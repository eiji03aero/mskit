package proxy

import (
	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
)

type Proxy struct {
	client *rabbitmq.Client
}

func New(
	c *rabbitmq.Client,
) *Proxy {
	return &Proxy{
		client: c,
	}
}
