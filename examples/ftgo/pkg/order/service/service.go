package service

import (
	logcommon "common/log"
	"log"

	orderdmn "order/domain/order"
	"order/pb"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/utils"
)

type service struct {
	repository *mskit.Repository
}

type Service interface {
	CreateOrder(params pb.CreateOrder) (id string, err error)
	GetOrder(id string) (order *orderdmn.Order, err error)
}

func New(r *mskit.Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateOrder(params pb.CreateOrder) (string, error) {
	id, err := utils.UUID()
	if err != nil {
		return "", err
	}

	order := &orderdmn.Order{}
	params.Id = id

	events, err := order.Process(params)
	if err != nil {
		return "", err
	}

	for _, e := range events {
		err = s.repository.Save(order, e)
		if err != nil {
			return "", err
		}
	}

	log.Println("order created: ")
	logcommon.PrintJsonln(order)
	return id, nil
}

func (s *service) GetOrder(id string) (*orderdmn.Order, error) {
	order := &orderdmn.Order{}
	err := s.repository.Load(id, order)

	log.Println("get order: ")
	logcommon.PrintJsonln(order)
	return order, err
}
