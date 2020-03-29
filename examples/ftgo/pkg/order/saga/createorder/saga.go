package createorder

import (
	"fmt"
)

type state struct {
	OrderId string `json:"order_id"`
}

func NewState(id string) *state {
	return &state{
		OrderId: id,
	}
}

func assertStruct(value interface{}) (s *state, err error) {
	s, ok := value.(*state)
	if !ok {
		err = fmt.Errorf("createorder.assertStruct: invalid data", value)
		return
	}

	return
}
