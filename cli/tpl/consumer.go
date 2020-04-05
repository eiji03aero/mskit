package tpl

func ConsumerTemplate() string {
	return `package consumer

import (
	"{{ .PkgName }}"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/streadway/amqp"
)

type consumer struct {
	client   *rabbitmq.Client
	service {{ .PkgName }}.Service
}

func New(c *rabbitmq.Client, svc {{ .PkgName }}.Service) *consumer {
	return &consumer{
		client:   c,
		service: svc,
	}
}

func (c *consumer) Run() error {
	go c.runSample()

	return nil
}

func (c *consumer) runSample() {
	c.client.NewConsumer().
		Configure(
			rabbitmq.TopicConsumerOption{
				ExchangeName: "topic-sample",
				RoutingKey:   "sample.sample.created",
			},
		).
		OnDelivery(func(d amqp.Delivery) {
			// logic here
		}).
		Exec()
}`
}
