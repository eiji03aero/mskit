package main

import (
	"net/http"

	orderdmn "order/domain/order"
	restaurantrepo "order/repository/restaurant"
	"order/saga/createorder"
	"order/saga/reviseorder"
	"order/service"
	"order/transport/consumer"
	httptransport "order/transport/http"
	accountingpxy "order/transport/proxy/accounting"
	consumerpxy "order/transport/proxy/consumer"
	kitchenpxy "order/transport/proxy/kitchen"
	orderpxy "order/transport/proxy/order"
	"order/transport/rpcendpoint"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/db/postgres"
	"github.com/eiji03aero/mskit/db/postgres/eventstore"
	"github.com/eiji03aero/mskit/db/postgres/sagastore"
	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/eiji03aero/mskit/utils/logger"
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

	eventRegistry := mskit.NewEventRegistry()
	eventRegistry.Set(orderdmn.OrderCreated{})
	eventRegistry.Set(orderdmn.OrderRejected{})
	eventRegistry.Set(orderdmn.OrderApproved{})
	eventRegistry.Set(orderdmn.OrderRevisionBegan{})
	eventRegistry.Set(orderdmn.UndoOrderRevisionBegan{})
	eventRegistry.Set(orderdmn.OrderRevisionConfirmed{})
	eventRegistry.Set(orderdmn.OrderTicketIdSet{})

	eventStore, err := eventstore.New(dbOption, eventRegistry)
	if err != nil {
		panic(err)
	}
	sagaStore, err := sagastore.New(dbOption)
	if err != nil {
		panic(err)
	}
	eventRepository := mskit.NewEventRepository(eventStore, &mskit.StubEventPublisher{})
	sagaRepository := mskit.NewSagaRepository(sagaStore)
	restaurantRepository := restaurantrepo.New(db)

	eb, err := rabbitmq.NewClient(rabbitmqOption)
	if err != nil {
		panic(err)
	}

	orderProxy := orderpxy.New(eb)
	consumerProxy := consumerpxy.New(eb)
	kitchenProxy := kitchenpxy.New(eb)
	accountingProxy := accountingpxy.New(eb)

	svc := service.New(
		eventRepository,
		restaurantRepository,
	)

	createOrderSagaManager := createorder.NewManager(
		sagaRepository,
		svc,
		orderProxy,
		consumerProxy,
		kitchenProxy,
		accountingProxy,
	)
	go createOrderSagaManager.Subscribe()

	reviseOrderSagaManager := reviseorder.NewManager(
		sagaRepository,
		svc,
		orderProxy,
		kitchenProxy,
		accountingProxy,
	)
	go reviseOrderSagaManager.Subscribe()

	svc.InjectSagaManagers(
		createOrderSagaManager,
		reviseOrderSagaManager,
	)

	err = consumer.New(eb, svc).Run()
	if err != nil {
		panic(err)
	}

	err = rpcendpoint.New(eb, svc).Run()
	if err != nil {
		panic(err)
	}

	mux := httptransport.New(svc)
	logger.Println("server starting to listen ...")
	http.ListenAndServe(":3000", mux)
}
