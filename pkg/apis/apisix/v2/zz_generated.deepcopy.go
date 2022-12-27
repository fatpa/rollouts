//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Kruise Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v2

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRoute) DeepCopyInto(out *ApisixRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRoute.
func (in *ApisixRoute) DeepCopy() *ApisixRoute {
	if in == nil {
		return nil
	}
	out := new(ApisixRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApisixRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteAuthentication) DeepCopyInto(out *ApisixRouteAuthentication) {
	*out = *in
	out.KeyAuth = in.KeyAuth
	out.JwtAuth = in.JwtAuth
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteAuthentication.
func (in *ApisixRouteAuthentication) DeepCopy() *ApisixRouteAuthentication {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteAuthenticationJwtAuth) DeepCopyInto(out *ApisixRouteAuthenticationJwtAuth) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteAuthenticationJwtAuth.
func (in *ApisixRouteAuthenticationJwtAuth) DeepCopy() *ApisixRouteAuthenticationJwtAuth {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteAuthenticationJwtAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteAuthenticationKeyAuth) DeepCopyInto(out *ApisixRouteAuthenticationKeyAuth) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteAuthenticationKeyAuth.
func (in *ApisixRouteAuthenticationKeyAuth) DeepCopy() *ApisixRouteAuthenticationKeyAuth {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteAuthenticationKeyAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteHTTP) DeepCopyInto(out *ApisixRouteHTTP) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(UpstreamTimeout)
		**out = **in
	}
	in.Match.DeepCopyInto(&out.Match)
	if in.Backends != nil {
		in, out := &in.Backends, &out.Backends
		*out = make([]ApisixRouteHTTPBackend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make([]ApisixRoutePlugin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = new(ApisixRouteAuthentication)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteHTTP.
func (in *ApisixRouteHTTP) DeepCopy() *ApisixRouteHTTP {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteHTTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteHTTPBackend) DeepCopyInto(out *ApisixRouteHTTPBackend) {
	*out = *in
	out.ServicePort = in.ServicePort
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteHTTPBackend.
func (in *ApisixRouteHTTPBackend) DeepCopy() *ApisixRouteHTTPBackend {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteHTTPBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteHTTPMatch) DeepCopyInto(out *ApisixRouteHTTPMatch) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RemoteAddrs != nil {
		in, out := &in.RemoteAddrs, &out.RemoteAddrs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NginxVars != nil {
		in, out := &in.NginxVars, &out.NginxVars
		*out = make([]ApisixRouteHTTPMatchExpr, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteHTTPMatch.
func (in *ApisixRouteHTTPMatch) DeepCopy() *ApisixRouteHTTPMatch {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteHTTPMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteHTTPMatchExpr) DeepCopyInto(out *ApisixRouteHTTPMatchExpr) {
	*out = *in
	out.Subject = in.Subject
	if in.Set != nil {
		in, out := &in.Set, &out.Set
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteHTTPMatchExpr.
func (in *ApisixRouteHTTPMatchExpr) DeepCopy() *ApisixRouteHTTPMatchExpr {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteHTTPMatchExpr)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteHTTPMatchExprSubject) DeepCopyInto(out *ApisixRouteHTTPMatchExprSubject) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteHTTPMatchExprSubject.
func (in *ApisixRouteHTTPMatchExprSubject) DeepCopy() *ApisixRouteHTTPMatchExprSubject {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteHTTPMatchExprSubject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteList) DeepCopyInto(out *ApisixRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApisixRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteList.
func (in *ApisixRouteList) DeepCopy() *ApisixRouteList {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApisixRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRoutePlugin) DeepCopyInto(out *ApisixRoutePlugin) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRoutePlugin.
func (in *ApisixRoutePlugin) DeepCopy() *ApisixRoutePlugin {
	if in == nil {
		return nil
	}
	out := new(ApisixRoutePlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteSpec) DeepCopyInto(out *ApisixRouteSpec) {
	*out = *in
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = make([]ApisixRouteHTTP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Stream != nil {
		in, out := &in.Stream, &out.Stream
		*out = make([]ApisixRouteStream, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteSpec.
func (in *ApisixRouteSpec) DeepCopy() *ApisixRouteSpec {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteStream) DeepCopyInto(out *ApisixRouteStream) {
	*out = *in
	out.Match = in.Match
	out.Backend = in.Backend
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make([]ApisixRoutePlugin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteStream.
func (in *ApisixRouteStream) DeepCopy() *ApisixRouteStream {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteStream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteStreamBackend) DeepCopyInto(out *ApisixRouteStreamBackend) {
	*out = *in
	out.ServicePort = in.ServicePort
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteStreamBackend.
func (in *ApisixRouteStreamBackend) DeepCopy() *ApisixRouteStreamBackend {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteStreamBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixRouteStreamMatch) DeepCopyInto(out *ApisixRouteStreamMatch) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixRouteStreamMatch.
func (in *ApisixRouteStreamMatch) DeepCopy() *ApisixRouteStreamMatch {
	if in == nil {
		return nil
	}
	out := new(ApisixRouteStreamMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixStatus) DeepCopyInto(out *ApisixStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixStatus.
func (in *ApisixStatus) DeepCopy() *ApisixStatus {
	if in == nil {
		return nil
	}
	out := new(ApisixStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpstreamTimeout) DeepCopyInto(out *UpstreamTimeout) {
	*out = *in
	out.Connect = in.Connect
	out.Send = in.Send
	out.Read = in.Read
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpstreamTimeout.
func (in *UpstreamTimeout) DeepCopy() *UpstreamTimeout {
	if in == nil {
		return nil
	}
	out := new(UpstreamTimeout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixUpstream) DeepCopyInto(out *ApisixUpstream) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(ApisixUpstreamSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixUpstream.
func (in *ApisixUpstream) DeepCopy() *ApisixUpstream {
	if in == nil {
		return nil
	}
	out := new(ApisixUpstream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApisixUpstream) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixUpstreamConfig) DeepCopyInto(out *ApisixUpstreamConfig) {
	*out = *in
	if in.LoadBalancer != nil {
		in, out := &in.LoadBalancer, &out.LoadBalancer
		*out = new(LoadBalancer)
		**out = **in
	}
	if in.Retries != nil {
		in, out := &in.Retries, &out.Retries
		*out = new(int)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(UpstreamTimeout)
		**out = **in
	}
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(HealthCheck)
		(*in).DeepCopyInto(*out)
	}
	if in.TLSSecret != nil {
		in, out := &in.TLSSecret, &out.TLSSecret
		*out = new(ApisixSecret)
		**out = **in
	}
	if in.Subsets != nil {
		in, out := &in.Subsets, &out.Subsets
		*out = make([]ApisixUpstreamSubset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixUpstreamConfig.
func (in *ApisixUpstreamConfig) DeepCopy() *ApisixUpstreamConfig {
	if in == nil {
		return nil
	}
	out := new(ApisixUpstreamConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixUpstreamExternalNode) DeepCopyInto(out *ApisixUpstreamExternalNode) {
	*out = *in
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixUpstreamExternalNode.
func (in *ApisixUpstreamExternalNode) DeepCopy() *ApisixUpstreamExternalNode {
	if in == nil {
		return nil
	}
	out := new(ApisixUpstreamExternalNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixUpstreamExternalType) DeepCopyInto(out *ApisixUpstreamExternalType) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixUpstreamExternalType.
func (in *ApisixUpstreamExternalType) DeepCopy() *ApisixUpstreamExternalType {
	if in == nil {
		return nil
	}
	out := new(ApisixUpstreamExternalType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixUpstreamList) DeepCopyInto(out *ApisixUpstreamList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApisixUpstream, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixUpstreamList.
func (in *ApisixUpstreamList) DeepCopy() *ApisixUpstreamList {
	if in == nil {
		return nil
	}
	out := new(ApisixUpstreamList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApisixUpstreamList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixUpstreamSpec) DeepCopyInto(out *ApisixUpstreamSpec) {
	*out = *in
	if in.ExternalNodes != nil {
		in, out := &in.ExternalNodes, &out.ExternalNodes
		*out = make([]ApisixUpstreamExternalNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ApisixUpstreamConfig.DeepCopyInto(&out.ApisixUpstreamConfig)
	if in.PortLevelSettings != nil {
		in, out := &in.PortLevelSettings, &out.PortLevelSettings
		*out = make([]PortLevelSettings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixUpstreamSpec.
func (in *ApisixUpstreamSpec) DeepCopy() *ApisixUpstreamSpec {
	if in == nil {
		return nil
	}
	out := new(ApisixUpstreamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApisixUpstreamSubset) DeepCopyInto(out *ApisixUpstreamSubset) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApisixUpstreamSubset.
func (in *ApisixUpstreamSubset) DeepCopy() *ApisixUpstreamSubset {
	if in == nil {
		return nil
	}
	out := new(ApisixUpstreamSubset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheck) DeepCopyInto(out *HealthCheck) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(ActiveHealthCheck)
		(*in).DeepCopyInto(*out)
	}
	if in.Passive != nil {
		in, out := &in.Passive, &out.Passive
		*out = new(PassiveHealthCheck)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheck.
func (in *HealthCheck) DeepCopy() *HealthCheck {
	if in == nil {
		return nil
	}
	out := new(HealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancer) DeepCopyInto(out *LoadBalancer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancer.
func (in *LoadBalancer) DeepCopy() *LoadBalancer {
	if in == nil {
		return nil
	}
	out := new(LoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PassiveHealthCheck) DeepCopyInto(out *PassiveHealthCheck) {
	*out = *in
	if in.Healthy != nil {
		in, out := &in.Healthy, &out.Healthy
		*out = new(PassiveHealthCheckHealthy)
		(*in).DeepCopyInto(*out)
	}
	if in.Unhealthy != nil {
		in, out := &in.Unhealthy, &out.Unhealthy
		*out = new(PassiveHealthCheckUnhealthy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PassiveHealthCheck.
func (in *PassiveHealthCheck) DeepCopy() *PassiveHealthCheck {
	if in == nil {
		return nil
	}
	out := new(PassiveHealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PassiveHealthCheckHealthy) DeepCopyInto(out *PassiveHealthCheckHealthy) {
	*out = *in
	if in.HTTPCodes != nil {
		in, out := &in.HTTPCodes, &out.HTTPCodes
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PassiveHealthCheckHealthy.
func (in *PassiveHealthCheckHealthy) DeepCopy() *PassiveHealthCheckHealthy {
	if in == nil {
		return nil
	}
	out := new(PassiveHealthCheckHealthy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PassiveHealthCheckUnhealthy) DeepCopyInto(out *PassiveHealthCheckUnhealthy) {
	*out = *in
	if in.HTTPCodes != nil {
		in, out := &in.HTTPCodes, &out.HTTPCodes
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PassiveHealthCheckUnhealthy.
func (in *PassiveHealthCheckUnhealthy) DeepCopy() *PassiveHealthCheckUnhealthy {
	if in == nil {
		return nil
	}
	out := new(PassiveHealthCheckUnhealthy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortLevelSettings) DeepCopyInto(out *PortLevelSettings) {
	*out = *in
	in.ApisixUpstreamConfig.DeepCopyInto(&out.ApisixUpstreamConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortLevelSettings.
func (in *PortLevelSettings) DeepCopy() *PortLevelSettings {
	if in == nil {
		return nil
	}
	out := new(PortLevelSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveHealthCheck) DeepCopyInto(out *ActiveHealthCheck) {
	*out = *in
	if in.StrictTLS != nil {
		in, out := &in.StrictTLS, &out.StrictTLS
		*out = new(bool)
		**out = **in
	}
	if in.RequestHeaders != nil {
		in, out := &in.RequestHeaders, &out.RequestHeaders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Healthy != nil {
		in, out := &in.Healthy, &out.Healthy
		*out = new(ActiveHealthCheckHealthy)
		(*in).DeepCopyInto(*out)
	}
	if in.Unhealthy != nil {
		in, out := &in.Unhealthy, &out.Unhealthy
		*out = new(ActiveHealthCheckUnhealthy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveHealthCheck.
func (in *ActiveHealthCheck) DeepCopy() *ActiveHealthCheck {
	if in == nil {
		return nil
	}
	out := new(ActiveHealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveHealthCheckHealthy) DeepCopyInto(out *ActiveHealthCheckHealthy) {
	*out = *in
	in.PassiveHealthCheckHealthy.DeepCopyInto(&out.PassiveHealthCheckHealthy)
	out.Interval = in.Interval
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveHealthCheckHealthy.
func (in *ActiveHealthCheckHealthy) DeepCopy() *ActiveHealthCheckHealthy {
	if in == nil {
		return nil
	}
	out := new(ActiveHealthCheckHealthy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveHealthCheckUnhealthy) DeepCopyInto(out *ActiveHealthCheckUnhealthy) {
	*out = *in
	in.PassiveHealthCheckUnhealthy.DeepCopyInto(&out.PassiveHealthCheckUnhealthy)
	out.Interval = in.Interval
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveHealthCheckUnhealthy.
func (in *ActiveHealthCheckUnhealthy) DeepCopy() *ActiveHealthCheckUnhealthy {
	if in == nil {
		return nil
	}
	out := new(ActiveHealthCheckUnhealthy)
	in.DeepCopyInto(out)
	return out
}
