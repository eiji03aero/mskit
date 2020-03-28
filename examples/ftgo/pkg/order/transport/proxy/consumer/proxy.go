package consumer

import (
	"order"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
)

type proxy struct {
	client *rabbitmq.Client
}

func New(
	client *rabbitmq.Client,
) order.ConsumerProxy {
	return &proxy{
		client: client,
	}
}
