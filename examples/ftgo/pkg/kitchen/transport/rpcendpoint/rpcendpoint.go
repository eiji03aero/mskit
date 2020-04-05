package rpcendpoint

import (
	"encoding/json"

	"kitchen"
	ticketdmn "kitchen/domain/ticket"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
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
	go re.runConfirmTicket()
	go re.runBeginReviseTicket()
	go re.runUndoBeginReviseTicket()
	go re.runConfirmReviseTicket()

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
			cmd := ticketdmn.CreateTicket{}
			err := json.Unmarshal(d.Body, &cmd)
			if err != nil {
				return rabbitmq.MakeFailResponse(p, err)
			}

			ticketId, err := re.service.CreateTicket(cmd)
			if err != nil {
				return rabbitmq.MakeFailResponse(p, err)
			}

			response := struct {
				TicketId string `json:"ticket_id"`
			}{
				TicketId: ticketId,
			}
			resJson, err := json.Marshal(response)
			if err != nil {
				return rabbitmq.MakeFailResponse(p, err)
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
			cmd := ticketdmn.CancelTicket{}
			err := json.Unmarshal(d.Body, &cmd)
			if err != nil {
				return rabbitmq.MakeFailResponse(p, err)
			}

			err = re.service.CancelTicket(cmd)
			if err != nil {
				return rabbitmq.MakeFailResponse(p, err)
			}

			return rabbitmq.MakeSuccessResponse(p)
		}).
		Exec()
}

func (re *rpcEndpoint) runConfirmTicket() {
	re.client.NewRPCEndpoint().
		Configure(
			rabbitmq.QueueOption{
				Name: "kitchen.rpc.confirm-ticket",
			},
		).
		OnDelivery(func(d amqp.Delivery) (p amqp.Publishing) {
			cmd := ticketdmn.ConfirmTicket{}
			err := json.Unmarshal(d.Body, &cmd)
			if err != nil {
				return rabbitmq.MakeFailResponse(p, err)
			}

			err = re.service.ConfirmTicket(cmd)
			if err != nil {
				return rabbitmq.MakeFailResponse(p, err)
			}

			return rabbitmq.MakeSuccessResponse(p)
		}).
		Exec()
}

func (re *rpcEndpoint) runBeginReviseTicket() {
	re.client.NewRPCEndpoint().
		Configure(
			rabbitmq.QueueOption{
				Name: "kitchen.rpc.begin-revise-ticket",
			},
		).
		OnDelivery(func(d amqp.Delivery) (p amqp.Publishing) {
			cmd := ticketdmn.BeginReviseTicket{}
			err := json.Unmarshal(d.Body, &cmd)
			if err != nil {
				return rabbitmq.MakeFailResponse(p, err)
			}

			err = re.service.BeginReviseTicket(cmd)
			if err != nil {
				return rabbitmq.MakeFailResponse(p, err)
			}

			return rabbitmq.MakeSuccessResponse(p)
		}).
		Exec()
}

func (re *rpcEndpoint) runUndoBeginReviseTicket() {
	re.client.NewRPCEndpoint().
		Configure(
			rabbitmq.QueueOption{
				Name: "kitchen.rpc.undo-begin-revise-ticket",
			},
		).
		OnDelivery(func(d amqp.Delivery) (p amqp.Publishing) {
			cmd := ticketdmn.UndoBeginReviseTicket{}
			err := json.Unmarshal(d.Body, &cmd)
			if err != nil {
				return rabbitmq.MakeFailResponse(p, err)
			}

			err = re.service.UndoBeginReviseTicket(cmd)
			if err != nil {
				return rabbitmq.MakeFailResponse(p, err)
			}

			return rabbitmq.MakeSuccessResponse(p)
		}).
		Exec()
}

func (re *rpcEndpoint) runConfirmReviseTicket() {
	re.client.NewRPCEndpoint().
		Configure(
			rabbitmq.QueueOption{
				Name: "kitchen.rpc.confirm-revise-ticket",
			},
		).
		OnDelivery(func(d amqp.Delivery) (p amqp.Publishing) {
			cmd := ticketdmn.ConfirmReviseTicket{}
			err := json.Unmarshal(d.Body, &cmd)
			if err != nil {
				return rabbitmq.MakeFailResponse(p, err)
			}

			err = re.service.ConfirmReviseTicket(cmd)
			if err != nil {
				return rabbitmq.MakeFailResponse(p, err)
			}

			return rabbitmq.MakeSuccessResponse(p)
		}).
		Exec()
}
