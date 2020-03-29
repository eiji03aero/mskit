package kitchen

import (
	"encoding/json"

	"order"
	orderdmn "order/domain/order"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/eiji03aero/mskit/utils/logger"
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
	lineItems []orderdmn.OrderLineItem,
) (ticketId string, err error) {
	logger.PrintFuncCall(p.CreateTicket, restaurantId, lineItems)

	reqBody := struct {
		Id        string                   `json:"restaurant_id"`
		LineItems []orderdmn.OrderLineItem `json:"line_items"`
	}{
		Id:        restaurantId,
		LineItems: lineItems,
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
	logger.PrintFuncCall(p.CancelTicket, id)

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
