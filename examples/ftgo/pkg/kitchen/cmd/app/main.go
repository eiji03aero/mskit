package main

import (
	logcommon "common/log"
	kitchendmn "kitchen/domain/kitchen"
	kitchensvc "kitchen/service"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/eventstore/mongo"
)

func main() {
	dbOption := mongo.DBOption{
		Host: "ftgo-kitchen-mongo",
		Port: "27017",
	}

	er := mskit.NewEventRegistry()
	er.Set(kitchendmn.TicketCreated{})

	eventStore, err := mongo.New(dbOption, er)
	if err != nil {
		panic(err)
	}

	repository := mskit.NewRepository(eventStore)

	svc := kitchensvc.New(repository)

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

	id, _ := svc.CreateTicket(cmd)

	ticket := kitchendmn.Ticket{}
	err = eventStore.Load(id, &ticket)
	if err != nil {
		panic(err)
	}

	logcommon.PrintJsonln(ticket)
}
