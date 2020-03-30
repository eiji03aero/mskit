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
	case CancelTicket:
		return t.processCancelTicket(c)
	case ConfirmTicket:
		return t.processConfirmTicket(c)
	default:
		return mskit.Events{}, errbdr.NewErrUnknownParams(t.Process, c)
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

func (t *Ticket) processCancelTicket(cmd CancelTicket) (mskit.Events, error) {
	events := mskit.NewEventsSingle(
		cmd.Id,
		Ticket{},
		TicketCancelled{
			Id: cmd.Id,
		},
	)

	return events, nil
}

func (t *Ticket) processConfirmTicket(cmd ConfirmTicket) (mskit.Events, error) {
	events := mskit.NewEventsSingle(
		cmd.Id,
		Ticket{},
		TicketConfirmed{
			Id: cmd.Id,
		},
	)

	return events, nil
}

func (t *Ticket) Apply(event interface{}) error {
	switch e := event.(type) {
	case TicketCreated:
		return t.applyTicketCreated(e)
	case TicketCancelled:
		return t.applyTicketCancelled(e)
	case TicketConfirmed:
		return t.applyTicketConfirmed(e)
	default:
		return errbdr.NewErrUnknownParams(t.Apply, e)
	}
}

func (t *Ticket) applyTicketCreated(event TicketCreated) error {
	t.Id = event.Id
	t.RestaurantId = event.RestaurantId
	t.TicketLineItems = event.TicketLineItems

	t.State = TicketState_CreatePending
	t.PreviousState = TicketState_CreatePending

	return nil
}

func (t *Ticket) applyTicketCancelled(event TicketCancelled) error {
	t.State = TicketState_Canceled
	t.PreviousState = t.State

	return nil
}

func (t *Ticket) applyTicketConfirmed(event TicketConfirmed) error {
	t.State = TicketState_Preparing
	t.PreviousState = t.State

	return nil
}
