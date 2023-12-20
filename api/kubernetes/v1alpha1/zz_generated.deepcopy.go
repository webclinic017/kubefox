//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/xigxog/kubefox/api"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *App) DeepCopyInto(out *App) {
	*out = *in
	out.App = in.App
	in.CommitTime.DeepCopyInto(&out.CommitTime)
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make(map[string]*Component, len(*in))
		for key, val := range *in {
			var outVal *Component
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(Component)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new App.
func (in *App) DeepCopy() *App {
	if in == nil {
		return nil
	}
	out := new(App)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeployment) DeepCopyInto(out *AppDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	in.Details.DeepCopyInto(&out.Details)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeployment.
func (in *AppDeployment) DeepCopy() *AppDeployment {
	if in == nil {
		return nil
	}
	out := new(AppDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeploymentDetails) DeepCopyInto(out *AppDeploymentDetails) {
	*out = *in
	out.App = in.App
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make(map[string]api.Details, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeploymentDetails.
func (in *AppDeploymentDetails) DeepCopy() *AppDeploymentDetails {
	if in == nil {
		return nil
	}
	out := new(AppDeploymentDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeploymentList) DeepCopyInto(out *AppDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeploymentList.
func (in *AppDeploymentList) DeepCopy() *AppDeploymentList {
	if in == nil {
		return nil
	}
	out := new(AppDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeploymentSpec) DeepCopyInto(out *AppDeploymentSpec) {
	*out = *in
	in.App.DeepCopyInto(&out.App)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeploymentSpec.
func (in *AppDeploymentSpec) DeepCopy() *AppDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(AppDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDeploymentStatus) DeepCopyInto(out *AppDeploymentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDeploymentStatus.
func (in *AppDeploymentStatus) DeepCopy() *AppDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(AppDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerSpec) DeepCopyInto(out *BrokerSpec) {
	*out = *in
	in.PodSpec.DeepCopyInto(&out.PodSpec)
	in.ContainerSpec.DeepCopyInto(&out.ContainerSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerSpec.
func (in *BrokerSpec) DeepCopy() *BrokerSpec {
	if in == nil {
		return nil
	}
	out := new(BrokerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVirtualEnv) DeepCopyInto(out *ClusterVirtualEnv) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Data.DeepCopyInto(&out.Data)
	in.Details.DeepCopyInto(&out.Details)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVirtualEnv.
func (in *ClusterVirtualEnv) DeepCopy() *ClusterVirtualEnv {
	if in == nil {
		return nil
	}
	out := new(ClusterVirtualEnv)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterVirtualEnv) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVirtualEnvList) DeepCopyInto(out *ClusterVirtualEnvList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterVirtualEnv, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVirtualEnvList.
func (in *ClusterVirtualEnvList) DeepCopy() *ClusterVirtualEnvList {
	if in == nil {
		return nil
	}
	out := new(ClusterVirtualEnvList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterVirtualEnvList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVirtualEnvSpec) DeepCopyInto(out *ClusterVirtualEnvSpec) {
	*out = *in
	if in.ReleasePolicies != nil {
		in, out := &in.ReleasePolicies, &out.ReleasePolicies
		*out = new(ReleasePolicies)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVirtualEnvSpec.
func (in *ClusterVirtualEnvSpec) DeepCopy() *ClusterVirtualEnvSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterVirtualEnvSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Component) DeepCopyInto(out *Component) {
	*out = *in
	in.ComponentDefinition.DeepCopyInto(&out.ComponentDefinition)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component.
func (in *Component) DeepCopy() *Component {
	if in == nil {
		return nil
	}
	out := new(Component)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentStatus) DeepCopyInto(out *ComponentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentStatus.
func (in *ComponentStatus) DeepCopy() *ComponentStatus {
	if in == nil {
		return nil
	}
	out := new(ComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorSource) DeepCopyInto(out *ErrorSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorSource.
func (in *ErrorSource) DeepCopy() *ErrorSource {
	if in == nil {
		return nil
	}
	out := new(ErrorSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventsSpec) DeepCopyInto(out *EventsSpec) {
	*out = *in
	out.MaxSize = in.MaxSize.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventsSpec.
func (in *EventsSpec) DeepCopy() *EventsSpec {
	if in == nil {
		return nil
	}
	out := new(EventsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPAdapter) DeepCopyInto(out *HTTPAdapter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Details = in.Details
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPAdapter.
func (in *HTTPAdapter) DeepCopy() *HTTPAdapter {
	if in == nil {
		return nil
	}
	out := new(HTTPAdapter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HTTPAdapter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPAdapterList) DeepCopyInto(out *HTTPAdapterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HTTPAdapter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPAdapterList.
func (in *HTTPAdapterList) DeepCopy() *HTTPAdapterList {
	if in == nil {
		return nil
	}
	out := new(HTTPAdapterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HTTPAdapterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPAdapterSpec) DeepCopyInto(out *HTTPAdapterSpec) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPAdapterSpec.
func (in *HTTPAdapterSpec) DeepCopy() *HTTPAdapterSpec {
	if in == nil {
		return nil
	}
	out := new(HTTPAdapterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSrvPorts) DeepCopyInto(out *HTTPSrvPorts) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSrvPorts.
func (in *HTTPSrvPorts) DeepCopy() *HTTPSrvPorts {
	if in == nil {
		return nil
	}
	out := new(HTTPSrvPorts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSrvService) DeepCopyInto(out *HTTPSrvService) {
	*out = *in
	out.Ports = in.Ports
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSrvService.
func (in *HTTPSrvService) DeepCopy() *HTTPSrvService {
	if in == nil {
		return nil
	}
	out := new(HTTPSrvService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSrvSpec) DeepCopyInto(out *HTTPSrvSpec) {
	*out = *in
	in.PodSpec.DeepCopyInto(&out.PodSpec)
	in.ContainerSpec.DeepCopyInto(&out.ContainerSpec)
	in.Service.DeepCopyInto(&out.Service)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSrvSpec.
func (in *HTTPSrvSpec) DeepCopy() *HTTPSrvSpec {
	if in == nil {
		return nil
	}
	out := new(HTTPSrvSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATSSpec) DeepCopyInto(out *NATSSpec) {
	*out = *in
	in.PodSpec.DeepCopyInto(&out.PodSpec)
	in.ContainerSpec.DeepCopyInto(&out.ContainerSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATSSpec.
func (in *NATSSpec) DeepCopy() *NATSSpec {
	if in == nil {
		return nil
	}
	out := new(NATSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Platform) DeepCopyInto(out *Platform) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Details = in.Details
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Platform.
func (in *Platform) DeepCopy() *Platform {
	if in == nil {
		return nil
	}
	out := new(Platform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Platform) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformDetails) DeepCopyInto(out *PlatformDetails) {
	*out = *in
	out.Details = in.Details
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformDetails.
func (in *PlatformDetails) DeepCopy() *PlatformDetails {
	if in == nil {
		return nil
	}
	out := new(PlatformDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformList) DeepCopyInto(out *PlatformList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Platform, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformList.
func (in *PlatformList) DeepCopy() *PlatformList {
	if in == nil {
		return nil
	}
	out := new(PlatformList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlatformList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformSpec) DeepCopyInto(out *PlatformSpec) {
	*out = *in
	in.Events.DeepCopyInto(&out.Events)
	in.Broker.DeepCopyInto(&out.Broker)
	in.HTTPSrv.DeepCopyInto(&out.HTTPSrv)
	in.NATS.DeepCopyInto(&out.NATS)
	out.Logger = in.Logger
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformSpec.
func (in *PlatformSpec) DeepCopy() *PlatformSpec {
	if in == nil {
		return nil
	}
	out := new(PlatformSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformStatus) DeepCopyInto(out *PlatformStatus) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]ComponentStatus, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformStatus.
func (in *PlatformStatus) DeepCopy() *PlatformStatus {
	if in == nil {
		return nil
	}
	out := new(PlatformStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Release) DeepCopyInto(out *Release) {
	*out = *in
	out.AppDeployment = in.AppDeployment
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Release.
func (in *Release) DeepCopy() *Release {
	if in == nil {
		return nil
	}
	out := new(Release)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseAppDeployment) DeepCopyInto(out *ReleaseAppDeployment) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseAppDeployment.
func (in *ReleaseAppDeployment) DeepCopy() *ReleaseAppDeployment {
	if in == nil {
		return nil
	}
	out := new(ReleaseAppDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseAppDeploymentStatus) DeepCopyInto(out *ReleaseAppDeploymentStatus) {
	*out = *in
	out.ReleaseAppDeployment = in.ReleaseAppDeployment
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseAppDeploymentStatus.
func (in *ReleaseAppDeploymentStatus) DeepCopy() *ReleaseAppDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(ReleaseAppDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseError) DeepCopyInto(out *ReleaseError) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(ErrorSource)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseError.
func (in *ReleaseError) DeepCopy() *ReleaseError {
	if in == nil {
		return nil
	}
	out := new(ReleaseError)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseHistoryLimits) DeepCopyInto(out *ReleaseHistoryLimits) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseHistoryLimits.
func (in *ReleaseHistoryLimits) DeepCopy() *ReleaseHistoryLimits {
	if in == nil {
		return nil
	}
	out := new(ReleaseHistoryLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleasePolicies) DeepCopyInto(out *ReleasePolicies) {
	*out = *in
	if in.HistoryLimits != nil {
		in, out := &in.HistoryLimits, &out.HistoryLimits
		*out = new(ReleaseHistoryLimits)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleasePolicies.
func (in *ReleasePolicies) DeepCopy() *ReleasePolicies {
	if in == nil {
		return nil
	}
	out := new(ReleasePolicies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseStatus) DeepCopyInto(out *ReleaseStatus) {
	*out = *in
	out.AppDeployment = in.AppDeployment
	in.RequestTime.DeepCopyInto(&out.RequestTime)
	if in.ActivationTime != nil {
		in, out := &in.ActivationTime, &out.ActivationTime
		*out = (*in).DeepCopy()
	}
	if in.ArchiveTime != nil {
		in, out := &in.ArchiveTime, &out.ArchiveTime
		*out = (*in).DeepCopy()
	}
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]ReleaseError, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseStatus.
func (in *ReleaseStatus) DeepCopy() *ReleaseStatus {
	if in == nil {
		return nil
	}
	out := new(ReleaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualEnv) DeepCopyInto(out *VirtualEnv) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Data.DeepCopyInto(&out.Data)
	in.Details.DeepCopyInto(&out.Details)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualEnv.
func (in *VirtualEnv) DeepCopy() *VirtualEnv {
	if in == nil {
		return nil
	}
	out := new(VirtualEnv)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualEnv) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualEnvData) DeepCopyInto(out *VirtualEnvData) {
	*out = *in
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make(map[string]*api.Val, len(*in))
		for key, val := range *in {
			var outVal *api.Val
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(api.Val)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make(map[string]*api.Val, len(*in))
		for key, val := range *in {
			var outVal *api.Val
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(api.Val)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualEnvData.
func (in *VirtualEnvData) DeepCopy() *VirtualEnvData {
	if in == nil {
		return nil
	}
	out := new(VirtualEnvData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualEnvDetails) DeepCopyInto(out *VirtualEnvDetails) {
	*out = *in
	out.Details = in.Details
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make(map[string]api.Details, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make(map[string]api.Details, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualEnvDetails.
func (in *VirtualEnvDetails) DeepCopy() *VirtualEnvDetails {
	if in == nil {
		return nil
	}
	out := new(VirtualEnvDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualEnvList) DeepCopyInto(out *VirtualEnvList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualEnv, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualEnvList.
func (in *VirtualEnvList) DeepCopy() *VirtualEnvList {
	if in == nil {
		return nil
	}
	out := new(VirtualEnvList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualEnvList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualEnvSnapshot) DeepCopyInto(out *VirtualEnvSnapshot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(VirtualEnvData)
		(*in).DeepCopyInto(*out)
	}
	in.Details.DeepCopyInto(&out.Details)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualEnvSnapshot.
func (in *VirtualEnvSnapshot) DeepCopy() *VirtualEnvSnapshot {
	if in == nil {
		return nil
	}
	out := new(VirtualEnvSnapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualEnvSnapshot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualEnvSnapshotList) DeepCopyInto(out *VirtualEnvSnapshotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualEnvSnapshot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualEnvSnapshotList.
func (in *VirtualEnvSnapshotList) DeepCopy() *VirtualEnvSnapshotList {
	if in == nil {
		return nil
	}
	out := new(VirtualEnvSnapshotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualEnvSnapshotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualEnvSnapshotSpec) DeepCopyInto(out *VirtualEnvSnapshotSpec) {
	*out = *in
	out.Source = in.Source
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualEnvSnapshotSpec.
func (in *VirtualEnvSnapshotSpec) DeepCopy() *VirtualEnvSnapshotSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualEnvSnapshotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualEnvSource) DeepCopyInto(out *VirtualEnvSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualEnvSource.
func (in *VirtualEnvSource) DeepCopy() *VirtualEnvSource {
	if in == nil {
		return nil
	}
	out := new(VirtualEnvSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualEnvSpec) DeepCopyInto(out *VirtualEnvSpec) {
	*out = *in
	if in.Release != nil {
		in, out := &in.Release, &out.Release
		*out = new(Release)
		**out = **in
	}
	if in.ReleasePolicies != nil {
		in, out := &in.ReleasePolicies, &out.ReleasePolicies
		*out = new(ReleasePolicies)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualEnvSpec.
func (in *VirtualEnvSpec) DeepCopy() *VirtualEnvSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualEnvSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualEnvStatus) DeepCopyInto(out *VirtualEnvStatus) {
	*out = *in
	if in.ActiveRelease != nil {
		in, out := &in.ActiveRelease, &out.ActiveRelease
		*out = new(ReleaseStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.PendingRelease != nil {
		in, out := &in.PendingRelease, &out.PendingRelease
		*out = new(ReleaseStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.ReleaseHistory != nil {
		in, out := &in.ReleaseHistory, &out.ReleaseHistory
		*out = make([]ReleaseStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualEnvStatus.
func (in *VirtualEnvStatus) DeepCopy() *VirtualEnvStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualEnvStatus)
	in.DeepCopyInto(out)
	return out
}
