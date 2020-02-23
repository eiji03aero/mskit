package order

import (
	errorscommon "common/errors"
	"order/pb"
)

type OrderLineItem struct {
	Quantity   int64  `json:"quantity"`
	MenuItemId string `json:"menu_item_id"`
}

func (oli *OrderLineItem) Merge(obj interface{}) error {
	switch o := obj.(type) {
	case *pb.OrderLineItem:
		oli.Quantity = o.Quantity
		oli.MenuItemId = o.MenuItemId
	default:
		return errorscommon.ErrNotSupportedParams(oli.Merge, o)
	}
	return nil
}
