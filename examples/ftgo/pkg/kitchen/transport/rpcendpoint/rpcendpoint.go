package rpcendpoint

import (
	"encoding/json"

	"kitchen"
	ticketdmn "kitchen/domain/ticket"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/eiji03aero/mskit/utils/logger"
	"github.com/streadway/amqp"
)

type rpcEndpoint struct {
	client  *rabbitmq.Client
	service kitchen.Service
}

func New(c *rabbitmq.Client, svc kitchen.Service) *rpcEndpoint {
	return &rpcEndpoint{
		client:  c,
		service: svc,
	}
}

func (re *rpcEndpoint) Run() (err error) {
	go re.runCreateTicket()
	go re.runCancelTicket()

	return
}

func (re *rpcEndpoint) runCreateTicket() {
	re.client.NewRPCEndpoint().
		Configure(
			rabbitmq.QueueOption{
				Name: "kitchen.rpc.create-ticket",
			},
		).
		OnDelivery(func(d amqp.Delivery) (p amqp.Publishing) {
			logger.PrintFuncCall(re.runCreateTicket, d.Body)

			cmd := ticketdmn.CreateTicket{}
			err := json.Unmarshal(d.Body, &cmd)
			if err != nil {
				return rabbitmq.MakeFailResponse(p)
			}

			ticketId, err := re.service.CreateTicket(cmd)
			if err != nil {
				return rabbitmq.MakeFailResponse(p)
			}

			response := struct {
				TicketId string `json:"ticket_id"`
			}{
				TicketId: ticketId,
			}
			resJson, err := json.Marshal(response)
			if err != nil {
				return rabbitmq.MakeFailResponse(p)
			}

			p.Body = resJson

			return rabbitmq.MakeSuccessResponse(p)
		}).
		Exec()
}

func (re *rpcEndpoint) runCancelTicket() {
	re.client.NewRPCEndpoint().
		Configure(
			rabbitmq.QueueOption{
				Name: "kitchen.rpc.cancel-ticket",
			},
		).
		OnDelivery(func(d amqp.Delivery) (p amqp.Publishing) {
			logger.PrintFuncCall(re.runCancelTicket, d.Body)

			cmd := ticketdmn.CancelTicket{}
			err := json.Unmarshal(d.Body, &cmd)
			if err != nil {
				return rabbitmq.MakeFailResponse(p)
			}

			err = re.service.CancelTicket(cmd)
			if err != nil {
				return rabbitmq.MakeFailResponse(p)
			}

			return rabbitmq.MakeSuccessResponse(p)
		}).
		Exec()
}
