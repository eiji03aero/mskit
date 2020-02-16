package service

import (
	"github.com/eiji03aero/mskit"
	ordersvc "github.com/eiji03aero/mskit/examples/ftgo/pkg/order"
	orderdomain "github.com/eiji03aero/mskit/examples/ftgo/pkg/order/domain/order"
	"github.com/eiji03aero/mskit/utils"
)

type service struct {
	repository mskit.Repository
}

func New(repo mskit.Repository) ordersvc.Service {
	return &service{
		repository: repo,
	}
}

func (s *service) CreateOrder(params *ordersvc.CreateOrderParams) (string, error) {
	id, err := utils.UUID()
	if err != nil {
		return "", err
	}

	order := &orderdomain.Order{}
	createOrder := &orderdomain.CreateOrder{
		ID:   id,
		Name: params.Name,
	}

	events, err := order.Process(createOrder)
	if err != nil {
		return "", err
	}

	for _, e := range events {
		order.Apply(e)
		err := s.repository.Save(e)
		if err != nil {
			return "", err
		}
	}

	return id, nil
}
