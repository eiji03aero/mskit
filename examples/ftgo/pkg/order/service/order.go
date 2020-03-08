package service

import (
	logcommon "common/log"
	orderdmn "order/domain/order"

	"github.com/eiji03aero/mskit/utils"
)

func (s *service) CreateOrder(cmd orderdmn.CreateOrder) (string, error) {
	id, err := utils.UUID()
	if err != nil {
		return "", err
	}

	order := &orderdmn.Order{}
	cmd.Id = id

	err = s.repository.ExecuteCommand(order, cmd)
	if err != nil {
		return "", err
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
