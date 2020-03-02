package main

import (
	"log"
	"net/http"
	restaurantdmn "restaurant/domain/restaurant"
	restaurantsvc "restaurant/service"
	httptransport "restaurant/transport/http"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/eventstore/mongo"
)

func main() {
	dbOption := mongo.DBOption{
		Host: "ftgo-restaurant-mongo",
		Port: "27017",
	}

	er := mskit.NewEventRegistry()
	er.Set(restaurantdmn.RestaurantCreated{})

	eventStore, err := mongo.New(dbOption, er)
	if err != nil {
		panic(err)
	}

	repository := mskit.NewRepository(eventStore)
	svc := restaurantsvc.New(repository)
	mux := httptransport.New(svc)

	log.Println("server starting to listen ...")
	http.ListenAndServe(":3002", mux)
}
