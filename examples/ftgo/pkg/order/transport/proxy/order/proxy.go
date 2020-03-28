package order

import (
	orderroot "order"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
)

type proxy struct {
	client *rabbitmq.Client
}

func New(
	client *rabbitmq.Client,
) orderroot.OrderProxy {
	return &proxy{
		client: client,
	}
}
