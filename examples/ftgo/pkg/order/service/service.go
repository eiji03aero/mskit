package service

import (
	logcommon "common/log"
	orderdmn "order/domain/order"
	restaurantdmn "order/domain/restaurant"
	restaurantrepo "order/repository/restaurant"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/utils"
)

type service struct {
	repository           *mskit.Repository
	restaurantRepository *restaurantrepo.Repository
}

type Service interface {
	CreateOrder(params orderdmn.CreateOrder) (id string, err error)
	GetOrder(id string) (order *orderdmn.Order, err error)
	CreateRestaurant(restaurantdmn.Restaurant) (err error)
	GetRestaurant(id string) (restaurant *restaurantdmn.Restaurant, err error)
}

func New(
	r *mskit.Repository,
	rrepo *restaurantrepo.Repository,
) Service {
	return &service{
		repository:           r,
		restaurantRepository: rrepo,
	}
}

func (s *service) CreateOrder(params orderdmn.CreateOrder) (string, error) {
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

	logcommon.PrintCreated(order)
	return id, nil
}

func (s *service) GetOrder(id string) (*orderdmn.Order, error) {
	order := &orderdmn.Order{}
	err := s.repository.Load(id, order)

	logcommon.PrintGet(order)
	return order, err
}
