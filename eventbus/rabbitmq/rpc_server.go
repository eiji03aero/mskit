package rabbitmq

import (
	"github.com/streadway/amqp"
)

type RPCServer struct {
	conn            *amqp.Connection
	QueueOption     QueueOption
	ConsumeOption   ConsumeOption
	DeliveryHandler func(d amqp.Delivery) amqp.Publishing
}

func NewRPCServer(conn *amqp.Connection) *RPCServer {
	return &RPCServer{
		conn: conn,
	}
}

func (rs *RPCServer) Configure(
	qopt QueueOption,
) *RPCServer {
	rs.QueueOption = qopt
	return rs
}

func (rs *RPCServer) OnDelivery(cb func(d amqp.Delivery) amqp.Publishing) *RPCServer {
	rs.DeliveryHandler = cb
	return rs
}

func (rs *RPCServer) Exec() (err error) {
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