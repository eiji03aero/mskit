package consumer

import (
	"encoding/json"
	"log"
	restaurantdmn "order/domain/restaurant"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/streadway/amqp"
)

func (p *proxy) ValidateOrder(orderId string, total int) (err error) {
	log.Println("Proxy#ValidateOrder: ", orderId, total)

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
