package service

import (
	logcommon "common/log"
	orderdmn "order/domain/order"
	"order/saga/createorder"

	"github.com/eiji03aero/mskit/utils"
)

func (s *service) CreateOrder(cmd orderdmn.CreateOrder) (id string, err error) {
	restaurant, err := s.GetRestaurant(cmd.RestaurantId)
	if err != nil {
		return
	}

	err = s.validateMenuItems(restaurant, cmd.OrderLineItems)
	if err != nil {
		return
	}

	id, err = utils.UUID()
	if err != nil {
		return
	}

	cmd.Id = id
	order := &orderdmn.Order{}

	err = s.eventRepository.ExecuteCommand(order, cmd)
	if err != nil {
		return
	}

	sagaState := createorder.NewState(order.Id)
	go s.createOrderSagaManager.Create(sagaState)

	logcommon.PrintCreated(order)
	return
}

func (s *service) GetOrder(id string) (*orderdmn.Order, error) {
	order := &orderdmn.Order{}
	err := s.eventRepository.Load(id, order)

	logcommon.PrintGet(order)
	return order, err
}

func (s *service) RejectOrder(cmd orderdmn.RejectOrder) (err error) {
	order := &orderdmn.Order{}
	err = s.eventRepository.Load(cmd.Id, order)
	if err != nil {
		return
	}

	err = s.eventRepository.ExecuteCommand(order, cmd)
	if err != nil {
		return
	}

	logcommon.PrintlnWithJson("order rejected: ", order)
	return
}
