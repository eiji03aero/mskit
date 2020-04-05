package rpcendpoint

import (
	"accounting"
	"encoding/json"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/eiji03aero/mskit/utils/logger"
	"github.com/streadway/amqp"
)

type rpcEndpoint struct {
	client  *rabbitmq.Client
	service accounting.Service
}

func New(c *rabbitmq.Client, svc accounting.Service) *rpcEndpoint {
	return &rpcEndpoint{
		client:  c,
		service: svc,
	}
}

func (re *rpcEndpoint) Run() (err error) {
	go re.runAuthorize()
	return
}

func (re *rpcEndpoint) runAuthorize() {
	re.client.NewRPCEndpoint().
		Configure(
			rabbitmq.QueueOption{
				Name: "accounting.rpc.authorize",
			},
		).
		OnDelivery(func(d amqp.Delivery) (p amqp.Publishing) {
			request := struct {
				ConsumerId string `json:"consumer_id"`
			}{}
			logger.PrintFuncCall(re.runAuthorize, string(d.Body))

			err := json.Unmarshal(d.Body, &request)
			if err != nil {
				return rabbitmq.MakeFailResponse(p, err)
			}

			err = re.service.Authorize(request.ConsumerId)
			if err != nil {
				return rabbitmq.MakeFailResponse(p, err)
			}

			return rabbitmq.MakeSuccessResponse(p)
		}).
		Exec()
}
