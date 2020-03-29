package logger

import (
	"encoding/json"
	"log"

	"github.com/eiji03aero/mskit/utils"
	"github.com/fatih/color"
)

func Println(
	args ...interface{},
) {
	args = formatData(
		[]interface{}{},
		args,
	)

	log.Println(args...)
}

func PrintFuncCall(
	f interface{},
	rest ...interface{},
) {
	fname := utils.GetFunctionNameParent(f)
	args := formatData(
		[]interface{}{
			CyanString(fname),
		},
		rest,
	)

	log.Println(args...)
}

func PrintResource(
	resource interface{},
	rest ...interface{},
) {
	args := formatData(
		[]interface{}{
			resource,
		},
		rest,
	)

	log.Println(args...)
}

func PrintResourceCreated(
	resource interface{},
	rest ...interface{},
) {
	resourceName := utils.GetTypeName(resource)
	args := formatData(
		[]interface{}{
			CyanString(resourceName),
			BlueString("created"),
			resource,
		},
		rest,
	)

	log.Println(args...)
}

func PrintResourceGet(
	resource interface{},
	rest ...interface{},
) {
	resourceName := utils.GetTypeName(resource)
	args := formatData(
		[]interface{}{
			CyanString(resourceName),
			BlueString("get"),
			resource,
		},
		rest,
	)

	log.Println(args...)
}

func formatData(args []interface{}, rest []interface{}) (result []interface{}) {
	args = append(args, rest...)
	return formatDataToJson(args)
}

func formatDataToJson(args []interface{}) (result []interface{}) {
	for _, arg := range args {
		var aresult interface{}

		switch a := arg.(type) {
		case string:
			aresult = a
		case []byte:
			aresult = string(a)
		default:
			aJson, err := json.Marshal(a)
			if err != nil {
				panic(err)
			}
			aresult = string(aJson)
		}

		result = append(result, aresult)
	}

	return
}

func RedString(s string, rest ...interface{}) string {
	return color.RedString(s, rest...)
}

func BlueString(s string, rest ...interface{}) string {
	return color.BlueString(s, rest...)
}

func CyanString(s string, rest ...interface{}) string {
	return color.CyanString(s, rest...)
}
