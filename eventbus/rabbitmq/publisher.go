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

	err = channel.ExchangeDeclare(
		p.ExchangeOption.Name,
		p.ExchangeOption.Type,
		p.ExchangeOption.Durable,
		p.ExchangeOption.AutoDeleted,
		p.ExchangeOption.Internal,
		p.ExchangeOption.NoWait,
		p.ExchangeOption.Arguments,
	)
	if err != nil {
		return err
	}

	err = channel.Publish(
		p.ExchangeOption.Name,
		p.PublishArgs.RoutingKey,
		p.PublishArgs.Mandatory,
		p.PublishArgs.Immediate,
		p.PublishArgs.Publishing,
	)
	if err != nil {
		return err
	}

	return nil
}
