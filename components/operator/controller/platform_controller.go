// Copyright 2023 XigXog
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package controller

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	vapi "github.com/hashicorp/vault/api"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/xigxog/kubefox/api"
	"github.com/xigxog/kubefox/api/kubernetes/v1alpha1"
	"github.com/xigxog/kubefox/build"
	"github.com/xigxog/kubefox/components/operator/defaults"
	"github.com/xigxog/kubefox/components/operator/templates"
	opvault "github.com/xigxog/kubefox/components/operator/vault"
	"github.com/xigxog/kubefox/core"
	"github.com/xigxog/kubefox/k8s"
	"github.com/xigxog/kubefox/logkf"
	"github.com/xigxog/kubefox/vault"
)

const (
	TenYears     string = "87600h"
	HundredYears string = "876000h"
)

// PlatformReconciler reconciles a Platform object
type PlatformReconciler struct {
	*Client

	Instance  string
	Namespace string
	VaultURL  string

	LogLevel  string
	LogFormat string

	CompMgr *ComponentManager

	setupMap map[string]bool

	log *logkf.Logger

	mutex sync.Mutex
}

// SetupWithManager sets up the controller with the Manager.
func (r *PlatformReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.log = logkf.Global.With(logkf.KeyController, "Platform")
	r.setupMap = make(map[string]bool)
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Platform{}).
		Owns(&appsv1.DaemonSet{}).
		Owns(&appsv1.Deployment{}).
		Owns(&appsv1.StatefulSet{}).
		Complete(r)
}

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *PlatformReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.log.With(
		"namespace", req.Namespace,
		"name", req.Name,
	)
	setupKey := fmt.Sprintf("%s/%s", req.Namespace, req.Name)

	ns := &v1.Namespace{}
	if err := r.Get(ctx, k8s.Key("", req.Namespace), ns); err != nil {
		r.setSetup(setupKey, false)
		return ctrl.Result{}, k8s.IgnoreNotFound(err)
	}
	if ns.Status.Phase == v1.NamespaceTerminating {
		log.Debugf("Namespace '%s' is terminating", ns.Name)
		r.setSetup(setupKey, false)
		return ctrl.Result{}, nil
	}

	platform := &v1alpha1.Platform{}
	if err := r.Get(ctx, req.NamespacedName, platform); err != nil {
		r.setSetup(setupKey, false)
		return ctrl.Result{}, k8s.IgnoreNotFound(err)
	}

	log.Debugf("reconciling '%s'", k8s.ToString(platform))
	defer log.Debugf("reconciling '%s' complete", k8s.ToString(platform))

	reconcileErr := r.reconcile(ctx, platform, log)
	if reconcileErr != nil {
		platform.Status.Conditions = k8s.UpdateConditions(metav1.Now(), platform.Status.Conditions, &metav1.Condition{
			Type:               api.ConditionTypeAvailable,
			Status:             metav1.ConditionUnknown,
			ObservedGeneration: platform.ObjectMeta.Generation,
			Reason:             api.ConditionReasonReconcileFailed,
			Message:            reconcileErr.Error(),
		})

	} else {
		if _, err := r.CompMgr.ReconcileApps(ctx, req.Namespace); err != nil {
			reconcileErr = err
		}
	}

	if err := r.updateComponentsStatus(ctx, platform); err != nil {
		if reconcileErr == nil {
			reconcileErr = err
		}
	}

	if err := r.Status().Update(ctx, platform); err != nil {
		if reconcileErr == nil {
			reconcileErr = err
		}
	}

	return RetryConflictWebhookErr(reconcileErr)
}

