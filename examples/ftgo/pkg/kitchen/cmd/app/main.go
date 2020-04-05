package main

import (
	ticketdmn "kitchen/domain/ticket"
	kitchensvc "kitchen/service"
	"kitchen/transport/rpcendpoint"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/db/mongo"
	"github.com/eiji03aero/mskit/db/mongo/eventstore"
	"github.com/eiji03aero/mskit/eventbus/rabbitmq"
	"github.com/eiji03aero/mskit/utils/logger"
)

var (
	dbOption = mongo.DBOption{
		Host: "ftgo-kitchen-mongo",
		Port: "27017",
	}
	rabbitmqOption = rabbitmq.Option{
		Host: "ftgo-rabbitmq",
		Port: "5672",
	}
)

func main() {
	er := mskit.NewEventRegistry()
	er.Set(ticketdmn.TicketCreated{})
	er.Set(ticketdmn.TicketCancelled{})
	er.Set(ticketdmn.TicketConfirmed{})
	er.Set(ticketdmn.TicketRevisionBegan{})
	er.Set(ticketdmn.UndoTicketRevisionBegan{})
	er.Set(ticketdmn.TicketRevisionConfirmed{})

	es, err := eventstore.New(dbOption, er)
	if err != nil {
		panic(err)
	}

	eb, err := rabbitmq.NewClient(rabbitmqOption)

	svc := kitchensvc.New(
		mskit.NewEventRepository(es, &mskit.StubEventPublisher{}),
	)

	err = rpcendpoint.New(eb, svc).Run()
	if err != nil {
		panic(err)
	}

	logger.Println("server starting to listen ...")
	bff := make(chan bool, 1)
	<-bff
}
