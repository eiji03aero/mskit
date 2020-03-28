package order

import (
	"encoding/json"
	"log"
	orderdmn "order/domain/order"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/streadway/amqp"
)

func (p *proxy) RejectOrder(id string) (err error) {
	log.Println("Proxy#RejectOrder: ", id)

	cmdJson, err := json.Marshal(orderdmn.RejectOrder{
		Id: id,
	})
	if err != nil {
		return
	}

	_, err = p.client.NewRPCClient().
		Configure(
			rabbitmq.PublishArgs{
				RoutingKey: "order.rpc.reject-order",
				Publishing: amqp.Publishing{
					Body: cmdJson,
				},
			},
		).
		Exec()

	return
}
