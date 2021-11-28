package mate

import (
	"errors"
	"reflect"
)

func parsePoolFunc(f interface{}) (outType reflect.Type, e error) {
	fType := reflect.TypeOf(f)
	if fType.Kind() != reflect.Func {
		e = errors.New("it's not a func")
		return
	}
	if fType.NumOut() != 1 {
		e = errors.New("return must be an object pointer")
		return
	}
	outType = fType.Out(0)
	if outType.Kind() != reflect.Ptr {
		e = errors.New("return must be an object pointer")
		return
	}
	return
}
func allFields(dest interface{}, call func(reflect.Value)) {
	destVal := indirect(reflect.ValueOf(dest))
	destType := destVal.Type()
	if destType.Kind() != reflect.Struct && destType.Kind() != reflect.Interface {
		return
	}

	for index := 0; index < destVal.NumField(); index++ {
		if destType.Field(index).Anonymous {
			allFields(destVal.Field(index).Addr().Interface(), call)
			continue
		}
		val := destVal.Field(index)
		kind := val.Kind()
		if kind != reflect.Ptr && kind != reflect.Interface {
			continue
		}
		call(val)
	}
}

func indirect(reflectValue reflect.Value) reflect.Value {
	for reflectValue.Kind() == reflect.Ptr || reflectValue.Kind() == reflect.Interface {
		reflectValue = reflectValue.Elem()
	}
	return reflectValue
}