package rabbitmq

import (
	"github.com/streadway/amqp"
)

type PublishArgs struct {
	RoutingKey string
	Mandatory  bool
	Immediate  bool
	Publishing amqp.Publishing
}

var DefaultPublishArgs = PublishArgs{
	RoutingKey: "",
	Mandatory:  false,
	Immediate:  false,
	Publishing: amqp.Publishing{},
}

type TopicPublisherOption struct {
	ExchangeName string
	RoutingKey   string
	Publishing   amqp.Publishing
}
