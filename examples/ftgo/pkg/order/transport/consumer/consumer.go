package consumer

import (
	"encoding/json"
	"log"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/streadway/amqp"
	restaurantdmn "order/domain/restaurant"
	ordersvc "order/service"
)

type consumer struct {
	c   *rabbitmq.Client
	svc ordersvc.Service
}

func New(c *rabbitmq.Client, svc ordersvc.Service) *consumer {
	return &consumer{
		c:   c,
		svc: svc,
	}
}

func (c *consumer) Run() error {
	go func() {
		go c.c.NewConsumer().
			Configure(
				rabbitmq.TopicConsumerOption{
					ExchangeName: "topic-restaurant",
					RoutingKey:   "restaurant.restaurant.created",
				},
			).
			OnDelivery(func(d amqp.Delivery) {
				var restaurant restaurantdmn.Restaurant
				err := json.Unmarshal(d.Body, &restaurant)
				if err != nil {
					log.Println(err)
					return
				}

				err = c.svc.CreateRestaurant(restaurant)
				if err != nil {
					log.Println(err)
					return
				}
			}).
			Exec()
	}()

	return nil
}
