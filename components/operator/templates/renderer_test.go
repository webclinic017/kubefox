package templates

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestRenderVault(t *testing.T) {
	d := &Data{
		Instance: Instance{
			Name:      "kubefox",
			Namespace: "kubefox-system",
		},
		Component: Component{
			Name:  "vault",
			Image: "ghcr.io/xigxog/vault:1.14.1-v0.2.1-alpha",
		},
	}
	if s, err := renderStr("list.tpl", "vault/*", d); err != nil {
		t.Errorf("%v", err)
	} else {
		t.Log(s)
	}
}

func TestRenderPlatform(t *testing.T) {
	d := &Data{
		Instance: Instance{
			Name:      "kubefox",
			Namespace: "kubefox-system",
		},
		Platform: Platform{
			Name:      "dev",
			Namespace: "kubefox-platform",
		},
		Owner: []*metav1.OwnerReference{
			{
				APIVersion: "kubefox.xigxog.io/v1alpha1",
				Kind:       "Platform",
				UID:        "123",
				Name:       "kubefox-dev",
			},
		},
	}
	if s, err := renderStr("list.tpl", "platform/*", d); err != nil {
		t.Errorf("%v", err)
	} else {
		t.Logf("\n%s", s)
	}
}

func TestRenderNATS(t *testing.T) {
	d := &Data{
		Instance: Instance{
			Name:      "kubefox",
			Namespace: "kubefox-system",
		},
		Platform: Platform{
			Name:      "dev",
			Namespace: "kubefox-platform",
		},
		Component: Component{
			Name:  "nats",
			Image: "nats:2.9.21-alpine",
		},
		Owner: []*metav1.OwnerReference{
			{
				APIVersion: "kubefox.xigxog.io/v1alpha1",
				Kind:       "Platform",
				UID:        "123",
				Name:       "kubefox-dev",
			},
		},
	}
	if s, err := renderStr("list.tpl", "nats/*", d); err != nil {
		t.Errorf("%v", err)
	} else {
		t.Logf("\n%s", s)
	}
}

func TestRenderBroker(t *testing.T) {
	d := &Data{
		Instance: Instance{
			Name:      "kubefox",
			Namespace: "kubefox-system",
		},
		Platform: Platform{
			Name:      "dev",
			Namespace: "kubefox-platform",
		},
		Component: Component{
			Name:  "broker",
			Image: "ghcr.io/xigxog/kubefox/broker:v0.0.1",
		},
		Owner: []*metav1.OwnerReference{
			{
				APIVersion: "kubefox.xigxog.io/v1alpha1",
				Kind:       "Platform",
				UID:        "123",
				Name:       "kubefox-dev",
			},
		},
	}
	if s, err := renderStr("list.tpl", "broker/*", d); err != nil {
		t.Errorf("%v", err)
	} else {
		t.Logf("\n%s", s)
	}
}

func TestRenderComponent(t *testing.T) {
	d := &Data{
		Instance: Instance{
			Name:      "kubefox",
			Namespace: "kubefox-system",
		},
		Platform: Platform{
			Name:      "dev",
			Namespace: "kubefox-platform",
		},
		Component: Component{
			Name:  "hello",
			Image: "ghcr.io/xigxog/kubefox/hello:v0.0.1",
		},
		Owner: []*metav1.OwnerReference{
			{
				APIVersion: "kubefox.xigxog.io/v1alpha1",
				Kind:       "Platform",
				UID:        "123",
				Name:       "kubefox-dev",
			},
			{
				APIVersion: "kubefox.xigxog.io/v1alpha1",
				Kind:       "Deployment",
				UID:        "123",
				Name:       "kubefox-dev",
			},
		},
	}
	if s, err := renderStr("list.tpl", "component/*", d); err != nil {
		t.Errorf("%v", err)
	} else {
		t.Logf("\n%s", s)
	}
}
