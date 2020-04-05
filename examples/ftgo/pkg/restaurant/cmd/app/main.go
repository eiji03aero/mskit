package main

import (
	"net/http"
	restaurantdmn "restaurant/domain/restaurant"
	restaurantsvc "restaurant/service"
	httptransport "restaurant/transport/http"
	"restaurant/transport/publisher"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/db/mongo"
	"github.com/eiji03aero/mskit/db/mongo/eventstore"
	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/eiji03aero/mskit/utils/logger"
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

	eb, err := rabbitmq.NewClient(rabbitmqOption)
	if err != nil {
		panic(err)
	}
	dep := publisher.New(eb)

	es, err := eventstore.New(dbOption, er)
	if err != nil {
		panic(err)
	}
	erp := mskit.NewEventRepository(es, dep)

	svc := restaurantsvc.New(erp, dep)
	mux := httptransport.New(svc)

	logger.Println("server starting to listen ...")
	http.ListenAndServe(":3002", mux)
}
