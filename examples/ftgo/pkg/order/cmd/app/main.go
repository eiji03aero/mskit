package main

import (
	"log"
	"net/http"

	orderdmn "order/domain/order"
	restaurantrepo "order/repository/restaurant"
	"order/saga/createorder"
	"order/service"
	"order/transport/consumer"
	httptransport "order/transport/http"
	consumerpxy "order/transport/proxy/consumer"
	orderpxy "order/transport/proxy/order"
	"order/transport/rpcendpoint"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/db/postgres"
	"github.com/eiji03aero/mskit/db/postgres/eventstore"
	"github.com/eiji03aero/mskit/db/postgres/sagastore"
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
	ss, err := sagastore.New(dbOption)
	if err != nil {
		panic(err)
	}
	sr := mskit.NewSagaRepository(ss)

	eventBusClient, err := rabbitmq.NewClient(rabbitmqOption)
	if err != nil {
		panic(err)
	}

	orderProxy := orderpxy.New(eventBusClient)
	consumerProxy := consumerpxy.New(eventBusClient)

	svc := service.New(
		mskit.NewEventRepository(es, &mskit.StubDomainEventPublisher{}),
		restaurantrepo.New(db),
	)

	createOrderSagaManager := createorder.NewManager(
		sr,
		svc,
		orderProxy,
		consumerProxy,
	)
	go createOrderSagaManager.Subscribe()

	svc.InjectSagaManagers(
		createOrderSagaManager,
	)

	err = consumer.New(eventBusClient, svc).Run()
	if err != nil {
		panic(err)
	}

	err = rpcendpoint.New(eventBusClient, svc).Run()
	if err != nil {
		panic(err)
	}

	mux := httptransport.New(svc)
	log.Println("server starting to listen ...")
	http.ListenAndServe(":3000", mux)
}
