package order

import (
	"sort"

	"github.com/eiji03aero/mskit/utils/errbdr"
)

type OrderLineItems struct {
	LineItems []OrderLineItem `json:"line_items"`
}

func (oli *OrderLineItems) Merge(obj interface{}) error {
	switch o := obj.(type) {
	case OrderLineItems:
		for _, argOli := range o.LineItems {
			item, ok := oli.GetItemById(argOli.MenuItemId)

			if ok {
				item.Merge(argOli)
			} else {
				oli.LineItems = append(oli.LineItems, OrderLineItem{
					Quantity:   argOli.Quantity,
					MenuItemId: argOli.MenuItemId,
				})
			}
		}
	default:
		return errbdr.NewErrUnknownParams(oli.Merge, o)
	}

	return nil
}

func (oli *OrderLineItems) GetItemById(id string) (item OrderLineItem, found bool) {
	for _, item = range oli.LineItems {
		i := sort.Search(oli.Len(), func(i int) bool {
			return oli.LineItems[i].MenuItemId == id
		})

		if i < oli.Len() {
			return item, true
		}
	}
	return item, false
}

func (oli *OrderLineItems) Len() int {
	return len(oli.LineItems)
}
