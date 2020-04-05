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
	case BeginReviseTicket:
		return t.processBeginReviseTicket(c)
	case UndoBeginReviseTicket:
		return t.processUndoBeginReviseTicket(c)
	case ConfirmReviseTicket:
		return t.processConfirmReviseTicket(c)
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

func (t *Ticket) processBeginReviseTicket(cmd BeginReviseTicket) (events mskit.Events, err error) {
	switch t.State {
	case TicketState_Accepted:
		fallthrough
	case TicketState_AwaitingAcceptance:
		events = mskit.NewEventsSingle(
			cmd.Id,
			Ticket{},
			TicketRevisionBegan{
				Id: cmd.Id,
			},
		)
	default:
		err = errbdr.NewErrUnsupportedStateTransition(t, t.State)
	}

	return
}

func (t *Ticket) processUndoBeginReviseTicket(cmd UndoBeginReviseTicket) (events mskit.Events, err error) {
	switch t.State {
	case TicketState_RevisionPending:
		events = mskit.NewEventsSingle(
			cmd.Id,
			Ticket{},
			UndoTicketRevisionBegan{
				Id: cmd.Id,
			},
		)
	default:
		err = errbdr.NewErrUnsupportedStateTransition(t, t.State)
	}

	return
}

func (t *Ticket) processConfirmReviseTicket(cmd ConfirmReviseTicket) (events mskit.Events, err error) {
	switch t.State {
	case TicketState_RevisionPending:
		events = mskit.NewEventsSingle(
			cmd.Id,
			Ticket{},
			TicketRevisionConfirmed{
				Id:              cmd.Id,
				TicketLineItems: cmd.TicketLineItems,
			},
		)
	default:
		err = errbdr.NewErrUnsupportedStateTransition(t, t.State)
	}

	return
}

func (t *Ticket) Apply(event interface{}) error {
	switch e := event.(type) {
	case TicketCreated:
		return t.applyTicketCreated(e)
	case TicketCancelled:
		return t.applyTicketCancelled(e)
	case TicketConfirmed:
		return t.applyTicketConfirmed(e)
	case TicketRevisionBegan:
		return t.applyTicketRevisionBegan(e)
	case UndoTicketRevisionBegan:
		return t.applyUndoTicketRevisionBegan(e)
	case TicketRevisionConfirmed:
		return t.applyTicketRevisionConfirmed(e)
	default:
		return errbdr.NewErrUnknownParams(t.Apply, e)
	}
}

func (t *Ticket) applyTicketCreated(event TicketCreated) error {
	t.Id = event.Id
	t.RestaurantId = event.RestaurantId
	t.TicketLineItems = event.TicketLineItems

	t.PreviousState = TicketState_CreatePending
	t.State = TicketState_CreatePending

	return nil
}

func (t *Ticket) applyTicketCancelled(event TicketCancelled) error {
	t.PreviousState = t.State
	t.State = TicketState_Canceled

	return nil
}

func (t *Ticket) applyTicketConfirmed(event TicketConfirmed) error {
	t.PreviousState = t.State
	t.State = TicketState_AwaitingAcceptance

	return nil
}

func (t *Ticket) applyTicketRevisionBegan(event TicketRevisionBegan) error {
	t.PreviousState = t.State
	t.State = TicketState_RevisionPending

	return nil
}

func (t *Ticket) applyUndoTicketRevisionBegan(event UndoTicketRevisionBegan) error {
	t.State = t.PreviousState

	return nil
}

func (t *Ticket) applyTicketRevisionConfirmed(event TicketRevisionConfirmed) error {
	t.State = t.PreviousState
	t.TicketLineItems = event.TicketLineItems

	return nil
}
