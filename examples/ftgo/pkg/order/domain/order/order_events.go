package order

import (
	"github.com/eiji03aero/mskit/facade"
)

func init() {
	facade.RegisterEvent(OrderCreated{})
	facade.RegisterEvent(OrderRejected{})
	facade.RegisterEvent(OrderApproved{})
	facade.RegisterEvent(OrderRevisionBegan{})
	facade.RegisterEvent(UndoOrderRevisionBegan{})
	facade.RegisterEvent(OrderRevisionConfirmed{})
	facade.RegisterEvent(OrderTicketIdSet{})
}

type OrderCreated struct {
	Id                  string              `json:"id"`
	ConsumerId          string              `json:"consumer_id"`
	RestaurantId        string              `json:"restaurant_id"`
	PaymentInformation  PaymentInformation  `json:"payment_information"`
	DeliveryInformation DeliveryInformation `json:"delivery_information"`
	OrderLineItems      OrderLineItems      `json:"order_line_items"`
}

type OrderRejected struct {
	Id string `json:"id"`
}

type OrderApproved struct {
	Id string `json:"id"`
}

type OrderRevisionBegan struct {
	Id string `json:"id"`
}

type UndoOrderRevisionBegan struct {
	Id string `json:"id"`
}

type OrderRevisionConfirmed struct {
	Id             string         `json:"id"`
	OrderLineItems OrderLineItems `json:"order_line_items"`
}

type OrderTicketIdSet struct {
	TicketId string `json:"ticket_id"`
}
