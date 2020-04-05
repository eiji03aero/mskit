package kitchen

type TicketState int

const (
	TicketState_Unknown TicketState = iota
	TicketState_CreatePending
	TicketState_AwaitingAcceptance
	TicketState_Accepted
	TicketState_Preparing
	TicketState_ReadyForPickup
	TicketState_PickedUp
	TicketState_CancelPending
	TicketState_Canceled
	TicketState_RevisionPending
)
