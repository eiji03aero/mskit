package service

import (
	orderdmn "order/domain/order"
	restaurantdmn "order/domain/restaurant"
	restaurantrepo "order/repository/restaurant"

	"github.com/eiji03aero/mskit"
)

type service struct {
	eventRepository      *mskit.EventRepository
	restaurantRepository *restaurantrepo.Repository
}

type Service interface {
	CreateOrder(params orderdmn.CreateOrder) (id string, err error)
	GetOrder(id string) (order *orderdmn.Order, err error)
	CreateRestaurant(restaurantdmn.Restaurant) (err error)
	GetRestaurant(id string) (restaurant *restaurantdmn.Restaurant, err error)
}

func New(
	r *mskit.EventRepository,
	rrepo *restaurantrepo.Repository,
) Service {
	return &service{
		eventRepository:      r,
		restaurantRepository: rrepo,
	}
}
