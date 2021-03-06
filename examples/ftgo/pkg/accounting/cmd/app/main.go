package main

import (
	"accounting/adapter/consumer"
	"accounting/adapter/rpcendpoint"
	accountdmn "accounting/domain/account"
	"accounting/service"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/db/mongo"
	"github.com/eiji03aero/mskit/db/mongo/eventstore"
	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/eiji03aero/mskit/utils/logger"
)

var (
	dbOption = mongo.DBOption{
		Host: "ftgo-accounting-mongo",
		Port: "27017",
	}
	rabbitmqOption = rabbitmq.Option{
		Host: "ftgo-rabbitmq",
		Port: "5672",
	}
)

func main() {
	er := mskit.NewEventRegistry()
	er.Set(accountdmn.AccountCreated{})

	eb, err := rabbitmq.NewClient(rabbitmqOption)
	if err != nil {
		panic(err)
	}

	es, err := eventstore.New(dbOption, er)
	if err != nil {
		panic(err)
	}
	erp := mskit.NewEventRepository(es, &mskit.StubEventPublisher{})

	svc := service.New(erp)

	err = consumer.New(eb, svc).Run()
	if err != nil {
		panic(err)
	}

	err = rpcendpoint.New(eb, svc).Run()
	if err != nil {
		panic(err)
	}

	logger.Println("server starting to listen ...")
	bff := make(chan bool)
	<-bff
}
