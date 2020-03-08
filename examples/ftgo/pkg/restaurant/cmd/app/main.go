package main

import (
	"log"
	"net/http"
	restaurantdmn "restaurant/domain/restaurant"
	restaurantsvc "restaurant/service"
	httptransport "restaurant/transport/http"
	"restaurant/transport/publisher"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/eiji03aero/mskit/eventstore/mongo"
)

var (
	dbOption = mongo.DBOption{
		Host: "ftgo-restaurant-mongo",
		Port: "27017",
	}
	rabbitmqOption = rabbitmq.Option{
		Host: "ftgo-rabbitmq",
		Port: "5672",
	}
)

func main() {
	er := mskit.NewEventRegistry()
	er.Set(restaurantdmn.RestaurantCreated{})

	eventBusClient, err := rabbitmq.NewClient(rabbitmqOption)
	if err != nil {
		panic(err)
	}
	dep := publisher.New(eventBusClient)

	eventStore, err := mongo.New(dbOption, er)
	if err != nil {
		panic(err)
	}
	repository := mskit.NewRepository(eventStore, dep)

	svc := restaurantsvc.New(repository, dep)
	mux := httptransport.New(svc)

	log.Println("server starting to listen ...")
	http.ListenAndServe(":3002", mux)
}
