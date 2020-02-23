package order

import (
	errorscommon "common/errors"
	"order/pb"
	"sort"
)

type OrderLineItems struct {
	LineItems []OrderLineItem `json:"line_items"`
}

func (oli *OrderLineItems) Merge(obj interface{}) error {
	switch o := obj.(type) {
	case *pb.OrderLineItems:
		for _, argOli := range o.LineItems {
			i := sort.Search(len(oli.LineItems), func(i int) bool {
				return oli.LineItems[i].MenuItemId == argOli.MenuItemId
			})

			if i < len(oli.LineItems) {
				oli.LineItems[i].Merge(argOli)
			} else {
				oli.LineItems = append(oli.LineItems, OrderLineItem{
					Quantity:   argOli.Quantity,
					MenuItemId: argOli.MenuItemId,
				})
			}
		}
	default:
		return errorscommon.ErrNotSupportedParams(oli.Merge, o)
	}

	return nil
}

func (oli *OrderLineItems) Len() int {
	return len(oli.LineItems)
}
