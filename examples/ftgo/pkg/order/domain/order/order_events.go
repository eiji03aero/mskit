package order

type OrderCreated struct {
	Id                  string              `json:"id"`
	ConsumerId          string              `json:"consumer_id"`
	RestaurantId        string              `json:"restaurant_id"`
	PaymentInformation  PaymentInformation  `json:"payment_information"`
	DeliveryInformation DeliveryInformation `json:"delivery_information"`
	OrderLineItems      OrderLineItems      `json:"order_line_items"`
}