func (r *PlatformReconciler) reconcile(ctx context.Context, platform *v1alpha1.Platform, log *logkf.Logger) error {
	if platform.Spec.Debug.Enabled && platform.Name != "debug" && platform.Namespace != "kubefox-debug" {
		log.Warn("debug is enabled for Platform but name or namespace are incorrect, skipping reconcile")
		log.Debug("to enable debug Platform name must be 'debug' and namespace must be 'kubefox-debug'")

		return nil
	}

	setupKey := k8s.ToString(platform)

	cm := &v1.ConfigMap{}
	if err := r.Get(ctx, k8s.Key(r.Namespace, r.Instance+"-root-ca"), cm); err != nil {
		r.setSetup(setupKey, false)
		return log.ErrorN("unable to fetch root CA configmap: %w", err)
	}

	maxEventSize := platform.Spec.Events.MaxSize.Value()
	if platform.Spec.Events.MaxSize.IsZero() {
		maxEventSize = api.DefaultMaxEventSizeBytes
	}
	platformTD := &TemplateData{
		Data: templates.Data{
			Instance: templates.Instance{
				Name:           r.Instance,
				Namespace:      r.Namespace,
				RootCA:         cm.Data["ca.crt"],
				BootstrapImage: BootstrapImage,
			},
			Platform: templates.Platform{
				Name:       platform.Name,
				Namespace:  platform.Namespace,
				BrokerAddr: "$(KUBEFOX_HOST_IP):6060",
			},
			Owner: []*metav1.OwnerReference{
				metav1.NewControllerRef(platform, platform.GroupVersionKind()),
			},
			BuildInfo: build.Info,
			Telemetry: platform.Spec.Telemetry,
			Values: map[string]any{
				api.ValKeyMaxEventSize: maxEventSize,
				api.ValKeyVaultURL:     r.VaultURL,
			},
		},
	}
	if d := platform.Spec.Debug; d.Enabled && d.BrokerAddr != "" {
		platformTD.Data.Platform.BrokerAddr = d.BrokerAddr
	}

	if r.isSetup(setupKey) {
		r.log.Debugf("Platform '%s' already setup ", setupKey)

	} else {
		// Ensure there are valid commits for Platform components.
		if !api.RegexpCommit.MatchString(build.Info.BrokerCommit) ||
			!api.RegexpCommit.MatchString(build.Info.HTTPSrvCommit) {
			log.Error("broker or httpsrv commit from build info is invalid")
			return nil
		}
		if err := r.setupVaultPlatform(ctx, platform); err != nil {
			return log.ErrorN("problem setting up vault: %w", err)
		}

		r.setSetup(setupKey, true)
	}

	if err := r.ApplyTemplate(ctx, "platform", &platformTD.Data, log); err != nil {
		return log.ErrorN("problem setting up Platform: %w", err)
	}

	td := platformTD.ForComponent(api.PlatformComponentNATS, &appsv1.StatefulSet{}, &defaults.NATS, templates.Component{
		Component: core.NewPlatformComponent(
			api.ComponentTypeNATS,
			api.PlatformComponentNATS,
			"",
		),
		Image:               NATSImage,
		ImagePullPolicy:     platform.Spec.ImagePullPolicy,
		PodSpec:             platform.Spec.NATS.PodSpec,
		ContainerSpec:       platform.Spec.NATS.ContainerSpec,
		IsPlatformComponent: true,
	})
	if err := r.setupVaultComponent(ctx, td, false); err != nil {
		return err
	}
	if rdy, err := r.CompMgr.SetupComponent(ctx, td); !rdy || err != nil {
		platform.Status.Conditions = k8s.UpdateConditions(metav1.Now(), platform.Status.Conditions, &metav1.Condition{
			Type:               api.ConditionTypeAvailable,
			Status:             metav1.ConditionFalse,
			ObservedGeneration: platform.ObjectMeta.Generation,
			Reason:             api.ConditionReasonNATSUnavailable,
			Message:            fmt.Sprintf(`NATS StatefulSet "%s" is unavailable.`, td.Obj.GetName()),
		})
		return chill(err)
	}

	td = platformTD.ForComponent(api.PlatformComponentBroker, &appsv1.DaemonSet{}, &defaults.Broker, templates.Component{
		Component: core.NewPlatformComponent(
			api.ComponentTypeBroker,
			api.PlatformComponentBroker,
			build.Info.BrokerCommit,
		),
		Image:               BrokerImage,
		ImagePullPolicy:     platform.Spec.ImagePullPolicy,
		PodSpec:             platform.Spec.Broker.PodSpec,
		ContainerSpec:       platform.Spec.Broker.ContainerSpec,
		IsPlatformComponent: true,
	})
	if err := r.setupVaultComponent(ctx, td, true); err != nil {
		return err
	}
	if rdy, err := r.CompMgr.SetupComponent(ctx, td); !rdy || err != nil {
		platform.Status.Conditions = k8s.UpdateConditions(metav1.Now(), platform.Status.Conditions, &metav1.Condition{
			Type:               api.ConditionTypeAvailable,
			Status:             metav1.ConditionFalse,
			ObservedGeneration: platform.ObjectMeta.Generation,
			Reason:             api.ConditionReasonBrokerUnavailable,
			Message:            fmt.Sprintf(`Broker DaemonSet "%s" is unavailable.`, td.Obj.GetName()),
		})
		return chill(err)
	}

	td = platformTD.ForComponent(api.PlatformComponentHTTPSrv, &appsv1.Deployment{}, &defaults.HTTPSrv, templates.Component{
		Component: core.NewPlatformComponent(
			api.ComponentTypeHTTPAdapter,
			api.PlatformComponentHTTPSrv,
			build.Info.HTTPSrvCommit,
		),
		Image:               HTTPSrvImage,
		ImagePullPolicy:     platform.Spec.ImagePullPolicy,
		PodSpec:             platform.Spec.HTTPSrv.PodSpec,
		ContainerSpec:       platform.Spec.HTTPSrv.ContainerSpec,
		IsPlatformComponent: true,
	})
	td.Values["serviceLabels"] = platform.Spec.HTTPSrv.Service.Labels
	td.Values["serviceAnnotations"] = platform.Spec.HTTPSrv.Service.Annotations
	td.Values["serviceType"] = platform.Spec.HTTPSrv.Service.Type
	td.Values["httpPort"] = platform.Spec.HTTPSrv.Service.Ports.HTTP
	td.Values["httpsPort"] = platform.Spec.HTTPSrv.Service.Ports.HTTPS
	if err := r.setupVaultComponent(ctx, td, false); err != nil {
		return err
	}
	if rdy, err := r.CompMgr.SetupComponent(ctx, td); !rdy || err != nil {
		platform.Status.Conditions = k8s.UpdateConditions(metav1.Now(), platform.Status.Conditions, &metav1.Condition{
			Type:               api.ConditionTypeAvailable,
			Status:             metav1.ConditionFalse,
			ObservedGeneration: platform.ObjectMeta.Generation,
			Reason:             api.ConditionReasonHTTPSrvUnavailable,
			Message:            fmt.Sprintf(`HTTPSrv Deployment "%s" is unavailable.`, td.Obj.GetName()),
		})
		return chill(err)
	}

	platform.Status.Conditions = k8s.UpdateConditions(metav1.Now(), platform.Status.Conditions, &metav1.Condition{
		Type:               api.ConditionTypeAvailable,
		Status:             metav1.ConditionTrue,
		ObservedGeneration: platform.ObjectMeta.Generation,
		Reason:             api.ConditionReasonPlatformComponentsAvailable,
		Message:            "Platform Components are available.",
	})

	return nil
}

