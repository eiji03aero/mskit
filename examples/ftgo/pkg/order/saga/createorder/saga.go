package createorder

import "encoding/json"

type state struct {
	OrderId string `json:"order_id"`
}

func NewState(id string) *state {
	return &state{
		OrderId: id,
	}
}

func assertStruct(value interface{}) (s *state, err error) {
	s = &state{}
	str, ok := value.(string)
	if !ok {
		return
	}

	err = json.Unmarshal([]byte(str), s)

	return
}
