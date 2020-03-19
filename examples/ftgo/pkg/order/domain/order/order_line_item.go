package order

import (
	errorscommon "common/errors"
)

type OrderLineItem struct {
	Quantity   int64  `json:"quantity"`
	MenuItemId string `json:"menu_item_id"`
}

func (oli *OrderLineItem) Merge(obj interface{}) error {
	switch o := obj.(type) {
	case OrderLineItem:
		oli.Quantity = o.Quantity
		oli.MenuItemId = o.MenuItemId
	default:
		return errorscommon.NewErrNotSupportedParams(oli.Merge, o)
	}
	return nil
}
