package main

import (
	"fmt"
	"time"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/streadway/amqp"
)

func main() {
	client, err := rabbitmq.NewClient(rabbitmq.Option{
		Host: "ftgo-rabbitmq",
		Port: "5672",
	})
	if err != nil {
		panic(err)
	}

	finish := make(chan error)

	go func() {
		client.NewRPCServer().
			Configure(
				rabbitmq.QueueOption{
					Name: "calc-onegai-shimasu",
				},
			).
			OnDelivery(func(d amqp.Delivery) (p amqp.Publishing) {
				p.Body = []byte("kore ga calc no kekka")
				return
			}).
			Exec()
	}()

	time.Sleep(1 * time.Second)

	go func() {
		delivery, err := client.NewRPCClient().
			Configure(
				rabbitmq.PublishArgs{
					RoutingKey: "calc-onegai-shimasu",
				},
			).
			Exec()
		if err != nil {
			panic(err)
		}

		fmt.Println(string(delivery.Body))
		finish <- nil
	}()

	<-finish
}
