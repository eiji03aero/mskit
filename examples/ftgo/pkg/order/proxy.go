package order

type OrderProxy interface {
	RejectOrder(id string) (err error)
}

type ConsumerProxy interface {
	ValidateOrder(orderId string, total int) (err error)
}
