package main

import (
	"log"
	"net/http"

	"order/pb"
	"order/service"
	httptransport "order/transport/http"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/eventstore/postgres"
)

func main() {
	dbOptions := postgres.DBOption{
		User:     "ftgo",
		Password: "ftgo123",
		Host:     "ftgo-order-postgres",
		Port:     "5432",
		Name:     "ftgo",
	}
	// err := postgres.InitializeDB(dbOptions)
	// if err != nil {
	// 	panic(err)
	// }

	er := mskit.NewEventRegistry()
	er.Set(pb.OrderCreated{})

	eventStore, err := postgres.New(dbOptions, er)
	if err != nil {
		panic(err)
	}

	repository := mskit.NewRepository(eventStore)
	svc := service.New(repository)
	mux := httptransport.New(svc)

	log.Println("server starting to listen ...")
	http.ListenAndServe(":3000", mux)
}
