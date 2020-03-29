package kitchen

import (
	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/utils/errbdr"
)

type Ticket struct {
	mskit.BaseAggregate
	State           TicketState     `json:"state"`
	PreviousState   TicketState     `json:"previous_state"`
	RestaurantId    string          `json:"restaurantId"`
	TicketLineItems TicketLineItems `json:"ticket_line_items"`
}

func (t *Ticket) Validate() (errs []error) {
	return errs
}

func (t *Ticket) Process(cmd interface{}) (mskit.Events, error) {
	switch c := cmd.(type) {
	case CreateTicket:
		return t.processCreateTicket(c)
	default:
		return mskit.Events{}, errbdr.NewErrUnknownParams(t.Process, c)
	}
}

func (t *Ticket) Apply(event interface{}) error {
	switch e := event.(type) {
	case TicketCreated:
		return t.applyTicketCreated(e)
	default:
		return errbdr.NewErrUnknownParams(t.Apply, e)
	}
}

func (t *Ticket) processCreateTicket(cmd CreateTicket) (mskit.Events, error) {
	events := mskit.NewEventsSingle(
		cmd.Id,
		Ticket{},
		TicketCreated{
			Id:              cmd.Id,
			RestaurantId:    cmd.RestaurantId,
			TicketLineItems: cmd.TicketLineItems,
		},
	)

	return events, nil
}

func (t *Ticket) applyTicketCreated(event TicketCreated) error {
	t.Id = event.Id
	t.RestaurantId = event.RestaurantId
	t.TicketLineItems = event.TicketLineItems

	t.State = TicketState_CreatePending
	t.PreviousState = TicketState_CreatePending

	return nil
}
