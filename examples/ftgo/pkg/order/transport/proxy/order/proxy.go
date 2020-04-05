package order

import (
	"encoding/json"

	orderroot "order"
	orderdmn "order/domain/order"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/streadway/amqp"
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

func (p *proxy) RejectOrder(id string) (err error) {
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

func (p *proxy) ApproveOrder(id string) (err error) {
	cmdJson, err := json.Marshal(orderdmn.ApproveOrder{
		Id: id,
	})
	if err != nil {
		return
	}

	_, err = p.client.NewRPCClient().
		Configure(
			rabbitmq.PublishArgs{
				RoutingKey: "order.rpc.approve-order",
				Publishing: amqp.Publishing{
					Body: cmdJson,
				},
			},
		).
		Exec()

	return
}

func (p *proxy) BeginReviseOrder(id string) (err error) {
	cmdJson, err := json.Marshal(orderdmn.BeginReviseOrder{
		Id: id,
	})
	if err != nil {
		return
	}

	_, err = p.client.NewRPCClient().
		Configure(
			rabbitmq.PublishArgs{
				RoutingKey: "order.rpc.begin-revise-order",
				Publishing: amqp.Publishing{
					Body: cmdJson,
				},
			},
		).
		Exec()

	return
}

func (p *proxy) UndoBeginReviseOrder(id string) (err error) {
	cmdJson, err := json.Marshal(orderdmn.UndoBeginReviseOrder{
		Id: id,
	})
	if err != nil {
		return
	}

	_, err = p.client.NewRPCClient().
		Configure(
			rabbitmq.PublishArgs{
				RoutingKey: "order.rpc.undo-begin-revise-order",
				Publishing: amqp.Publishing{
					Body: cmdJson,
				},
			},
		).
		Exec()

	return
}

func (p *proxy) ConfirmReviseOrder(id string, orderLineItems orderdmn.OrderLineItems) (err error) {
	cmdJson, err := json.Marshal(orderdmn.ConfirmReviseOrder{
		Id:             id,
		OrderLineItems: orderLineItems,
	})
	if err != nil {
		return
	}

	_, err = p.client.NewRPCClient().
		Configure(
			rabbitmq.PublishArgs{
				RoutingKey: "order.rpc.confirm-revise-order",
				Publishing: amqp.Publishing{
					Body: cmdJson,
				},
			},
		).
		Exec()

	return
}
