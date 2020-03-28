package service

import (
	logcommon "common/log"
	"fmt"
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

func (s *service) GetOrderTotal(id string) (total int, err error) {
	order, err := s.GetOrder(id)
	if err != nil {
		return
	}

	restaurant, err := s.GetRestaurant(order.RestaurantId)
	if err != nil {
		return
	}

	for _, oli := range order.OrderLineItems.LineItems {
		menuItem, found := restaurant.GetItemById(oli.MenuItemId)
		if !found {
			err = fmt.Errorf(
				"RestaurantMenuItem not found: id=%s restaurantId=%s",
				oli.MenuItemId, restaurant.Id,
			)
			break
		}

		total += menuItem.Price * oli.Quantity
	}

	return
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
