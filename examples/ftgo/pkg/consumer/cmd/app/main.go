package main

import (
	"log"
	"net/http"

	consumerdmn "consumer/domain/consumer"
	consumersvc "consumer/service"
	httptransport "consumer/transport/http"
	"consumer/transport/rpcendpoint"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/db/mongo"
	"github.com/eiji03aero/mskit/db/mongo/eventstore"
	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
)

var (
	dbOption = mongo.DBOption{
		Host: "ftgo-consumer-mongo",
		Port: "27017",
	}
	rabbitmqOption = rabbitmq.Option{
		Host: "ftgo-rabbitmq",
		Port: "5672",
	}
)

func main() {
	er := mskit.NewEventRegistry()
	er.Set(consumerdmn.ConsumerCreated{})

	eventBusClient, err := rabbitmq.NewClient(rabbitmqOption)
	if err != nil {
		panic(err)
	}

	es, err := eventstore.New(dbOption, er)
	if err != nil {
		panic(err)
	}
	eventRepository := mskit.NewEventRepository(es, &mskit.StubDomainEventPublisher{})

	svc := consumersvc.New(eventRepository)

	err = rpcendpoint.New(eventBusClient, svc).Run()
	if err != nil {
		panic(err)
	}

	mux := httptransport.New(svc)

	log.Println("server starting to listen ...")
	http.ListenAndServe(":3003", mux)
}
