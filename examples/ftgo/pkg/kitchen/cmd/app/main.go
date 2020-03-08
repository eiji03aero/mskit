package main

import (
	logcommon "common/log"
	kitchendmn "kitchen/domain/kitchen"
	kitchensvc "kitchen/service"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/eventstore/mongo"
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

	eventStore, err := mongo.New(dbOption, er)
	if err != nil {
		panic(err)
	}

	svc := kitchensvc.New(
		mskit.NewRepository(eventStore, &mskit.StubDomainEventPublisher{}),
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
	err = eventStore.Load(id, &ticket)
	if err != nil {
		panic(err)
	}

	logcommon.PrintJsonln(ticket)
}
