package rabbitmq

import (
	"github.com/streadway/amqp"
)

func ExchangeDeclare(ch *amqp.Channel, eo ExchangeOption) error {
	return ch.ExchangeDeclare(
		eo.Name,
		eo.Type,
		eo.Durable,
		eo.AutoDeleted,
		eo.Internal,
		eo.NoWait,
		eo.Arguments,
	)
}

func Publish(ch *amqp.Channel, ename string, pubargs PublishArgs) error {
	return ch.Publish(
		ename,
		pubargs.RoutingKey,
		pubargs.Mandatory,
		pubargs.Immediate,
		pubargs.Publishing,
	)
}

func QueueDeclare(ch *amqp.Channel, qo QueueOption) (amqp.Queue, error) {
	return ch.QueueDeclare(
		qo.Name,
		qo.Durable,
		qo.AutoDeleted,
		qo.Exclusive,
		qo.NoWait,
		qo.Arguments,
	)
}

func QueueBind(ch *amqp.Channel, queueName string, exchangeName string, qbo QueueBindOption) error {
	return ch.QueueBind(
		queueName,
		qbo.RoutingKey,
		exchangeName,
		qbo.NoWait,
		qbo.Arguments,
	)
}

func Consume(ch *amqp.Channel, queueName string, tag string, co ConsumeOption) (<-chan amqp.Delivery, error) {
	return ch.Consume(
		queueName,
		tag,
		co.AutoAck,
		co.Exclusive,
		co.NoLocal,
		co.NoWait,
		co.Arguments,
	)
}
