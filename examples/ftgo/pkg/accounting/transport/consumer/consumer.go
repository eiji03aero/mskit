package consumer

import (
	"accounting"
	accountdmn "accounting/domain/account"
	"encoding/json"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/eiji03aero/mskit/utils/logger"
	"github.com/streadway/amqp"
)

type consumer struct {
	client  *rabbitmq.Client
	service accounting.Service
}

func New(c *rabbitmq.Client, svc accounting.Service) *consumer {
	return &consumer{
		client:  c,
		service: svc,
	}
}

func (c *consumer) Run() error {
	go c.runConsumerCreated()

	return nil
}

func (c *consumer) runConsumerCreated() {
	c.client.NewConsumer().
		Configure(
			rabbitmq.TopicConsumerOption{
				ExchangeName: "topic-consumer",
				RoutingKey:   "consumer.consumer.created",
			},
		).
		OnDelivery(func(d amqp.Delivery) {
			logger.PrintFuncCall(c.runConsumerCreated, d.Body)

			response := struct {
				Id string `json:"id"`
			}{}
			err := json.Unmarshal(d.Body, &response)
			if err != nil {
				return
			}

			cmd := accountdmn.CreateAccount{
				ConsumerId: response.Id,
			}
			_, err = c.service.CreateAccount(cmd)
		}).
		Exec()
}
