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

func NewConsumer(c *amqp.Connection) *Consumer {
	return &Consumer{
		conn:            c,
		done:            make(chan error),
		Tag:             generateDefaultTag(),
		ExchangeOption:  DefaultExchangeOption,
		QueueOption:     DefaultQueueOption,
		QueueBindOption: DefaultQueueBindOption,
		ConsumeOption:   DefaultConsumeOption,
	}
}

func generateDefaultTag() string {
	id, _ := utils.UUID()
	return id
}

func (c *Consumer) Configure(opts ...interface{}) *Consumer {
	for _, opt := range opts {
		switch o := opt.(type) {
		case string:
			// FIXME: cannot force all the string to be Tag. has to be better way :(
			c.Tag = o
		case ExchangeOption:
			c.ExchangeOption = o
		case QueueOption:
			c.QueueOption = o
		case QueueBindOption:
			c.QueueBindOption = o
		case ConsumeOption:
			c.ConsumeOption = o

		case TopicConsumerOption:
			c.ExchangeOption.Type = "topic"
			c.ExchangeOption.Name = o.ExchangeName
			c.QueueBindOption.RoutingKey = o.RoutingKey
		}
	}

	return c
}

func (c *Consumer) OnDelivery(cb func(d amqp.Delivery)) *Consumer {
	c.DeliveryHandler = cb
	return c
}

func (c *Consumer) Exec() (err error) {
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
