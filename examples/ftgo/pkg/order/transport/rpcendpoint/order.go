package rpcendpoint

import (
	"encoding/json"
	"log"

	orderdmn "order/domain/order"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/streadway/amqp"
)

func (c *client) runRejectOrder() {
	c.c.NewRPCEndpoint().
		Configure(
			rabbitmq.QueueOption{
				Name: "order.rpc.reject-order",
			},
		).
		OnDelivery(func(d amqp.Delivery) (p amqp.Publishing) {
			log.Println("client#runRejectOrder: ", string(d.Body))
			p = amqp.Publishing{}

			command := orderdmn.RejectOrder{}
			err := json.Unmarshal(d.Body, &command)
			if err != nil {
				return rabbitmq.MakeFailResponse(p)
			}

			err = c.svc.RejectOrder(command)
			if err != nil {
				return rabbitmq.MakeFailResponse(p)
			}

			return rabbitmq.MakeSuccessResponse(p)
		}).
		Exec()
}
