package rpcendpoint

import (
	"encoding/json"

	consumerroot "consumer"
	consumerdmn "consumer/domain/consumer"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/streadway/amqp"
)

type rpcEndpoint struct {
	client  *rabbitmq.Client
	service consumerroot.Service
}

func New(c *rabbitmq.Client, svc consumerroot.Service) *rpcEndpoint {
	return &rpcEndpoint{
		client:  c,
		service: svc,
	}
}

func (re *rpcEndpoint) Run() (err error) {
	go re.runValidateOrder()

	return
}

func (re *rpcEndpoint) runValidateOrder() {
	re.client.NewRPCEndpoint().
		Configure(
			rabbitmq.QueueOption{
				Name: "consumer.rpc.validate-order",
			},
		).
		OnDelivery(func(d amqp.Delivery) (p amqp.Publishing) {
			command := consumerdmn.ValidateOrder{}
			err := json.Unmarshal(d.Body, &command)
			if err != nil {
				return rabbitmq.MakeFailResponse(p, err)
			}

			err = re.service.ValidateOrder(command)
			if err != nil {
				return rabbitmq.MakeFailResponse(p, err)
			}

			return rabbitmq.MakeSuccessResponse(p)
		}).
		Exec()
}
