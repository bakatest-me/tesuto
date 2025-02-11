package jsrepo

import (
	"github.com/dop251/goja"
)

type SafeGet struct {
	v  goja.Value
	vm *goja.Runtime
}

func NewSafeGet(vm *goja.Runtime, key string) SafeGet {
	return SafeGet{
		v:  vm.Get(key),
		vm: vm,
	}
}

func (m SafeGet) String() string {
	if m.v == nil {
		return ""
	}
	return m.v.String()
}

func (m SafeGet) ToBoolean() bool {
	if m.v == nil {
		return false
	}
	return m.v.ToBoolean()
}

func (m SafeGet) Raw() goja.Value {
	return m.v
}

func (m SafeGet) ToObject() saftGetObject {
	if m.v == nil {
		return saftGetObject{}
	}
	return saftGetObject{
		v:  m.v.ToObject(m.vm),
		vm: m.vm,
	}
}

type saftGetObject struct {
	v  *goja.Object
	vm *goja.Runtime
}

func (m saftGetObject) SafeGet(key string) SafeGet {
	return SafeGet{
		v:  m.v.Get(key),
		vm: m.vm,
	}
}

func (m saftGetObject) Raw() *goja.Object {
	return m.v
}

func (m saftGetObject) Keys() []string {
	if m.v == nil {
		return []string{}
	}
	return m.v.Keys()
}
