package order

type CreateOrder struct {
	Id                  string              `json:"id"`
	PaymentInformation  PaymentInformation  `json:"payment_information"`
	DeliveryInformation DeliveryInformation `json:"delivery_information"`
	OrderLineItems      OrderLineItems      `json:"order_line_items"`
}
