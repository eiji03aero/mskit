package service

import (
	consumerdmn "consumer/domain/consumer"

	"github.com/eiji03aero/mskit/utils"
)

func (s *service) CreateConsumer(cmd consumerdmn.CreateConsumer) (id string, err error) {
	id, err = utils.UUID()
	if err != nil {
		return
	}

	consumerAggregate := consumerdmn.NewConsumerAggregate()
	cmd.Id = id

	err = s.eventRepository.ExecuteCommand(consumerAggregate, cmd)
	if err != nil {
		return
	}

	return
}

func (s *service) GetConsumer(id string) (consumer *consumerdmn.Consumer, err error) {
	consumerAggregate := consumerdmn.NewConsumerAggregate()
	err = s.eventRepository.Load(id, consumerAggregate)
	if err != nil {
		return
	}

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
