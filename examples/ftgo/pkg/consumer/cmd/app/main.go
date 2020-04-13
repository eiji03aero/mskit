package main

import (
	"net/http"

	httpadapter "consumer/adapter/http"
	"consumer/adapter/publisher"
	"consumer/adapter/rpcendpoint"
	consumerdmn "consumer/domain/consumer"
	consumersvc "consumer/service"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/db/mongo"
	"github.com/eiji03aero/mskit/db/mongo/eventstore"
	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/eiji03aero/mskit/utils/logger"
)

var (
	dbOption = mongo.DBOption{
		Host: "ftgo-consumer-mongo",
		Port: "27017",
	}
	rabbitmqOption = rabbitmq.Option{
		Host: "ftgo-rabbitmq",
		Port: "5672",
	}
)

func main() {
	er := mskit.NewEventRegistry()
	er.Set(consumerdmn.ConsumerCreated{})

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

	svc := consumersvc.New(erp)

	err = rpcendpoint.New(eb, svc).Run()
	if err != nil {
		panic(err)
	}

	mux := httpadapter.New(svc)

	logger.Println("server starting to listen ...")
	http.ListenAndServe(":3003", mux)
}
