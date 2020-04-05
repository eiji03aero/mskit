package order

import (
	orderdmn "order/domain/order"
)

type OrderProxy interface {
	RejectOrder(id string) (err error)
	ApproveOrder(id string) (err error)
	BeginReviseOrder(id string) (err error)
	UndoBeginReviseOrder(id string) (err error)
	ConfirmReviseOrder(
		id string,
		orderLineItems orderdmn.OrderLineItems,
	) (err error)
}

type ConsumerProxy interface {
	ValidateOrder(orderId string, total int) (err error)
}

type KitchenProxy interface {
	CreateTicket(
		restaurantId string,
		orderLineItems orderdmn.OrderLineItems,
	) (ticketId string, err error)
	CancelTicket(id string) (err error)
	ConfirmTicket(id string) (err error)
	BeginReviseTicket(
		id string,
		orderLineItems orderdmn.OrderLineItems,
	) (err error)
	UndoBeginReviseTicket(id string) (err error)
	ConfirmReviseTicket(
		id string,
		orderLineItems orderdmn.OrderLineItems,
	) (err error)
}

type AccountingProxy interface {
	Authorize(consumerId string) (err error)
}
