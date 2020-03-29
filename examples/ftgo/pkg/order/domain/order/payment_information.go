package order

import (
	"github.com/eiji03aero/mskit/utils/errbdr"
)

type PaymentInformation struct {
	Token string `json:"token"`
}

func (pi *PaymentInformation) Merge(obj interface{}) error {
	switch o := obj.(type) {
	case PaymentInformation:
		pi.Token = o.Token
	default:
		return errbdr.NewErrUnknownParams(pi.Merge, o)
	}
	return nil
}
