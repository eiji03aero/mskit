package rabbitmq

import (
	"github.com/eiji03aero/mskit/utils"
	"github.com/streadway/amqp"
)

type Publishing amqp.Publishing

type Publisher struct {
	conn           *amqp.Connection
	ExchangeOption ExchangeOption
	PublishArgs    PublishArgs
}

type PublishArgs struct {
	RoutingKey string
	Mandatory  bool
	Immediate  bool
	Publishing amqp.Publishing
}

func NewPublisher(c *amqp.Connection) *Publisher {
	return &Publisher{
		conn: c,
	}
}

func (p *Publisher) Configure(exopt ExchangeOption, pubargs PublishArgs) *Publisher {
	p.ExchangeOption = exopt
	p.PublishArgs = pubargs
	return p
}

func (p *Publisher) Exec() error {
	if &(p.ExchangeOption) == nil || &(p.PublishArgs) == nil {
		return utils.NewErrNotEnoughPropertiesSet([][]interface{}{
			{"ExchangeOption", p.ExchangeOption},
			{"PublishArgs", p.PublishArgs},
		})
	}

	channel, err := p.conn.Channel()
	if err != nil {
		return err
	}

	err = ExchangeDeclare(channel, p.ExchangeOption)
	if err != nil {
		return err
	}

	err = Publish(channel, p.ExchangeOption.Name, p.PublishArgs)
	if err != nil {
		return err
	}

	return nil
}
