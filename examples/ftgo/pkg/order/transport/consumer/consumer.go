package consumer

import (
	"encoding/json"

	"order"
	restaurantdmn "order/domain/restaurant"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/eiji03aero/mskit/utils/logger"
	"github.com/streadway/amqp"
)

type consumer struct {
	c   *rabbitmq.Client
	svc order.Service
}

func New(c *rabbitmq.Client, svc order.Service) *consumer {
	return &consumer{
		c:   c,
		svc: svc,
	}
}

func (c *consumer) Run() error {
	go c.runRestaurantCreated()

	return nil
}

func (c *consumer) runRestaurantCreated() {
	c.c.NewConsumer().
		Configure(
			rabbitmq.TopicConsumerOption{
				ExchangeName: "topic-restaurant",
				RoutingKey:   "restaurant.restaurant.created",
			},
		).
		OnDelivery(func(d amqp.Delivery) {
			logger.PrintFuncCall(c.runRestaurantCreated, d.Body)

			var restaurant restaurantdmn.Restaurant
			err := json.Unmarshal(d.Body, &restaurant)
			if err != nil {
				logger.Println(err)
				return
			}

			err = c.svc.CreateRestaurant(restaurant)
			if err != nil {
				logger.Println(err)
				return
			}
		}).
		Exec()
}
