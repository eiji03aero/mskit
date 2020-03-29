package rpcendpoint

import (
	"encoding/json"

	ticketdmn "kitchen/domain/ticket"
	"kitchen/service"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/eiji03aero/mskit/utils/logger"
	"github.com/streadway/amqp"
)

type rpcEndpoint struct {
	client  *rabbitmq.Client
	service service.Service
}

func New(c *rabbitmq.Client, svc service.Service) *rpcEndpoint {
	return &rpcEndpoint{
		client:  c,
		service: svc,
	}
}

func (re *rpcEndpoint) Run() (err error) {
	go re.runCreateTicket()
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
