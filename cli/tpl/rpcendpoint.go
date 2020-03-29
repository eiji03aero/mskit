package tpl

func RPCEndpointTemplate() string {
	return `package rpcendpoint

import (
	"{{ .PkgName }}/service"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/eiji03aero/mskit/utils/logger"
	"github.com/streadway/amqp"
)

type rpcEndpoint struct {
	client  *rabbitmq.Client
	service service.Service
}

func New(c *rabbitmq.Client, svc service.Service) *rpcEndpoint {
	return &rpcEndpoint{
		client:  c,
		service: svc,
	}
}

func (re *rpcEndpoint) Run() (err error) {
	// Initilizing code comes here
	// go re.sample()
	return
}

func (re *rpcEndpoint) sample() {
	re.client.NewRPCEndpoint().
		Configure(
			rabbitmq.QueueOption{
				Name: "{{ .PkgName }}.rpc.",
			},
		).
		OnDelivery(func(d amqp.Delivery) (p amqp.Publishing) {
			logger.PrintFuncCall(re.sample, string(d.Body))
			return rabbitmq.MakeSuccessResponse(p)
		}).
		Exec()
}`
}
