package errbdr

import (
	"fmt"

	"github.com/eiji03aero/mskit/utils"
)

func NewErrNotEnoughPropertiesSet(args [][]interface{}) error {
	var propertiesString string
	for _, arg := range args {
		pairString := fmt.Sprintf(" %s=%v", arg[0], arg[1])
		propertiesString += pairString
	}
	return fmt.Errorf("Not enough properties set: %s", propertiesString)
}

func NewErrUnknownParams(f interface{}, params interface{}) error {
	funcName := utils.GetFunctionName(f)
	typeName := utils.GetTypeName(params)
	return fmt.Errorf("%s: unknown params %s", funcName, typeName)
}

func NewErrDataNotFound(v interface{}, id interface{}) error {
	dataName := utils.GetTypeName(v)
	return fmt.Errorf("Not found: dataName=%s id=%v", dataName, id)
}

func NewErrUnsupportedStateTransition(obj interface{}, state interface{}) error {
	objName := utils.GetTypeName(obj)
	return fmt.Errorf("Unsupported state transition: obj=%s state=%v", objName, state)
}
