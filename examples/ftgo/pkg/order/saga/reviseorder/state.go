package reviseorder

import (
	orderdmn "order/domain/order"

	"github.com/eiji03aero/mskit/utils/errbdr"
)

type state struct {
	OrderId        string                  `json:"order_id"`
	OrderLineItems orderdmn.OrderLineItems `json:"order_line_items"`
}

func NewState(orderId string, items orderdmn.OrderLineItems) *state {
	return &state{
		OrderId:        orderId,
		OrderLineItems: items,
	}
}

func assertStruct(value interface{}) (s *state, err error) {
	s, ok := value.(*state)
	if !ok {
		err = errbdr.NewErrUnknownParams(assertStruct, s)
		return
	}

	return
}
