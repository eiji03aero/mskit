package errors

import "fmt"

func ErrNotSupportedParams(f interface{}, v interface{}) error {
	return fmt.Errorf("Not supported params: function=%v value=%v", f, v)
}
