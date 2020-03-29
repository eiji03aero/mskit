package kitchen

import (
	ticketdmn "kitchen/domain/ticket"
)

type Service interface {
	CreateTicket(cmd ticketdmn.CreateTicket) (id string, err error)
	CancelTicket(cmd ticketdmn.CancelTicket) (err error)
}
