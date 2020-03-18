package rabbitmq

import (
	"github.com/streadway/amqp"
)

type RPCEndpoint struct {
	conn            *amqp.Connection
	QueueOption     QueueOption
	ConsumeOption   ConsumeOption
	DeliveryHandler func(d amqp.Delivery) amqp.Publishing
}

func NewRPCEndpoint(conn *amqp.Connection) *RPCEndpoint {
	return &RPCEndpoint{
		conn: conn,
	}
}

func (rs *RPCEndpoint) Configure(
	qopt QueueOption,
	copt ConsumeOption,
) *RPCEndpoint {
	rs.QueueOption = qopt
	rs.ConsumeOption = copt
	return rs
}

func (rs *RPCEndpoint) OnDelivery(cb func(d amqp.Delivery) amqp.Publishing) *RPCEndpoint {
	rs.DeliveryHandler = cb
	return rs
}

func (rs *RPCEndpoint) Exec() (err error) {
	channel, err := rs.conn.Channel()
	if err != nil {
		return
	}
	defer channel.Close()

	_, err = QueueDeclare(channel, rs.QueueOption)
	if err != nil {
		return
	}

	msgs, err := Consume(channel, rs.QueueOption.Name, "", rs.ConsumeOption)
	if err != nil {
		return
	}

	for d := range msgs {
		publishing := rs.DeliveryHandler(d)
		publishing.CorrelationId = d.CorrelationId
		err = Publish(channel, "", PublishArgs{
			RoutingKey: d.ReplyTo,
			Mandatory:  false,
			Immediate:  false,
			Publishing: publishing,
		})
		if err != nil {
			return
		}
	}

	return
}
