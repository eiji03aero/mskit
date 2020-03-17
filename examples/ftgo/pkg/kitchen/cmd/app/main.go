package main

import (
	logcommon "common/log"
	kitchendmn "kitchen/domain/kitchen"
	kitchensvc "kitchen/service"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/db/mongo"
	"github.com/eiji03aero/mskit/db/mongo/eventstore"
)

var (
	dbOption = mongo.DBOption{
		Host: "ftgo-kitchen-mongo",
		Port: "27017",
	}
)

func main() {
	er := mskit.NewEventRegistry()
	er.Set(kitchendmn.TicketCreated{})

	es, err := eventstore.New(dbOption, er)
	if err != nil {
		panic(err)
	}

	svc := kitchensvc.New(
		mskit.NewRepository(es, &mskit.StubDomainEventPublisher{}),
	)

	cmd := kitchendmn.CreateTicket{
		RestaurantId: "mac",
		TicketLineItems: kitchendmn.TicketLineItems{
			LineItems: []kitchendmn.TicketLineItem{
				kitchendmn.TicketLineItem{
					Quantity:   2,
					MenuItemId: "fries",
				},
				kitchendmn.TicketLineItem{
					Quantity:   5,
					MenuItemId: "fries L size",
				},
			},
		},
	}

	id, err := svc.CreateTicket(cmd)
	if err != nil {
		panic(err)
	}

	ticket := kitchendmn.Ticket{}
	err = es.Load(id, &ticket)
	if err != nil {
		panic(err)
	}

	logcommon.PrintJsonln(ticket)
}
