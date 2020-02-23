package order

import (
	"common/errors"
	"order/pb"
)

type DeliveryInformation struct {
	Address Address `json:"address"`
}

func (di *DeliveryInformation) Merge(obj interface{}) error {
	switch o := obj.(type) {
	case *pb.DeliveryInformation:
		di.Address.Merge(o.Address)
	default:
		return errors.ErrNotSupportedParams(di.Merge, o)
	}
	return nil
}
