package order

type CreateOrder struct {
	Id                  string              `json:"id"`
	ConsumerId          string              `json:"consumer_id"`
	RestaurantId        string              `json:"restaurant_id"`
	PaymentInformation  PaymentInformation  `json:"payment_information"`
	DeliveryInformation DeliveryInformation `json:"delivery_information"`
	OrderLineItems      OrderLineItems      `json:"order_line_items"`
}

type RejectOrder struct {
	Id string `json:"id"`
}

type ApproveOrder struct {
	Id string `json:"id"`
}

type ReviseOrder struct {
	Id             string         `json:"id"`
	OrderLineItems OrderLineItems `json:"order_line_items"`
}

type BeginReviseOrder struct {
	Id string `json:"id"`
}

type UndoBeginReviseOrder struct {
	Id string `json:"id"`
}

type HandleTicketCreated struct {
	Id       string `json:"id"`
	TicketId string `json:"ticket_id"`
}

type ConfirmReviseOrder struct {
	Id             string         `json:"id"`
	OrderLineItems OrderLineItems `json:"order_line_items"`
}
