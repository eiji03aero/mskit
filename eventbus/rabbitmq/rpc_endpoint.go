package rabbitmq

import (
	"github.com/eiji03aero/mskit/utils/logger"
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
		conn:          conn,
		QueueOption:   DefaultQueueOption,
		ConsumeOption: DefaultConsumeOption,
	}
}

func (rs *RPCEndpoint) Configure(opts ...interface{}) *RPCEndpoint {
	for _, opt := range opts {
		switch o := opt.(type) {
		case QueueOption:
			rs.QueueOption = o
		case ConsumeOption:
			rs.ConsumeOption = o
		}
	}

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
		logger.Println(
			logger.YellowString("RPCEndpoint received request:"),
			logger.CyanString(rs.QueueOption.Name),
			d.Body,
		)

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
