package rabbitmq

import (
	"github.com/streadway/amqp"
)

type QueueOption struct {
	Name        string
	Durable     bool
	AutoDeleted bool
	Exclusive   bool
	NoWait      bool
	Arguments   amqp.Table
}

var DefaultQueueOption = QueueOption{
	Name:        "",
	Durable:     false,
	AutoDeleted: true,
	Exclusive:   true,
	NoWait:      false,
	Arguments:   nil,
}

type QueueBindOption struct {
	RoutingKey string
	NoWait     bool
	Arguments  amqp.Table
}

var DefaultQueueBindOption = QueueBindOption{
	RoutingKey: "",
	NoWait:     false,
	Arguments:  nil,
}

type ConsumeOption struct {
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Arguments amqp.Table
}

var DefaultConsumeOption = ConsumeOption{
	AutoAck:   true,
	Exclusive: false,
	NoWait:    false,
	Arguments: nil,
}

type TopicConsumerOption struct {
	ExchangeName string
	RoutingKey   string
}
