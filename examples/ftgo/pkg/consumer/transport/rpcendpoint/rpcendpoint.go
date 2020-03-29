package rpcendpoint

import (
	"encoding/json"

	consumerdmn "consumer/domain/consumer"
	consumersvc "consumer/service"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/eiji03aero/mskit/utils/logger"
	"github.com/streadway/amqp"
)

type rpcEndpoint struct {
	c   *rabbitmq.Client
	svc consumersvc.Service
}

func New(c *rabbitmq.Client, svc consumersvc.Service) *rpcEndpoint {
	return &rpcEndpoint{
		c:   c,
		svc: svc,
	}
}

func (re *rpcEndpoint) Run() (err error) {
	go re.runValidateOrder()

	return
}

func (re *rpcEndpoint) runValidateOrder() {
	re.c.NewRPCEndpoint().
		Configure(
			rabbitmq.QueueOption{
				Name: "consumer.rpc.validate-order",
			},
		).
		OnDelivery(func(d amqp.Delivery) (p amqp.Publishing) {
			logger.PrintFuncCall(re.runValidateOrder, string(d.Body))

			command := consumerdmn.ValidateOrder{}
			err := json.Unmarshal(d.Body, command)
			if err != nil {
				return rabbitmq.MakeFailResponse(p)
			}

			err = re.svc.ValidateOrder(command)
			if err != nil {
				return rabbitmq.MakeFailResponse(p)
			}

			return rabbitmq.MakeSuccessResponse(p)
		}).
		Exec()
}
