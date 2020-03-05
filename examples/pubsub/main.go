package main

import (
	"fmt"
	"log"
	"time"

	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/streadway/amqp"
)

var exchangeMap = map[string]rabbitmq.ExchangeOption{
	"complain": rabbitmq.ExchangeOption{
		Name:        "complain",
		Type:        "topic",
		Durable:     false,
		AutoDeleted: true,
		Internal:    false,
		NoWait:      false,
	},
}

var queueMap = map[string]rabbitmq.QueueOption{
	"complain": rabbitmq.QueueOption{
		Name:        "",
		Durable:     false,
		AutoDeleted: true,
		Exclusive:   false,
		NoWait:      false,
	},
}

func main() {
	client, err := rabbitmq.NewClient(rabbitmq.Option{
		Host: "ftgo-rabbitmq",
		Port: "5672",
	})
	if err != nil {
		panic(err)
	}

	finish := make(chan error, 2)

	go func() {
		go client.NewConsumer().
			Configure(
				"complain-consumer-all",
				exchangeMap["complain"],
				queueMap["complain"],
				rabbitmq.QueueBindOption{
					RoutingKey: "complain.#",
					NoWait:     false,
				},
				rabbitmq.ConsumeOption{
					AutoAck:   true,
					Exclusive: false,
					NoLocal:   false,
					NoWait:    false,
				},
			).
			OnDelivery(func(d amqp.Delivery) {
				fmt.Println("complain-consumer-all: received and printing to log!")
				fmt.Println(string(d.Body))
			}).
			Exec()

		go client.NewConsumer().
			Configure(
				"complain-consumer-error",
				exchangeMap["complain"],
				queueMap["complain"],
				rabbitmq.QueueBindOption{
					RoutingKey: "complain.error",
					NoWait:     false,
				},
				rabbitmq.ConsumeOption{
					AutoAck:   true,
					Exclusive: false,
					NoLocal:   false,
					NoWait:    false,
				},
			).
			OnDelivery(func(d amqp.Delivery) {
				fmt.Println("complain-consumer-error: received and printing to log!")
				fmt.Println(string(d.Body))
			}).
			Exec()

		finish <- nil
	}()

	time.Sleep(2 * time.Second)

	go func() {
		complainExchange := rabbitmq.ExchangeOption{
			Name:        "complain",
			Type:        "topic",
			Durable:     false,
			AutoDeleted: true,
			Internal:    false,
			NoWait:      false,
		}

		complainPublishArgs := rabbitmq.PublishArgs{
			RoutingKey: "complain",
			Mandatory:  false,
			Immediate:  false,
			Publishing: amqp.Publishing{
				Body: []byte("mou dame desu yametaro"),
			},
		}
		complainPublishArgsError := rabbitmq.PublishArgs{
			RoutingKey: "complain.error",
			Mandatory:  false,
			Immediate:  false,
			Publishing: amqp.Publishing{
				Body: []byte("hontoni yabai yametaro"),
			},
		}

		go client.NewPublisher().
			Configure(
				complainExchange,
				complainPublishArgs,
			).
			Exec()

		go client.NewPublisher().
			Configure(
				complainExchange,
				complainPublishArgsError,
			).
			Exec()

		time.Sleep(2 * time.Second)
		finish <- nil
	}()

	log.Println("Starting publisher ... ")
	<-finish
	<-finish
	log.Println("Closing ... ")
}
