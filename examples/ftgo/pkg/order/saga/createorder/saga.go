package createorder

import (
	"github.com/eiji03aero/mskit/utils/errbdr"
)

type state struct {
	OrderId  string `json:"order_id"`
	TicketId string `json:"ticket_id"`
}

func NewState(id string) *state {
	return &state{
		OrderId: id,
	}
}

func assertStruct(value interface{}) (s *state, err error) {
	s, ok := value.(*state)
	if !ok {
		err = errbdr.NewErrUnknownParams(assertStruct, s)
		return
	}

	return
}
