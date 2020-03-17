package main

import (
	"log"
	"net/http"

	orderdmn "order/domain/order"
	restaurantrepo "order/repository/restaurant"
	"order/service"
	"order/transport/consumer"
	httptransport "order/transport/http"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/db/postgres"
	"github.com/eiji03aero/mskit/db/postgres/eventstore"
	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
)

var (
	dbOption = postgres.DBOption{
		User:     "ftgo",
		Password: "ftgo123",
		Host:     "ftgo-order-postgres",
		Port:     "5432",
		Name:     "ftgo",
	}
	rabbitmqOption = rabbitmq.Option{
		Host: "ftgo-rabbitmq",
		Port: "5672",
	}
)

func main() {
	db, err := postgres.GetDB(dbOption)
	if err != nil {
		panic(err)
	}

	er := mskit.NewEventRegistry()
	er.Set(orderdmn.OrderCreated{})

	es, err := eventstore.New(dbOption, er)
	if err != nil {
		panic(err)
	}

	eventBusClient, err := rabbitmq.NewClient(rabbitmqOption)
	if err != nil {
		panic(err)
	}

	svc := service.New(
		mskit.NewRepository(es, &mskit.StubDomainEventPublisher{}),
		restaurantrepo.New(db),
	)

	err = consumer.New(eventBusClient, svc).
		Run()
	if err != nil {
		panic(err)
	}

	mux := httptransport.New(svc)
	log.Println("server starting to listen ...")
	http.ListenAndServe(":3000", mux)
}
