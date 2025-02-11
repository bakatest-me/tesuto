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

func (m SafeGet) Boolean() bool {
	if m.v == nil {
		return false
	}
	return m.v.ToBoolean()
}

func (m SafeGet) Raw() goja.Value {
	return m.v
}

func (m SafeGet) Object() safeGetObject {
	if m.v == nil {
		return safeGetObject{}
	}
	return safeGetObject{
		v:  m.v.ToObject(m.vm),
		vm: m.vm,
	}
}

type safeGetObject struct {
	v  *goja.Object
	vm *goja.Runtime
}

func (m safeGetObject) SafeGet(key string) SafeGet {
	return SafeGet{
		v:  m.v.Get(key),
		vm: m.vm,
	}
}

func (m safeGetObject) Raw() *goja.Object {
	return m.v
}

func (m safeGetObject) Keys() []string {
	if m.v == nil {
		return []string{}
	}
	return m.v.Keys()
}
