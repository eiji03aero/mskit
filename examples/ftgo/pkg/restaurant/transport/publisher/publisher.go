package publisher

import (
	"encoding/json"

	restaurantdmn "restaurant/domain/restaurant"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/eiji03aero/mskit/utils/errbdr"
	"github.com/streadway/amqp"
)

type publisher struct {
	c *rabbitmq.Client
}

func New(c *rabbitmq.Client) mskit.EventPublisher {
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
		err = p.publishRestaurantCreated(e, ej)
	default:
		err = errbdr.NewErrUnknownParams(p.Publish, e)
	}

	return
}

func (p *publisher) publishRestaurantCreated(event restaurantdmn.RestaurantCreated, eventJson []byte) (err error) {
	return p.c.NewPublisher().
		Configure(
			rabbitmq.TopicPublisherOption{
				ExchangeName: "topic-restaurant",
				RoutingKey:   "restaurant.restaurant.created",
				Publishing: amqp.Publishing{
					Body: eventJson,
				},
			},
		).
		Exec()
}
