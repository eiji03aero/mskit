package rabbitmq

import (
	"github.com/eiji03aero/mskit/utils"
	"github.com/streadway/amqp"
)

type Delivery amqp.Delivery

type Consumer struct {
	conn            *amqp.Connection
	channel         *amqp.Channel
	done            chan error
	Tag             string
	ExchangeOption  ExchangeOption
	QueueOption     QueueOption
	QueueBindOption QueueBindOption
	ConsumeOption   ConsumeOption
	DeliveryHandler DeliveryHandler
}

type QueueOption struct {
	Name        string
	Durable     bool
	AutoDeleted bool
	Exclusive   bool
	NoWait      bool
	Arguments   amqp.Table
}

type QueueBindOption struct {
	Key       string
	NoWait    bool
	Arguments amqp.Table
}

type ConsumeOption struct {
	NoAck     bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Arguments amqp.Table
}

func NewConsumer(c *amqp.Connection) *Consumer {
	return &Consumer{
		conn: c,
		done: make(chan error),
	}
}

func (c *Consumer) Configure(
	tag string,
	exopt ExchangeOption,
	qopt QueueOption,
	qbopt QueueBindOption,
	copt ConsumeOption,
) *Consumer {
	c.Tag = tag
	c.ExchangeOption = exopt
	c.QueueOption = qopt
	c.QueueBindOption = qbopt
	c.ConsumeOption = copt
	return c
}

type DeliveryHandler func(msg amqp.Delivery)

func (c *Consumer) OnDelivery(cb DeliveryHandler) *Consumer {
	c.DeliveryHandler = cb
	return c
}

func (c *Consumer) Exec() (err error) {
	if c.Tag == "" ||
		&(c.ExchangeOption) == nil ||
		&(c.QueueOption) == nil ||
		&(c.QueueBindOption) == nil ||
		&(c.ConsumeOption) == nil ||
		&(c.DeliveryHandler) == nil {
		return utils.NewErrNotEnoughPropertiesSet([][]interface{}{
			{"Tag", c.Tag},
			{"ExchangeOption", c.ExchangeOption},
			{"QueueOption", c.QueueOption},
			{"QueueBindOption", c.QueueBindOption},
			{"ConsumeOption", c.ConsumeOption},
			{"DeliveryHandler", c.DeliveryHandler},
		})
	}

	c.channel, err = c.conn.Channel()
	if err != nil {
		return
	}

	err = c.channel.ExchangeDeclare(
		c.ExchangeOption.Name,
		c.ExchangeOption.Type,
		c.ExchangeOption.Durable,
		c.ExchangeOption.AutoDeleted,
		c.ExchangeOption.Internal,
		c.ExchangeOption.NoWait,
		c.ExchangeOption.Arguments,
	)
	if err != nil {
		return
	}

	queue, err := c.channel.QueueDeclare(
		c.QueueOption.Name,
		c.QueueOption.Durable,
		c.QueueOption.AutoDeleted,
		c.QueueOption.Exclusive,
		c.QueueOption.NoWait,
		c.QueueOption.Arguments,
	)
	if err != nil {
		return
	}

	err = c.channel.QueueBind(
		queue.Name,
		c.QueueBindOption.Key,
		c.ExchangeOption.Name,
		c.QueueBindOption.NoWait,
		c.QueueBindOption.Arguments,
	)
	if err != nil {
		return
	}

	deliveries, err := c.channel.Consume(
		queue.Name,
		c.Tag,
		c.ConsumeOption.NoAck,
		c.ConsumeOption.Exclusive,
		c.ConsumeOption.NoLocal,
		c.ConsumeOption.NoWait,
		c.ConsumeOption.Arguments,
	)
	if err != nil {
		return
	}

	go c.handleDelivery(deliveries)

	return
}

func (c *Consumer) handleDelivery(deliveries <-chan amqp.Delivery) {
	for d := range deliveries {
		c.DeliveryHandler(d)
	}
	c.done <- nil
}
