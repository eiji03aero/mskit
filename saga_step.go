package mskit

// SagaStep is a struct to express one action of saga
type SagaStep struct {
	invokeParticipantHandler func()
	replyHandler             func()
	compensationHandler      func()
}

// SagaStepInvokeParticipantOption is struct to set InvokeParticipant in SagaAction
type SagaStepInvokeParticipantOption struct {
	handler func()
}

// SagaStepOnReplyOption is struct to set OnReply in SagaAction
type SagaStepOnReplyOption struct {
	handler func()
}

// SagaStepCompensationOption is struct to set Compensation in SagaAction
type SagaStepCompensationOption struct {
	handler func()
}

// Configure sets options
func (sa *SagaStep) Configure(opts ...interface{}) {
	for _, opt := range opts {
		switch o := opt.(type) {
		case SagaStepInvokeParticipantOption:
			sa.invokeParticipantHandler = o.handler
		case SagaStepOnReplyOption:
			sa.replyHandler = o.handler
		case SagaStepCompensationOption:
			sa.compensationHandler = o.handler
		}
	}
}

// Validate returns error if SagaStep is not valid yet
func (sa *SagaStep) Validate() error {
	return nil
}
