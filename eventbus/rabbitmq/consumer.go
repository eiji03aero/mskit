package rabbitmq

import (
	"github.com/eiji03aero/mskit/utils"
	"github.com/streadway/amqp"
)

type Consumer struct {
	conn            *amqp.Connection
	channel         *amqp.Channel
	done            chan error
	Tag             string
	ExchangeOption  ExchangeOption
	QueueOption     QueueOption
	QueueBindOption QueueBindOption
	ConsumeOption   ConsumeOption
	DeliveryHandler func(d amqp.Delivery)
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
	AutoAck   bool
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

func (c *Consumer) OnDelivery(cb func(d amqp.Delivery)) *Consumer {
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

	err = ExchangeDeclare(c.channel, c.ExchangeOption)
	if err != nil {
		return
	}

	queue, err := QueueDeclare(c.channel, c.QueueOption)
	if err != nil {
		return
	}

	err = QueueBind(c.channel, queue.Name, c.ExchangeOption.Name, c.QueueBindOption)
	if err != nil {
		return
	}

	deliveries, err := Consume(c.channel, queue.Name, c.Tag, c.ConsumeOption)
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
