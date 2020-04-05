package logger

import (
	"log"

	"github.com/eiji03aero/mskit/utils"
)

func Println(
	args ...interface{},
) {
	args = formatArgsRest(
		[]interface{}{},
		args,
	)

	log.Println(args...)
}

func PrintFail(
	msg string,
	args ...interface{},
) {
	args = formatArgsRest(
		[]interface{}{
			RedString(msg),
		},
		args,
	)

	log.Println(args...)
}

func PrintFuncCall(
	f interface{},
	rest ...interface{},
) {
	fname := utils.GetFunctionNameParent(f)
	args := formatArgsRest(
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
	resourceName := utils.GetTypeName(resource)
	args := formatArgsRest(
		[]interface{}{
			CyanString(resourceName),
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
	args := formatArgsRest(
		[]interface{}{
			CyanString(resourceName),
			HiBlueString("created"),
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
	args := formatArgsRest(
		[]interface{}{
			CyanString(resourceName),
			HiBlueString("get"),
			resource,
		},
		rest,
	)

	log.Println(args...)
}