func (r *PlatformReconciler) updateComponentsStatus(ctx context.Context, p *v1alpha1.Platform) error {
	p.Status.Components = nil

	req, err := labels.NewRequirement(api.LabelK8sComponentType,
		selection.NotEquals, []string{string(api.ComponentTypeKubeFox)},
	)
	if err != nil {
		return err
	}
	podList := &v1.PodList{}
	if err := r.List(ctx, podList,
		client.InNamespace(p.Namespace),
		client.MatchingLabelsSelector{Selector: labels.NewSelector().Add(*req)},
	); err != nil {
		return err
	}

	for _, pod := range podList.Items {
		cond := k8s.PodCondition(&pod, v1.PodReady)
		p.Status.Components = append(p.Status.Components, v1alpha1.ComponentStatus{
			Ready:    cond.Status == v1.ConditionTrue,
			Name:     pod.Labels[api.LabelK8sComponent],
			Hash:     pod.Labels[api.LabelK8sComponentHash],
			Type:     api.ComponentType(pod.Labels[api.LabelK8sComponentType]),
			PodName:  pod.Name,
			PodIP:    pod.Status.PodIP,
			NodeName: pod.Spec.NodeName,
			NodeIP:   pod.Status.HostIP,
		})
	}

	return nil
}

func (r *PlatformReconciler) setupVaultPlatform(ctx context.Context, platform *v1alpha1.Platform) error {
	r.log.Debugf("setting up Vault for Platform '%s'", platform.Name)

	vaultCli, err := opvault.GetClient(ctx)
	if err != nil {
		return err
	}

	rootKey := vault.Key{Instance: r.Instance}
	key := vault.Key{
		Instance:  r.Instance,
		Namespace: platform.Namespace,
	}

	// Setup Env/VE Data stores.
	if err := vaultCli.CreateDataStore(ctx, ""); err != nil {
		return fmt.Errorf("error creating cluster data store: %w", err)
	}
	if err := vaultCli.CreateDataStore(ctx, platform.Namespace); err != nil {
		return fmt.Errorf("error creating namespace data store: %w", err)
	}

	// Setup PKI.
	// If cfg is non-nil then the mount already exists.
	if cfg, _ := vaultCli.Sys().MountConfig(vault.PKIPath(key)); cfg == nil {
		if err := vaultCli.Sys().Mount(vault.PKIPath(key), &vapi.MountInput{
			Type:        "pki",
			Description: "KubeFox Platform Intermediate CA",
			Config: vapi.MountConfigInput{
				MaxLeaseTTL: HundredYears,
			},
		}); err != nil {
			return err
		}
		if _, err := vaultCli.Logical().Write(vault.PKISubPath(key, "config/urls"),
			map[string]interface{}{
				"issuing_certificates":    fmt.Sprintf("%s/v1/%s", r.VaultURL, vault.PKISubPath(key, "ca")),
				"crl_distribution_points": fmt.Sprintf("%s/v1/%s", r.VaultURL, vault.PKISubPath(key, "crl")),
			},
		); err != nil {
			return err
		}

		// Generate intermediate cert and use root certificate to sign.
		intCert, err := vaultCli.Logical().Write(vault.PKISubPath(key, "intermediate/generate/internal"),
			map[string]interface{}{
				"common_name": "KubeFox Platform Intermediate CA",
				"issuer_name": "kubefox-platform-intermediate",
			},
		)
		if err != nil {
			return err
		}
		signedIntCert, err := vaultCli.Logical().Write(vault.PKISubPath(rootKey, "root/sign-intermediate"),
			map[string]interface{}{
				"csr":    intCert.Data["csr"],
				"format": "pem_bundle",
				"ttl":    HundredYears,
			},
		)
		if err != nil {
			return err
		}
		if _, err := vaultCli.Logical().Write(vault.PKISubPath(key, "intermediate/set-signed"),
			map[string]interface{}{
				"certificate": signedIntCert.Data["certificate"],
			},
		); err != nil {
			return err
		}
	}

	r.log.Debugf("Vault successfully setup for Platform '%s'", platform.Name)

	return nil
}

