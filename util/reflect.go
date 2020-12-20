package util

import "reflect"

func EnsureNotPtrReflectValue(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr {
		return reflect.Indirect(v)
	}
	return v
}
