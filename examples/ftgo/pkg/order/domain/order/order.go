package order

import (
	"errors"
	"fmt"

	"github.com/eiji03aero/mskit"
)

type Order struct {
	mskit.BaseAggregate
	Name string
}

func (o *Order) Process(cmd interface{}) ([]*mskit.Event, error) {
	switch c := cmd.(type) {
	case *CreateOrder:
		return o.processCreateOrder(c)
	default:
		return nil, errors.New("not imp in Process")
	}
}

func (o *Order) Apply(event *mskit.Event) error {
	switch e := event.Data.(type) {
	case *OrderCreated:
		return o.applyOrderCreated(e)
	default:
		return errors.New(fmt.Sprintf("not implemented in Apply: %v", e))
	}
}

type CreateOrder struct {
	ID   string
	Name string
}

type OrderCreated struct {
	Name string `json:"name"`
}

func (o *Order) processCreateOrder(cmd *CreateOrder) ([]*mskit.Event, error) {
	events := []*mskit.Event{
		mskit.NewEvent(
			cmd.ID,
			&Order{},
			&OrderCreated{
				Name: cmd.Name,
			},
		),
	}

	return events, nil
}

func (o *Order) applyOrderCreated(event *OrderCreated) error {
	o.Name = event.Name
	return nil
}
