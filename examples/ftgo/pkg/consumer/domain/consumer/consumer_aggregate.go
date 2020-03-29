package consumer

import (
	"fmt"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/utils/errbdr"
)

type ConsumerAggregate struct {
	mskit.BaseAggregate
	Consumer *Consumer `json:"consumer"`
}

func NewConsumerAggregate() *ConsumerAggregate {
	return &ConsumerAggregate{
		Consumer: &Consumer{},
	}
}

func (ca *ConsumerAggregate) ValidateOrder(total int) error {
	if total < 20 {
		return fmt.Errorf("total too little")
	}
	return nil
}

func (c *ConsumerAggregate) Validate() (errs []error) {
	return
}

func (c *ConsumerAggregate) Process(command interface{}) (mskit.Events, error) {
	switch cmd := command.(type) {
	case CreateConsumer:
		return c.processCreateConsumer(cmd)
	default:
		return mskit.Events{}, errbdr.NewErrUnknownParams(c.Process, cmd)
	}
}

func (c *ConsumerAggregate) processCreateConsumer(cmd CreateConsumer) (mskit.Events, error) {
	events := mskit.NewEventsSingle(
		cmd.Id,
		ConsumerAggregate{},
		ConsumerCreated{
			Id:   cmd.Id,
			Name: cmd.Name,
		},
	)

	return events, nil
}

func (c *ConsumerAggregate) Apply(event interface{}) error {
	switch e := event.(type) {
	case ConsumerCreated:
		return c.applyConsumerCreated(e)
	default:
		return errbdr.NewErrUnknownParams(c.Apply, e)
	}
}

func (c *ConsumerAggregate) applyConsumerCreated(event ConsumerCreated) error {
	c.Id = event.Id
	c.Consumer.Id = event.Id
	c.Consumer.Name = event.Name

	return nil
}
