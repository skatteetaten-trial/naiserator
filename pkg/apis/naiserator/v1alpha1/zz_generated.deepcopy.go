// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPolicy) DeepCopyInto(out *AccessPolicy) {
	*out = *in
	in.Ingress.DeepCopyInto(&out.Ingress)
	in.Egress.DeepCopyInto(&out.Egress)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPolicy.
func (in *AccessPolicy) DeepCopy() *AccessPolicy {
	if in == nil {
		return nil
	}
	out := new(AccessPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPolicyEgress) DeepCopyInto(out *AccessPolicyEgress) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]AccessPolicyGressRule, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPolicyEgress.
func (in *AccessPolicyEgress) DeepCopy() *AccessPolicyEgress {
	if in == nil {
		return nil
	}
	out := new(AccessPolicyEgress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPolicyGressRule) DeepCopyInto(out *AccessPolicyGressRule) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPolicyGressRule.
func (in *AccessPolicyGressRule) DeepCopy() *AccessPolicyGressRule {
	if in == nil {
		return nil
	}
	out := new(AccessPolicyGressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPolicyIngress) DeepCopyInto(out *AccessPolicyIngress) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]AccessPolicyGressRule, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPolicyIngress.
func (in *AccessPolicyIngress) DeepCopy() *AccessPolicyIngress {
	if in == nil {
		return nil
	}
	out := new(AccessPolicyIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Application) DeepCopyInto(out *Application) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Application.
func (in *Application) DeepCopy() *Application {
	if in == nil {
		return nil
	}
	out := new(Application)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Application) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationList) DeepCopyInto(out *ApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Application, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationList.
func (in *ApplicationList) DeepCopy() *ApplicationList {
	if in == nil {
		return nil
	}
	out := new(ApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpec) DeepCopyInto(out *ApplicationSpec) {
	*out = *in
	out.Liveness = in.Liveness
	out.Readiness = in.Readiness
	if in.Ingresses != nil {
		in, out := &in.Ingresses, &out.Ingresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.SecureLogs = in.SecureLogs
	out.Prometheus = in.Prometheus
	out.Replicas = in.Replicas
	out.Resources = in.Resources
	in.Vault.DeepCopyInto(&out.Vault)
	in.ConfigMaps.DeepCopyInto(&out.ConfigMaps)
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]EnvVar, len(*in))
		copy(*out, *in)
	}
	out.Strategy = in.Strategy
	out.Service = in.Service
	in.AccessPolicy.DeepCopyInto(&out.AccessPolicy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpec.
func (in *ApplicationSpec) DeepCopy() *ApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMaps) DeepCopyInto(out *ConfigMaps) {
	*out = *in
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMaps.
func (in *ConfigMaps) DeepCopy() *ConfigMaps {
	if in == nil {
		return nil
	}
	out := new(ConfigMaps)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVar) DeepCopyInto(out *EnvVar) {
	*out = *in
	out.ValueFrom = in.ValueFrom
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVar.
func (in *EnvVar) DeepCopy() *EnvVar {
	if in == nil {
		return nil
	}
	out := new(EnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVarSource) DeepCopyInto(out *EnvVarSource) {
	*out = *in
	out.FieldRef = in.FieldRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVarSource.
func (in *EnvVarSource) DeepCopy() *EnvVarSource {
	if in == nil {
		return nil
	}
	out := new(EnvVarSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectFieldSelector) DeepCopyInto(out *ObjectFieldSelector) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectFieldSelector.
func (in *ObjectFieldSelector) DeepCopy() *ObjectFieldSelector {
	if in == nil {
		return nil
	}
	out := new(ObjectFieldSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Probe) DeepCopyInto(out *Probe) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Probe.
func (in *Probe) DeepCopy() *Probe {
	if in == nil {
		return nil
	}
	out := new(Probe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusConfig) DeepCopyInto(out *PrometheusConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusConfig.
func (in *PrometheusConfig) DeepCopy() *PrometheusConfig {
	if in == nil {
		return nil
	}
	out := new(PrometheusConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Replicas) DeepCopyInto(out *Replicas) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Replicas.
func (in *Replicas) DeepCopy() *Replicas {
	if in == nil {
		return nil
	}
	out := new(Replicas)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequirements) DeepCopyInto(out *ResourceRequirements) {
	*out = *in
	out.Limits = in.Limits
	out.Requests = in.Requests
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequirements.
func (in *ResourceRequirements) DeepCopy() *ResourceRequirements {
	if in == nil {
		return nil
	}
	out := new(ResourceRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSpec) DeepCopyInto(out *ResourceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSpec.
func (in *ResourceSpec) DeepCopy() *ResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretPath) DeepCopyInto(out *SecretPath) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretPath.
func (in *SecretPath) DeepCopy() *SecretPath {
	if in == nil {
		return nil
	}
	out := new(SecretPath)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecureLogs) DeepCopyInto(out *SecureLogs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecureLogs.
func (in *SecureLogs) DeepCopy() *SecureLogs {
	if in == nil {
		return nil
	}
	out := new(SecureLogs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Strategy) DeepCopyInto(out *Strategy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Strategy.
func (in *Strategy) DeepCopy() *Strategy {
	if in == nil {
		return nil
	}
	out := new(Strategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vault) DeepCopyInto(out *Vault) {
	*out = *in
	if in.Mounts != nil {
		in, out := &in.Mounts, &out.Mounts
		*out = make([]SecretPath, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vault.
func (in *Vault) DeepCopy() *Vault {
	if in == nil {
		return nil
	}
	out := new(Vault)
	in.DeepCopyInto(out)
	return out
}
