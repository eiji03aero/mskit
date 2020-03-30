package order

import (
	orderdmn "order/domain/order"
)

type OrderProxy interface {
	RejectOrder(id string) (err error)
	ApproveOrder(id string) (err error)
}

type ConsumerProxy interface {
	ValidateOrder(orderId string, total int) (err error)
}

type KitchenProxy interface {
	CreateTicket(
		restaurantId string,
		lineItems []orderdmn.OrderLineItem,
	) (ticketId string, err error)
	CancelTicket(id string) (err error)
	ConfirmTicket(id string) (err error)
}

type AccountingProxy interface {
	Authorize(consumerId string) (err error)
}
