package mskit

type SagaInstanceState int

const (
	SagaInstanceState_Unknown SagaInstanceState = iota
	SagaInstanceState_Processing
	SagaInstanceState_Aborted
	SagaInstanceState_Done
)

type SagaInstance struct {
	Id        string            `json:"id"`
	SagaState SagaInstanceState `json:"saga_state"`
	Data      interface{}       `json:"data"`
}
