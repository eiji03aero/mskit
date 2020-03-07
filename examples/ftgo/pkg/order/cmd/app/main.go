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
	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/eiji03aero/mskit/eventstore/postgres"
)

func main() {
	dbOption := postgres.DBOption{
		User:     "ftgo",
		Password: "ftgo123",
		Host:     "ftgo-order-postgres",
		Port:     "5432",
		Name:     "ftgo",
	}

	er := mskit.NewEventRegistry()
	er.Set(orderdmn.OrderCreated{})

	eventStore, err := postgres.New(dbOption, er)
	if err != nil {
		panic(err)
	}

	eventBusClient, err := rabbitmq.NewClient(rabbitmq.Option{
		Host: "ftgo-rabbitmq",
		Port: "5672",
	})
	if err != nil {
		panic(err)
	}

	db, err := postgres.GetDB(dbOption)
	if err != nil {
		panic(err)
	}

	restaurantRepository := restaurantrepo.New(db)

	repository := mskit.NewRepository(eventStore)
	svc := service.New(
		repository,
		restaurantRepository,
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
