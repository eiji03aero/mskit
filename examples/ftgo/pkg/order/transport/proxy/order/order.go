package order

import (
	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
)

type proxy struct {
	c *rabbitmq.Client
}

func New(c *rabbitmq.Client) *proxy {
	return &proxy{
		c: c,
	}
}
