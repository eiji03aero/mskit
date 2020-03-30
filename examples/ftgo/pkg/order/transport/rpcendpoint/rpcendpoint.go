package rpcendpoint

import (
	"encoding/json"

	"order"
	orderdmn "order/domain/order"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/eiji03aero/mskit/utils/logger"
	"github.com/streadway/amqp"
)

type rpcEndpoint struct {
	c   *rabbitmq.Client
	svc order.Service
}

func New(c *rabbitmq.Client, svc order.Service) *rpcEndpoint {
	return &rpcEndpoint{
		c:   c,
		svc: svc,
	}
}

func (re *rpcEndpoint) Run() (err error) {
	go re.runRejectOrder()
	go re.runApproveOrder()
	return
}

func (re *rpcEndpoint) runRejectOrder() {
	re.c.NewRPCEndpoint().
		Configure(
			rabbitmq.QueueOption{
				Name: "order.rpc.reject-order",
			},
		).
		OnDelivery(func(d amqp.Delivery) (p amqp.Publishing) {
			logger.PrintFuncCall(re.runRejectOrder, string(d.Body))

			command := orderdmn.RejectOrder{}
			err := json.Unmarshal(d.Body, &command)
			if err != nil {
				return rabbitmq.MakeFailResponse(p)
			}

			err = re.svc.RejectOrder(command)
			if err != nil {
				return rabbitmq.MakeFailResponse(p)
			}

			return rabbitmq.MakeSuccessResponse(p)
		}).
		Exec()
}

func (re *rpcEndpoint) runApproveOrder() {
	re.c.NewRPCEndpoint().
		Configure(
			rabbitmq.QueueOption{
				Name: "order.rpc.approve-order",
			},
		).
		OnDelivery(func(d amqp.Delivery) (p amqp.Publishing) {
			logger.PrintFuncCall(re.runApproveOrder, string(d.Body))

			command := orderdmn.ApproveOrder{}
			err := json.Unmarshal(d.Body, &command)
			if err != nil {
				return rabbitmq.MakeFailResponse(p)
			}

			err = re.svc.ApproveOrder(command)
			if err != nil {
				return rabbitmq.MakeFailResponse(p)
			}

			return rabbitmq.MakeSuccessResponse(p)
		}).
		Exec()
}
