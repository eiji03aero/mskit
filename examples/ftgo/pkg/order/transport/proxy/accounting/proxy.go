package accounting

import (
	"encoding/json"
	"order"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/streadway/amqp"
)

type proxy struct {
	client *rabbitmq.Client
}

func New(c *rabbitmq.Client) order.AccountingProxy {
	return &proxy{
		client: c,
	}
}

func (p *proxy) Authorize(consumerId string) (err error) {
	reqBody := struct {
		ConsumerId string `json:"consumer_id"`
	}{
		ConsumerId: consumerId,
	}

	cmdJson, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	_, err = p.client.NewRPCClient().
		Configure(
			rabbitmq.PublishArgs{
				RoutingKey: "accounting.rpc.authorize",
				Publishing: amqp.Publishing{
					Body: cmdJson,
				},
			},
		).
		Exec()
	if err != nil {
		return
	}

	return
}
