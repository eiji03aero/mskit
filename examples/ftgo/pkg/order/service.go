package order

type Service interface {
	CreateOrder(params *CreateOrderParams) (id string, err error)
}

type CreateOrderParams struct {
	Name string `json:"name"`
}
