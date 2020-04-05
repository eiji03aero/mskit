package consumer

import (
	"encoding/json"

	"order"
	restaurantdmn "order/domain/restaurant"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/streadway/amqp"
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

func (p *proxy) ValidateOrder(orderId string, total int) (err error) {
	cmdJson, err := json.Marshal(restaurantdmn.ValidateOrder{
		OrderId: orderId,
		Total:   total,
	})
	if err != nil {
		return
	}

	_, err = p.client.NewRPCClient().
		Configure(
			rabbitmq.PublishArgs{
				RoutingKey: "consumer.rpc.validate-order",
				Publishing: amqp.Publishing{
					Body: cmdJson,
				},
			},
		).
		Exec()

	return
}
