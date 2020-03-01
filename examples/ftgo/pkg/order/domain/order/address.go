package order

import (
	"common/errors"
)

type Address struct {
	ZipCode string `json:"zip_code"`
}

func (a *Address) Merge(obj interface{}) error {
	switch o := obj.(type) {
	case Address:
		a.ZipCode = o.ZipCode
	default:
		return errors.ErrNotSupportedParams(a.Merge, o)
	}
	return nil
}
