package service

import (
	"fmt"
	orderdmn "order/domain/order"
	"order/saga/createorder"
	"order/saga/reviseorder"

	"github.com/eiji03aero/mskit/utils"
	"github.com/eiji03aero/mskit/utils/logger"
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

	logger.PrintResourceCreated(order)
	return
}

func (s *service) GetOrder(id string) (*orderdmn.Order, error) {
	order := &orderdmn.Order{}
	err := s.eventRepository.Load(id, order)

	logger.PrintResourceGet(order)
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
	order, err := s.GetOrder(cmd.Id)
	if err != nil {
		return
	}

	err = s.eventRepository.ExecuteCommand(order, cmd)
	if err != nil {
		return
	}

	logger.PrintResource(order, "rejected")
	return
}

func (s *service) ApproveOrder(cmd orderdmn.ApproveOrder) (err error) {
	order, err := s.GetOrder(cmd.Id)
	if err != nil {
		return
	}

	err = s.eventRepository.ExecuteCommand(order, cmd)
	if err != nil {
		return
	}

	logger.PrintResource(order, "confirmed")
	return
}

func (s *service) ReviseOrder(cmd orderdmn.ReviseOrder) (err error) {
	order, err := s.GetOrder(cmd.Id)
	if err != nil {
		return
	}

	sagaState := reviseorder.NewState(order.Id, order.OrderLineItems)
	go s.reviseOrderSagaManager.Create(sagaState)

	logger.PrintResource(order, "revise")
	return
}

func (s *service) BeginReviseOrder(cmd orderdmn.BeginReviseOrder) (err error) {
	order, err := s.GetOrder(cmd.Id)
	if err != nil {
		return
	}

	err = s.eventRepository.ExecuteCommand(order, cmd)
	if err != nil {
		return
	}

	logger.PrintResource(order, "revise began")
	return
}

func (s *service) UndoBeginReviseOrder(cmd orderdmn.UndoBeginReviseOrder) (err error) {
	order, err := s.GetOrder(cmd.Id)
	if err != nil {
		return
	}

	err = s.eventRepository.ExecuteCommand(order, cmd)
	if err != nil {
		return
	}

	logger.PrintResource(order, "undo revise began")
	return
}

func (s *service) ConfirmReviseOrder(cmd orderdmn.ConfirmReviseOrder) (err error) {
	order, err := s.GetOrder(cmd.Id)
	if err != nil {
		return
	}

	err = s.eventRepository.ExecuteCommand(order, cmd)
	if err != nil {
		return
	}

	logger.PrintResource(order, "confirm revise order")
	return
}

func (s *service) HandleTicketCreated(cmd orderdmn.HandleTicketCreated) (err error) {
	order, err := s.GetOrder(cmd.Id)
	if err != nil {
		return
	}

	err = s.eventRepository.ExecuteCommand(order, cmd)
	if err != nil {
		return
	}

	logger.PrintResource(order, "handle ticket created")
	return
}
