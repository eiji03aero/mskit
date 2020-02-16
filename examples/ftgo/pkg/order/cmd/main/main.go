package main

import (
	"log"
	"net/http"

	"github.com/eiji03aero/mskit/eventstore/postgres"
	"github.com/eiji03aero/mskit/examples/ftgo/pkg/order/service"
	httptransport "github.com/eiji03aero/mskit/examples/ftgo/pkg/order/transport/http"
)

func main() {
	repository, err := postgres.New(
		postgres.DBOption{
			User:     "ftgo",
			Password: "ftgo123",
			Host:     "ftgo-postgres",
			Port:     "5432",
			Name:     "ftgo",
		},
	)
	if err != nil {
		panic(err)
	}

	svc := service.New(repository)
	mux := httptransport.New(svc)

	log.Println("server starting to listen ...")
	http.ListenAndServe(":3000", mux)
}
