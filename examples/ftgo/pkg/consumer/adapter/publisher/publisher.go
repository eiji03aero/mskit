package publisher

import (
	"encoding/json"

	consumerdmn "consumer/domain/consumer"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/eiji03aero/mskit/utils/errbdr"
	"github.com/streadway/amqp"
)

type publisher struct {
	client *rabbitmq.Client
}

func New(c *rabbitmq.Client) mskit.EventPublisher {
	return &publisher{
		client: c,
	}
}

func (p *publisher) Publish(event interface{}) (err error) {
	ej, err := json.Marshal(event)
	if err != nil {
		return
	}

	switch e := event.(type) {
	case consumerdmn.ConsumerCreated:
		return p.publishConsumerCreated(e, ej)
	default:
		return errbdr.NewErrUnknownParams(p.Publish, e)
	}
}

func (p *publisher) publishConsumerCreated(event consumerdmn.ConsumerCreated, eventJson []byte) (err error) {
	return p.client.NewPublisher().
		Configure(
			rabbitmq.TopicPublisherOption{
				ExchangeName: "topic-consumer",
				RoutingKey:   "consumer.consumer.created",
				Publishing: amqp.Publishing{
					Body: eventJson,
				},
			},
		).
		Exec()
}
