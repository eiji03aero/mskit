package order

import (
	"github.com/eiji03aero/mskit/utils/errbdr"
)

type Address struct {
	ZipCode string `json:"zip_code"`
}

func (a *Address) Merge(obj interface{}) error {
	switch o := obj.(type) {
	case Address:
		a.ZipCode = o.ZipCode
	default:
		return errbdr.NewErrUnknownParams(a.Merge, o)
	}
	return nil
}
