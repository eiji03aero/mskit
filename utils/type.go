package utils

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

func GetType(v interface{}) (reflect.Type, string) {
	rawType := reflect.TypeOf(v)

	if rawType.Kind() == reflect.Ptr {
		rawType = rawType.Elem()
	}

	name := rawType.String()
	fragments := strings.Split(name, ".")
	return rawType, fragments[1]
}

func GetTypeName(v interface{}) (name string) {
	_, name = GetType(v)
	return
}

func GetFunctionNameFull(v interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()
}

func GetFunctionName(v interface{}) string {
	name := GetFunctionNameFull(v)
	fragments := strings.Split(name, ".")
	return fragments[len(fragments)-1]
}

func GetFunctionNameParent(v interface{}) string {
	name := GetFunctionNameFull(v)
	fragments := strings.Split(name, ".")
	if len(fragments) == 1 {
		return fragments[0]
	}

	parent := fragments[len(fragments)-2]
	fname := fragments[len(fragments)-1]
	return fmt.Sprintf("%s.%s", parent, fname)
}

func DereferenceIfPtr(v interface{}) interface{} {
	return reflect.Indirect(reflect.ValueOf(v)).Interface()
}
