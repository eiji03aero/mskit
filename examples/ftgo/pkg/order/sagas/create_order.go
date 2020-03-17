package sagas

import (
	"github.com/eiji03aero/mskit"
)

type CreateOrderSaga struct {
	mskit.Saga
}

func NewCreateOrderSaga() *CreateOrderSaga {
	saga := &CreateOrderSaga{}

	definition, err := mskit.NewSagaDefinitionBuilder().
		Step(
			mskit.SagaStepCompensationOption{
				handler: func() {},
			},
		).
		Step(
			mskit.SagaStepInvokeParticipantOption{
				handler: func() {},
			},
		).
		Step(
			mskit.SagaStepInvokeParticipantOption{
				handler: func() {},
			},
			mskit.SagaStepOnReplyOption{
				handler: func() {},
			},
			mskit.SagaStepCompensationOption{
				handler: func() {},
			},
		).
		Step(
			mskit.SagaStepInvokeParticipantOption{
				handler: func() {},
			},
		).
		Step(
			mskit.SagaStepInvokeParticipantOption{
				handler: func() {},
			},
		).
		Step(
			mskit.SagaStepInvokeParticipantOption{
				handler: func() {},
			},
		).
		Build()

	if err != nil {
		panic(err)
	}

	saga.Definition = definition

	return saga
}
