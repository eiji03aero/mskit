package service

import (
	logcommon "common/log"
	orderdmn "order/domain/order"

	"github.com/eiji03aero/mskit/utils"
)

func (s *service) CreateOrder(cmd orderdmn.CreateOrder) (id string, err error) {
	_, err = s.GetRestaurant(cmd.RestaurantId)
	if err != nil {
		return
	}

	id, err = utils.UUID()
	if err != nil {
		return "", err
	}

	cmd.Id = id
	order := &orderdmn.Order{}

	err = s.eventRepository.ExecuteCommand(order, cmd)
	if err != nil {
		return "", err
	}

	logcommon.PrintCreated(order)
	return id, nil
}

func (s *service) GetOrder(id string) (*orderdmn.Order, error) {
	order := &orderdmn.Order{}
	err := s.eventRepository.Load(id, order)

	logcommon.PrintGet(order)
	return order, err
}
