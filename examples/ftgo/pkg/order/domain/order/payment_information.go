package order

import (
	"common/errors"
)

type PaymentInformation struct {
	Token string `json:"token"`
}

func (pi *PaymentInformation) Merge(obj interface{}) error {
	switch o := obj.(type) {
	case PaymentInformation:
		pi.Token = o.Token
	default:
		return errors.ErrNotSupportedParams(pi.Merge, o)
	}
	return nil
}
