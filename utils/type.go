package utils

import (
	"reflect"
	"runtime"
	"strings"
)

func GetTypeName(v interface{}) (reflect.Type, string) {
	rawType := reflect.TypeOf(v)

	if rawType.Kind() == reflect.Ptr {
		rawType = rawType.Elem()
	}

	name := rawType.String()
	fragments := strings.Split(name, ".")
	return rawType, fragments[1]
}

func GetFunctionName(v interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()
}

func DereferenceIfPtr(v interface{}) interface{} {
	return reflect.Indirect(reflect.ValueOf(v)).Interface()
}
