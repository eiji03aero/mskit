package mskit

// SagaStep is a struct to express one action of saga
type SagaStep struct {
	invokeParticipantHandler func(sagaState interface{}) (result *SagaStepResult)
	compensationHandler      func(sagaState interface{}) (result *SagaStepResult)
}

// SagaStepInvokeParticipantOption is struct to set InvokeParticipant in SagaAction
type SagaStepInvokeParticipantOption struct {
	Handler func(sagaState interface{})
}

// SagaStepCompensationOption is struct to set Compensation in SagaAction
type SagaStepCompensationOption struct {
	Handler func(sagaState interface{})
}

// Configure sets options
func (sa *SagaStep) Configure(opts ...interface{}) {
	for _, opt := range opts {
		switch o := opt.(type) {
		case SagaStepInvokeParticipantOption:
			sa.invokeParticipantHandler = o.Handler
		case SagaStepCompensationOption:
			sa.compensationHandler = o.Handler
		}
	}
}

// Validate returns error if SagaStep is not valid yet
func (sa *SagaStep) Validate() error {
	return nil
}

type SagaStepResult struct {
	Id    string
	Error error
}
