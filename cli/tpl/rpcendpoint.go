package tpl

func RPCEndpointTemplate() string {
	return `package rpcendpoint

import (
	"{{ .PkgName }}/service"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
)

type rpcServer struct {
	c   *rabbitmq.Client
	svc service.Service
}

func New(c *rabbitmq.Client, svc service.Service) *rpcServer {
	return &client{
		c:   c,
		svc: svc,
	}
}

func (rs *rpcServer) Run() {
	// Initilizing code comes here
	// go rs.sample()
}

func (rs *rpcServer) sample() {
	c.c.NewRPCEndpoint().
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
