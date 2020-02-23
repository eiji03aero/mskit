package errors

import "fmt"

func ErrNotSupportedParams(f interface{}, v interface{}) error {
	return fmt.Errorf("Not supported params: %v %v", f, v)
}
