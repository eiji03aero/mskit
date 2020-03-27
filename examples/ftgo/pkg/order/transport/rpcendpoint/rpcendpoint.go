package rpcendpoint

import (
	ordersvc "order/service"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
)

type client struct {
	c   *rabbitmq.Client
	svc ordersvc.Service
}

func New(c *rabbitmq.Client, svc ordersvc.Service) *client {
	return &client{
		c:   c,
		svc: svc,
	}
}

func (c *client) Run() (err error) {
	go c.runRejectOrder()

	return
}
