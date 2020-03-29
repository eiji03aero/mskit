package order

import (
	"github.com/eiji03aero/mskit/utils/errbdr"
)

type OrderLineItem struct {
	Quantity   int    `json:"quantity"`
	MenuItemId string `json:"menu_item_id"`
}

func (oli *OrderLineItem) Merge(obj interface{}) error {
	switch o := obj.(type) {
	case OrderLineItem:
		oli.Quantity = o.Quantity
		oli.MenuItemId = o.MenuItemId
	default:
		return errbdr.NewErrUnknownParams(oli.Merge, o)
	}
	return nil
}
