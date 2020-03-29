package service

import (
	consumerdmn "consumer/domain/consumer"

	"github.com/eiji03aero/mskit/utils"
	"github.com/eiji03aero/mskit/utils/logger"
)

func (s *service) CreateConsumer(cmd consumerdmn.CreateConsumer) (id string, err error) {
	id, err = utils.UUID()
	if err != nil {
		return
	}

	consumerAggregate := consumerdmn.NewConsumerAggregate()
	cmd.Id = id

	s.eventRepository.ExecuteCommand(consumerAggregate, cmd)
	if err != nil {
		return
	}

	logger.PrintResourceCreated(consumerAggregate)
	return
}

func (s *service) GetConsumer(id string) (consumer *consumerdmn.Consumer, err error) {
	consumerAggregate := consumerdmn.NewConsumerAggregate()
	err = s.eventRepository.Load(id, consumerAggregate)
	if err != nil {
		return
	}

	logger.PrintResourceGet(consumerAggregate)
	consumer = consumerAggregate.Consumer
	return consumer, nil
}

func (s *service) ValidateOrder(cmd consumerdmn.ValidateOrder) (err error) {
	consumerAggregate := consumerdmn.NewConsumerAggregate()
	err = s.eventRepository.Load(cmd.Id, consumerAggregate)
	if err != nil {
		return
	}

	err = consumerAggregate.ValidateOrder(cmd.Total)

	return
}
