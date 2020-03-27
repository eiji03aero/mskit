package createorder

import (
	"errors"
	"order/transport/proxy"

	"github.com/eiji03aero/mskit"
)

func NewManager(
	repository *mskit.SagaRepository,
	pxy *proxy.Proxy,
) mskit.SagaManager {
	definition := createDefinition(
		pxy,
	)
	return mskit.NewSagaManager(
		definition,
		repository,
	)
}

func createDefinition(
	pxy *proxy.Proxy,
) *mskit.SagaDefinition {
	definition, err := mskit.NewSagaDefinitionBuilder().
		Step(
			mskit.SagaStepCompensationOption{
				Handler: func(ss interface{}) (err error) {
					sagaState, err := assertStruct(ss)
					if err != nil {
						return
					}

					err = pxy.RejectOrder(sagaState.OrderId)

					return
				},
			},
		).
		Step(
			mskit.SagaStepExecuteOption{
				Handler: func(_ interface{}) (err error) {
					err = errors.New("shippai")
					return
				},
			},
		).
		Build()
	if err != nil {
		panic(err)
	}

	return definition
}
