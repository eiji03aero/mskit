package kitchen

import (
	errorscommon "common/errors"
	"github.com/eiji03aero/mskit"
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
	case CreateTicketCommand:
		return t.processCreateTicketCommand(c)
	default:
		return mskit.Events{}, errorscommon.ErrNotSupportedParams(t.Process, c)
	}
}

func (t *Ticket) Apply(event interface{}) error {
	switch e := event.(type) {
	case *TicketCreatedEvent:
		return t.applyTicketCreatedEvent(e)
	default:
		return errorscommon.ErrNotSupportedParams(t.Apply, e)
	}
}

func (t *Ticket) processCreateTicketCommand(cmd CreateTicketCommand) (mskit.Events, error) {
	events := mskit.NewEventsSingle(
		cmd.Id,
		Ticket{},
		&TicketCreatedEvent{
			Id:              cmd.Id,
			RestaurantId:    cmd.RestaurantId,
			TicketLineItems: cmd.TicketLineItems,
		},
	)

	return events, nil
}

func (t *Ticket) applyTicketCreatedEvent(event *TicketCreatedEvent) error {
	t.Id = event.Id
	t.RestaurantId = event.RestaurantId
	t.TicketLineItems = event.TicketLineItems

	t.State = TicketState_CreatePending
	t.PreviousState = TicketState_CreatePending

	return nil
}