func (r *PlatformReconciler) setupVaultComponent(ctx context.Context, td *TemplateData, grantReadData bool) error {

	if r.isSetup(td.Name()) {
		r.log.Debugf("Vault already setup for Component '%s'", td.Component.Name)
		return nil
	}

	vaultCli, err := opvault.GetClient(ctx)
	if err != nil {
		return err
	}

	key := vault.Key{
		Instance:  r.Instance,
		Namespace: td.Platform.Namespace,
		Component: td.Component.Name,
	}

	r.log.Debugf("setting up Vault for Component '%s' with role '%s'", td.Component.Name, vault.RoleName(key))

	svcName := fmt.Sprintf("%s.%s", td.Name(), td.Platform.Namespace)
	if _, err := vaultCli.Logical().Write(vault.PKISubPath(key, "roles/"+vault.RoleName(key)),
		map[string]interface{}{
			"issuer_ref":         "default",
			"allowed_domains":    fmt.Sprintf("%s,%s.svc", svcName, svcName),
			"allow_localhost":    true,
			"allow_bare_domains": true,
			"max_ttl":            TenYears,
		},
	); err != nil {
		return err
	}

	var policies []string

	issueCertsPolicy := vault.PolicyName(key, "issue-certs")
	if err := vaultCli.Sys().PutPolicyWithContext(ctx, issueCertsPolicy, `
		path "`+vault.PKISubPath(key, "issue/"+vault.RoleName(key))+`" {
			capabilities = ["create", "update"]
		}
	`); err != nil {
		return err
	}
	policies = append(policies, issueCertsPolicy)

	if grantReadData {
		readDataPolicy := vault.PolicyName(key, "read-data")
		if err = vaultCli.Sys().PutPolicyWithContext(ctx, readDataPolicy, `
			path "`+vault.DataPath(api.DataKey{Instance: r.Instance})+`/*" {
				capabilities = ["read", "list"]
			}
			path "`+vault.DataPath(api.DataKey{Instance: r.Instance, Namespace: td.Platform.Namespace})+`/*" {
				capabilities = ["read", "list"]
			}
		`); err != nil {
			return err
		}
		policies = append(policies, readDataPolicy)
	}

	if _, err = vaultCli.Logical().Write(vault.KubernetesRolePath(key),
		map[string]interface{}{
			"bound_service_account_names":      td.Name(),
			"bound_service_account_namespaces": td.Platform.Namespace,
			"token_policies":                   strings.Join(policies, ","),
		},
	); err != nil {
		return err
	}

	r.setSetup(td.Name(), true)
	r.log.Debugf("Vault successfully setup for Component '%s'", td.Component.Name)

	return nil
}

func (r *PlatformReconciler) setSetup(key string, val bool) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.setupMap[key] = val
}

func (r *PlatformReconciler) isSetup(key string) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.setupMap[key]
}

// chill waits a few seconds for things to chillax.
func chill(err error) error {
	time.Sleep(time.Second * 3)
	return err
}
