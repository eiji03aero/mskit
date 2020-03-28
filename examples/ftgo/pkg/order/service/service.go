package service

import (
	"order"
	restaurantrepo "order/repository/restaurant"

	"github.com/eiji03aero/mskit"
)

type service struct {
	eventRepository      *mskit.EventRepository
	restaurantRepository *restaurantrepo.Repository

	createOrderSagaManager mskit.SagaManager
}

func New(
	r *mskit.EventRepository,
	rrepo *restaurantrepo.Repository,
) order.Service {
	return &service{
		eventRepository:      r,
		restaurantRepository: rrepo,
	}
}

func (s *service) InjectSagaManagers(
	createOrderSagaManager mskit.SagaManager,
) {
	s.createOrderSagaManager = createOrderSagaManager
}
