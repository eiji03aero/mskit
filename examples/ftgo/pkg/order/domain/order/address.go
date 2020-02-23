package order

import (
	"common/errors"
	"order/pb"
)

type Address struct {
	ZipCode string `json:"zip_code"`
}

func (a *Address) Merge(obj interface{}) error {
	switch o := obj.(type) {
	case *pb.Address:
		a.ZipCode = o.ZipCode
	default:
		return errors.ErrNotSupportedParams(a.Merge, o)
	}
	return nil
}
