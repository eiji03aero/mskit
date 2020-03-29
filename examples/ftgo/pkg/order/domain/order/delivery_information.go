package order

import (
	"github.com/eiji03aero/mskit/utils/errbdr"
)

type DeliveryInformation struct {
	Address Address `json:"address"`
}

func (di *DeliveryInformation) Merge(obj interface{}) error {
	switch o := obj.(type) {
	case DeliveryInformation:
		di.Address.Merge(o.Address)
	default:
		return errbdr.NewErrUnknownParams(di.Merge, o)
	}
	return nil
}
