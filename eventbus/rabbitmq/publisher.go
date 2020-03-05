package rabbitmq

import (
	"github.com/streadway/amqp"
)

type Publisher struct {
	conn           *amqp.Connection
	ExchangeOption ExchangeOption
	PublishArgs    PublishArgs
}

func NewPublisher(c *amqp.Connection) *Publisher {
	return &Publisher{
		conn:           c,
		ExchangeOption: DefaultExchangeOption,
		PublishArgs:    DefaultPublishArgs,
	}
}

func (p *Publisher) Configure(opts ...interface{}) *Publisher {
	for _, opt := range opts {
		switch o := opt.(type) {
		case ExchangeOption:
			p.ExchangeOption = o
		case PublishArgs:
			p.PublishArgs = o

		case TopicPublisherOption:
			p.ExchangeOption.Type = "topic"
			p.ExchangeOption.Name = o.ExchangeName
			p.PublishArgs.RoutingKey = o.RoutingKey
			p.PublishArgs.Publishing = o.Publishing
		}
	}

	return p
}

func (p *Publisher) Exec() error {
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
