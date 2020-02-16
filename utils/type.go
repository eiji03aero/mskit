package utils

import (
	"reflect"
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
