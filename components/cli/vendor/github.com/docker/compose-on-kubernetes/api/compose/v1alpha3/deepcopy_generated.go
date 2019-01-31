// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by C:\gohome\bin\deepcopy-gen.exe. DO NOT EDIT.

package v1alpha3

import (
	time "time"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigObjConfig) DeepCopyInto(out *ConfigObjConfig) {
	*out = *in
	out.External = in.External
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigObjConfig.
func (in *ConfigObjConfig) DeepCopy() *ConfigObjConfig {
	if in == nil {
		return nil
	}
	out := new(ConfigObjConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Constraint) DeepCopyInto(out *Constraint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Constraint.
func (in *Constraint) DeepCopy() *Constraint {
	if in == nil {
		return nil
	}
	out := new(Constraint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Constraints) DeepCopyInto(out *Constraints) {
	*out = *in
	if in.OperatingSystem != nil {
		in, out := &in.OperatingSystem, &out.OperatingSystem
		if *in == nil {
			*out = nil
		} else {
			*out = new(Constraint)
			**out = **in
		}
	}
	if in.Architecture != nil {
		in, out := &in.Architecture, &out.Architecture
		if *in == nil {
			*out = nil
		} else {
			*out = new(Constraint)
			**out = **in
		}
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		if *in == nil {
			*out = nil
		} else {
			*out = new(Constraint)
			**out = **in
		}
	}
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]Constraint, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Constraints.
func (in *Constraints) DeepCopy() *Constraints {
	if in == nil {
		return nil
	}
	out := new(Constraints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployConfig) DeepCopyInto(out *DeployConfig) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		if *in == nil {
			*out = nil
		} else {
			*out = new(uint64)
			**out = **in
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.UpdateConfig != nil {
		in, out := &in.UpdateConfig, &out.UpdateConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(UpdateConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.RestartPolicy != nil {
		in, out := &in.RestartPolicy, &out.RestartPolicy
		if *in == nil {
			*out = nil
		} else {
			*out = new(RestartPolicy)
			**out = **in
		}
	}
	in.Placement.DeepCopyInto(&out.Placement)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployConfig.
func (in *DeployConfig) DeepCopy() *DeployConfig {
	if in == nil {
		return nil
	}
	out := new(DeployConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *External) DeepCopyInto(out *External) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new External.
func (in *External) DeepCopy() *External {
	if in == nil {
		return nil
	}
	out := new(External)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileObjectConfig) DeepCopyInto(out *FileObjectConfig) {
	*out = *in
	out.External = in.External
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileObjectConfig.
func (in *FileObjectConfig) DeepCopy() *FileObjectConfig {
	if in == nil {
		return nil
	}
	out := new(FileObjectConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileReferenceConfig) DeepCopyInto(out *FileReferenceConfig) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		if *in == nil {
			*out = nil
		} else {
			*out = new(uint32)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileReferenceConfig.
func (in *FileReferenceConfig) DeepCopy() *FileReferenceConfig {
	if in == nil {
		return nil
	}
	out := new(FileReferenceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckConfig) DeepCopyInto(out *HealthCheckConfig) {
	*out = *in
	if in.Test != nil {
		in, out := &in.Test, &out.Test
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		if *in == nil {
			*out = nil
		} else {
			*out = new(time.Duration)
			**out = **in
		}
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		if *in == nil {
			*out = nil
		} else {
			*out = new(time.Duration)
			**out = **in
		}
	}
	if in.Retries != nil {
		in, out := &in.Retries, &out.Retries
		if *in == nil {
			*out = nil
		} else {
			*out = new(uint64)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckConfig.
func (in *HealthCheckConfig) DeepCopy() *HealthCheckConfig {
	if in == nil {
		return nil
	}
	out := new(HealthCheckConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Placement) DeepCopyInto(out *Placement) {
	*out = *in
	if in.Constraints != nil {
		in, out := &in.Constraints, &out.Constraints
		if *in == nil {
			*out = nil
		} else {
			*out = new(Constraints)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Placement.
func (in *Placement) DeepCopy() *Placement {
	if in == nil {
		return nil
	}
	out := new(Placement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		if *in == nil {
			*out = nil
		} else {
			*out = new(Resource)
			**out = **in
		}
	}
	if in.Reservations != nil {
		in, out := &in.Reservations, &out.Reservations
		if *in == nil {
			*out = nil
		} else {
			*out = new(Resource)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestartPolicy) DeepCopyInto(out *RestartPolicy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestartPolicy.
func (in *RestartPolicy) DeepCopy() *RestartPolicy {
	if in == nil {
		return nil
	}
	out := new(RestartPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretConfig) DeepCopyInto(out *SecretConfig) {
	*out = *in
	out.External = in.External
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretConfig.
func (in *SecretConfig) DeepCopy() *SecretConfig {
	if in == nil {
		return nil
	}
	out := new(SecretConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceConfig) DeepCopyInto(out *ServiceConfig) {
	*out = *in
	if in.CapAdd != nil {
		in, out := &in.CapAdd, &out.CapAdd
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CapDrop != nil {
		in, out := &in.CapDrop, &out.CapDrop
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Configs != nil {
		in, out := &in.Configs, &out.Configs
		*out = make([]ServiceConfigObjConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Deploy.DeepCopyInto(&out.Deploy)
	if in.Entrypoint != nil {
		in, out := &in.Entrypoint, &out.Entrypoint
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				outVal := *val
				(*out)[key] = &outVal
			}
		}
	}
	if in.ExtraHosts != nil {
		in, out := &in.ExtraHosts, &out.ExtraHosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		if *in == nil {
			*out = nil
		} else {
			*out = new(HealthCheckConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]ServicePortConfig, len(*in))
		copy(*out, *in)
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]ServiceSecretConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StopGracePeriod != nil {
		in, out := &in.StopGracePeriod, &out.StopGracePeriod
		if *in == nil {
			*out = nil
		} else {
			*out = new(time.Duration)
			**out = **in
		}
	}
	if in.Tmpfs != nil {
		in, out := &in.Tmpfs, &out.Tmpfs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		if *in == nil {
			*out = nil
		} else {
			*out = new(int64)
			**out = **in
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]ServiceVolumeConfig, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceConfig.
func (in *ServiceConfig) DeepCopy() *ServiceConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceConfigObjConfig) DeepCopyInto(out *ServiceConfigObjConfig) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		if *in == nil {
			*out = nil
		} else {
			*out = new(uint32)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceConfigObjConfig.
func (in *ServiceConfigObjConfig) DeepCopy() *ServiceConfigObjConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceConfigObjConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicePortConfig) DeepCopyInto(out *ServicePortConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicePortConfig.
func (in *ServicePortConfig) DeepCopy() *ServicePortConfig {
	if in == nil {
		return nil
	}
	out := new(ServicePortConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSecretConfig) DeepCopyInto(out *ServiceSecretConfig) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		if *in == nil {
			*out = nil
		} else {
			*out = new(uint32)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSecretConfig.
func (in *ServiceSecretConfig) DeepCopy() *ServiceSecretConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceSecretConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceVolumeConfig) DeepCopyInto(out *ServiceVolumeConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceVolumeConfig.
func (in *ServiceVolumeConfig) DeepCopy() *ServiceVolumeConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceVolumeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackSpec) DeepCopyInto(out *StackSpec) {
	*out = *in
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]ServiceConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make(map[string]SecretConfig, len(*in))
		for key, val := range *in {
			newVal := new(SecretConfig)
			val.DeepCopyInto(newVal)
			(*out)[key] = *newVal
		}
	}
	if in.Configs != nil {
		in, out := &in.Configs, &out.Configs
		*out = make(map[string]ConfigObjConfig, len(*in))
		for key, val := range *in {
			newVal := new(ConfigObjConfig)
			val.DeepCopyInto(newVal)
			(*out)[key] = *newVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackSpec.
func (in *StackSpec) DeepCopy() *StackSpec {
	if in == nil {
		return nil
	}
	out := new(StackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateConfig) DeepCopyInto(out *UpdateConfig) {
	*out = *in
	if in.Parallelism != nil {
		in, out := &in.Parallelism, &out.Parallelism
		if *in == nil {
			*out = nil
		} else {
			*out = new(uint64)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateConfig.
func (in *UpdateConfig) DeepCopy() *UpdateConfig {
	if in == nil {
		return nil
	}
	out := new(UpdateConfig)
	in.DeepCopyInto(out)
	return out
}