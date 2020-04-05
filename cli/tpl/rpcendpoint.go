package tpl

func RPCEndpointTemplate() string {
	return `package rpcendpoint

import (
	"{{ .PkgName }}"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/streadway/amqp"
)

type rpcEndpoint struct {
	client  *rabbitmq.Client
	service {{ .PkgName }}.Service
}

func New(c *rabbitmq.Client, svc {{ .PkgName }}.Service) *rpcEndpoint {
	return &rpcEndpoint{
		client:  c,
		service: svc,
	}
}

func (re *rpcEndpoint) Run() (err error) {
	go re.sample()
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
			return rabbitmq.MakeSuccessResponse(p)
		}).
		Exec()
}`
}
