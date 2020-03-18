package mskit

// SagaInstanceState defines state of SagaInstance
type SagaInstanceState int

const (
	SagaInstanceState_Unknown SagaInstanceState = iota
	SagaInstanceState_Processing
	SagaInstanceState_Aborted
	SagaInstanceState_Done
)

// SagaInstace is a struct to express an instance of saga
type SagaInstance struct {
	Id        string            `json:"id"`
	SagaState SagaInstanceState `json:"saga_state"`
	Data      interface{}       `json:"data"`
}
