package rpcendpoint

import (
	"encoding/json"

	consumerdmn "consumer/domain/consumer"
	consumersvc "consumer/service"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/streadway/amqp"
)

type client struct {
	c   *rabbitmq.Client
	svc consumersvc.Service
}

func New(c *rabbitmq.Client, svc consumersvc.Service) *client {
	return &client{
		c:   c,
		svc: svc,
	}
}

func (c *client) Run() (err error) {
	go c.runValidateOrder()

	return
}

func (c *client) runValidateOrder() {
	c.c.NewRPCEndpoint().
		Configure(
			rabbitmq.QueueOption{
				Name: "consumer.rpc.validate-order",
			},
		).
		OnDelivery(func(d amqp.Delivery) (p amqp.Publishing) {
			p = amqp.Publishing{}

			command := consumerdmn.ValidateOrder{}
			err := json.Unmarshal(d.Body, command)
			if err != nil {
				return rabbitmq.MakeFailResponse(p)
			}

			err = c.svc.ValidateOrder(command)
			if err != nil {
				return rabbitmq.MakeFailResponse(p)
			}

			return rabbitmq.MakeSuccessResponse(p)
		})
}
