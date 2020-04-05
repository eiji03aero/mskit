package rabbitmq

import (
	"errors"
	"time"

	"github.com/eiji03aero/mskit/utils"
	"github.com/eiji03aero/mskit/utils/logger"
	"github.com/streadway/amqp"
)

type RPCClient struct {
	conn        *amqp.Connection
	PublishArgs PublishArgs
}

func NewRPCClient(conn *amqp.Connection) *RPCClient {
	return &RPCClient{
		conn: conn,
	}
}

func (rc *RPCClient) Configure(
	pubargs PublishArgs,
) *RPCClient {
	rc.PublishArgs = pubargs
	return rc
}

func (rc *RPCClient) Exec() (delivery amqp.Delivery, err error) {
	channel, err := rc.conn.Channel()
	if err != nil {
		return
	}
	defer channel.Close()

	replyQueue, err := QueueDeclare(channel, QueueOption{
		Name:        "",
		Durable:     false,
		AutoDeleted: false,
		Exclusive:   true,
		NoWait:      false,
		Arguments:   nil,
	})
	if err != nil {
		return
	}

	msgs, err := Consume(channel, replyQueue.Name, "", ConsumeOption{
		AutoAck:   true,
		Exclusive: false,
		NoLocal:   false,
		NoWait:    false,
		Arguments: nil,
	})
	if err != nil {
		return
	}

	corrId, err := utils.UUID()
	if err != nil {
		return
	}

	rc.PublishArgs.Publishing.ReplyTo = replyQueue.Name
	rc.PublishArgs.Publishing.CorrelationId = corrId

	logger.Println(
		logger.YellowString("Send rpc request"),
		logger.CyanString(rc.PublishArgs.RoutingKey),
		rc.PublishArgs.Publishing.Body,
	)
	// Please take care of us default exchange, I'm too lazy :(
	err = Publish(channel, "", rc.PublishArgs)
	if err != nil {
		return
	}

	deliveryChan := make(chan amqp.Delivery, 1)

	go func() {
		for d := range msgs {
			if corrId == d.CorrelationId {
				deliveryChan <- d
				break
			}
		}
	}()

	select {
	case delivery = <-deliveryChan:
	case <-time.After(60 * time.Second):
		err = errors.New("rpc request timed out")
		logger.Println(
			logger.RedString(err.Error()),
			logger.CyanString(rc.PublishArgs.RoutingKey),
			rc.PublishArgs.Publishing.Body,
		)
		return
	}

	if !IsSuccessResponse(delivery) {
		errMsg := getErrorMessage(delivery)
		if errMsg == "" {
			errMsg = "rpc failed"
		}

		err = errors.New(errMsg)
	}

	return
}
