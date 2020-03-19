package errors

import (
	"fmt"
	"github.com/eiji03aero/mskit/utils"
)

func NewErrNotSupportedParams(f interface{}, v interface{}) error {
	funcName := utils.GetFunctionName(f)
	return fmt.Errorf("Not supported params: function=%s value=%v", funcName, v)
}

func NewErrDataNotFound(v interface{}, id interface{}) error {
	_, dataName := utils.GetTypeName(v)
	return fmt.Errorf("Not found: dataName=%s id=%v", dataName, id)
}
