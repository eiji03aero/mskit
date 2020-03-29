package tpl

func RPCEndpointTemplate() string {
	return `package rpcendpoint

import (
	"{{ .PkgName }}/service"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
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

func (re *rpcEndpoint) Run() {
	// Initilizing code comes here
	// go rs.sample()
}

func (re *rpcEndpoint) sample() {
	re.client.NewRPCEndpoint().
		Configure(
			rabbitmq.QueueOption{
				Name: "",
			},
		).
		OnDelivery(func(d amqp.Delivery) (p amqp.Publishing) {
			return rabbitmq.MakeSuccessResponse(p)
		}).
		Exec()
}`
}
