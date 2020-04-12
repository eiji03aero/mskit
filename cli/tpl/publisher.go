package tpl

func PublisherTemplate() string {
	return `package publisher

import (
	"encoding/json"

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
	case "sample":
		return p.publishSample(e, ej)
	default:
		return errbdr.NewErrUnknownParams(p.Publish, e)
	}
}

func (p *publisher) publishSample(event interface{}, eventJson []byte) (err error) {
	return p.client.NewPublisher().
		Configure(
			rabbitmq.TopicPublisherOption{
				ExchangeName: "topic-sample",
				RoutingKey:   "sample.sample.created",
				Publishing: amqp.Publishing{
					Body: eventJson,
				},
			},
		).
		Exec()
}`
}
