//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package core

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *App) DeepCopyInto(out *App) {
	*out = *in
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
func (in *ComponentSpec) DeepCopyInto(out *ComponentSpec) {
	*out = *in
	out.ComponentTypeVar = in.ComponentTypeVar
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]RouteSpec, len(*in))
		copy(*out, *in)
	}
	if in.EnvSchema != nil {
		in, out := &in.EnvSchema, &out.EnvSchema
		*out = make(map[string]*EnvVarSchema, len(*in))
		for key, val := range *in {
			var outVal *EnvVarSchema
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(EnvVarSchema)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make(map[string]*ComponentTypeVar, len(*in))
		for key, val := range *in {
			var outVal *ComponentTypeVar
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(ComponentTypeVar)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentSpec.
func (in *ComponentSpec) DeepCopy() *ComponentSpec {
	if in == nil {
		return nil
	}
	out := new(ComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentTypeVar) DeepCopyInto(out *ComponentTypeVar) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentTypeVar.
func (in *ComponentTypeVar) DeepCopy() *ComponentTypeVar {
	if in == nil {
		return nil
	}
	out := new(ComponentTypeVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVarSchema) DeepCopyInto(out *EnvVarSchema) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVarSchema.
func (in *EnvVarSchema) DeepCopy() *EnvVarSchema {
	if in == nil {
		return nil
	}
	out := new(EnvVarSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteSpec) DeepCopyInto(out *RouteSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSpec.
func (in *RouteSpec) DeepCopy() *RouteSpec {
	if in == nil {
		return nil
	}
	out := new(RouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Val) DeepCopyInto(out *Val) {
	*out = *in
	if in.arrayNumVal != nil {
		in, out := &in.arrayNumVal, &out.arrayNumVal
		*out = make([]float64, len(*in))
		copy(*out, *in)
	}
	if in.arrayStrVal != nil {
		in, out := &in.arrayStrVal, &out.arrayStrVal
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Val.
func (in *Val) DeepCopy() *Val {
	if in == nil {
		return nil
	}
	out := new(Val)
	in.DeepCopyInto(out)
	return out
}
