package mvc2

import "reflect"

func isContext(inTyp reflect.Type) bool {
	return inTyp.String() == "context.Context" // I couldn't find another way; context/context.go is not exported.
}

func indirectVal(v reflect.Value) reflect.Value {
	return reflect.Indirect(v)
}

func indirectTyp(typ reflect.Type) reflect.Type {
	switch typ.Kind() {
	case reflect.Ptr, reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return typ.Elem()
	}
	return typ
}

func goodVal(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		if v.IsNil() {
			return false
		}
	}

	return v.IsValid()
}

func isFunc(typ reflect.Type) bool {
	return typ.Kind() == reflect.Func
}