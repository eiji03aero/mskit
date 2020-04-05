package kitchen

import (
	"encoding/json"

	"order"
	orderdmn "order/domain/order"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/streadway/amqp"
)

type proxy struct {
	client *rabbitmq.Client
}

func New(c *rabbitmq.Client) order.KitchenProxy {
	return &proxy{
		client: c,
	}
}

func (p *proxy) CreateTicket(
	restaurantId string,
	orderLineItems orderdmn.OrderLineItems,
) (ticketId string, err error) {
	reqBody := struct {
		Id              string                  `json:"restaurant_id"`
		TicketLineItems orderdmn.OrderLineItems `json:"ticket_line_items"`
	}{
		Id:              restaurantId,
		TicketLineItems: orderLineItems,
	}

	cmdJson, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	delivery, err := p.client.NewRPCClient().
		Configure(
			rabbitmq.PublishArgs{
				RoutingKey: "kitchen.rpc.create-ticket",
				Publishing: amqp.Publishing{
					Body: cmdJson,
				},
			},
		).
		Exec()
	if err != nil {
		return
	}

	response := struct {
		TicketId string `json:"ticket_id"`
	}{}
	err = json.Unmarshal(delivery.Body, &response)
	if err != nil {
		return
	}

	ticketId = response.TicketId

	return
}

func (p *proxy) CancelTicket(id string) (err error) {
	reqBody := struct {
		Id string `json:"id"`
	}{
		Id: id,
	}

	cmdJson, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	_, err = p.client.NewRPCClient().
		Configure(
			rabbitmq.PublishArgs{
				RoutingKey: "kitchen.rpc.cancel-ticket",
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

func (p *proxy) ConfirmTicket(id string) (err error) {
	reqBody := struct {
		Id string `json:"id"`
	}{
		Id: id,
	}

	cmdJson, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	_, err = p.client.NewRPCClient().
		Configure(
			rabbitmq.PublishArgs{
				RoutingKey: "kitchen.rpc.confirm-ticket",
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

func (p *proxy) BeginReviseTicket(ticketId string, orderLineItems orderdmn.OrderLineItems) (err error) {
	reqBody := struct {
		Id              string                  `json:"id"`
		TicketLineItems orderdmn.OrderLineItems `json:"ticket_line_items"`
	}{
		Id:              ticketId,
		TicketLineItems: orderLineItems,
	}

	cmdJson, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	_, err = p.client.NewRPCClient().
		Configure(
			rabbitmq.PublishArgs{
				RoutingKey: "kitchen.rpc.begin-revise-ticket",
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

func (p *proxy) UndoBeginReviseTicket(ticketId string) (err error) {
	reqBody := struct {
		Id string `json:"id"`
	}{
		Id: ticketId,
	}

	cmdJson, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	_, err = p.client.NewRPCClient().
		Configure(
			rabbitmq.PublishArgs{
				RoutingKey: "kitchen.rpc.undo-begin-revise-ticket",
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

func (p *proxy) ConfirmReviseTicket(ticketId string, orderLineItems orderdmn.OrderLineItems) (err error) {
	reqBody := struct {
		Id              string                  `json:"id"`
		TicketLineItems orderdmn.OrderLineItems `json:"ticket_line_items"`
	}{
		Id:              ticketId,
		TicketLineItems: orderLineItems,
	}

	cmdJson, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	_, err = p.client.NewRPCClient().
		Configure(
			rabbitmq.PublishArgs{
				RoutingKey: "kitchen.rpc.confirm-revise-ticket",
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
