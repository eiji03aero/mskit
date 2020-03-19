package publisher

import (
	"encoding/json"

	errorscommon "common/errors"
	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/streadway/amqp"
	restaurantdmn "restaurant/domain/restaurant"
)

type publisher struct {
	c *rabbitmq.Client
}

func New(c *rabbitmq.Client) mskit.DomainEventPublisher {
	return &publisher{
		c: c,
	}
}

func (p *publisher) Publish(event interface{}) (err error) {
	ej, err := json.Marshal(event)
	if err != nil {
		return
	}

	switch e := event.(type) {
	case restaurantdmn.RestaurantCreated:
		err = p.c.NewPublisher().
			Configure(
				rabbitmq.TopicPublisherOption{
					ExchangeName: "topic-restaurant",
					RoutingKey:   "restaurant.restaurant.created",
					Publishing: amqp.Publishing{
						Body: ej,
					},
				},
			).
			Exec()

	default:
		err = errorscommon.NewErrNotSupportedParams(p.Publish, e)
	}

	return
}
