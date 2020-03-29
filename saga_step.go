package mskit

import (
	"github.com/eiji03aero/mskit/utils/errbdr"
)

// SagaStepResult expresses the result of SagaStepHandler
type SagaStepResult struct {
	Id    string
	Error error
}

// SagaStepHandler defines the signature for handle function for step
type SagaStepHandler func(sagaInstance *SagaInstance) (err error)

// SagaStep is a struct to express one action of saga
type SagaStep struct {
	executeHandler      SagaStepHandler
	compensationHandler SagaStepHandler
}

// SagaStepExecuteOption is struct to set executeHandler in SagaStep
type SagaStepExecuteOption struct {
	Handler SagaStepHandler
}

// SagaStepCompensationOption is struct to set compensationHandler in SagaStep
type SagaStepCompensationOption struct {
	Handler SagaStepHandler
}

// Configure sets options
func (sa *SagaStep) Configure(opts ...interface{}) {
	for _, opt := range opts {
		switch o := opt.(type) {
		case SagaStepExecuteOption:
			sa.executeHandler = o.Handler
		case SagaStepCompensationOption:
			sa.compensationHandler = o.Handler
		default:
			panic(errbdr.NewErrUnknownParams(sa.Configure, o))
		}
	}
}

// Validate returns error if SagaStep is not valid yet
func (sa *SagaStep) Validate() error {
	return nil
}
