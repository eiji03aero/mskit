package main

import (
	"log"
	"net/http"

	orderdmn "order/domain/order"
	"order/service"
	httptransport "order/transport/http"

	"github.com/eiji03aero/mskit"
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
	// err := postgres.InitializeDB(dbOption)
	// if err != nil {
	// 	panic(err)
	// }

	er := mskit.NewEventRegistry()
	er.Set(orderdmn.OrderCreated{})

	eventStore, err := postgres.New(dbOption, er)
	if err != nil {
		panic(err)
	}

	repository := mskit.NewRepository(eventStore)
	svc := service.New(repository)
	mux := httptransport.New(svc)

	log.Println("server starting to listen ...")
	http.ListenAndServe(":3000", mux)
}
