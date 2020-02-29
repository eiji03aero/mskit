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
		Host: "rabbitmq",
		Port: "5672",
	})
	if err != nil {
		panic(err)
	}

	go func() {
		go client.NewConsumer().
			Configure(
				"complain-consumer-all",
				exchangeMap["complain"],
				queueMap["complain"],
				rabbitmq.QueueBindOption{
					Key:    "complain.#",
					NoWait: false,
				},
				rabbitmq.ConsumeOption{
					NoAck:     false,
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
					Key:    "complain.error",
					NoWait: false,
				},
				rabbitmq.ConsumeOption{
					NoAck:     false,
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
	}()

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

		time.Sleep(3 * time.Second)

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
	}()

	log.Println("Starting publisher ... ")
	time.Sleep(10 * time.Second)
	log.Println("Closing ... ")
}
