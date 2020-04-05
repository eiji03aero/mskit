package tpl

func SagaManagerTemplate() string {
	return `package {{ .LowerName }}

import (
	"{{ .PkgName }}"

	"github.com/eiji03aero/mskit"
)

type client struct {
	sagaRepository  *mskit.SagaRepository
	service         {{ .PkgName }}.Service
}

func NewManager(
	sagaRepository *mskit.SagaRepository,
	svc {{ .PkgName }}.Service,
) mskit.SagaManager {
	c := &client{
		sagaRepository:      sagaRepository,
		service:         svc,
	}

	definition, err := mskit.NewSagaDefinitionBuilder().
		Step(
			mskit.SagaStepExecuteOption{
				Handler: c.sampleE,
			},
		).
		Build()
	if err != nil {
		panic(err)
	}

	return mskit.NewSagaManager(
		sagaRepository,
		definition,
		state{},
	)
}

func (c *client) sampleE(si *mskit.SagaInstance) (err error) {
	sagaState, err := assertStruct(si.Data)
	if err != nil {
		return
	}

	return
}`
}

func SagaStateTemplate() string {
	return `package {{ .LowerName }}

import (
	"github.com/eiji03aero/mskit/utils/errbdr"
)

type state struct {}

func NewState() *state {
	return &state{}
}

func assertStruct(value interface{}) (s *state, err error) {
	s, ok := value.(*state)
	if !ok {
		err = errbdr.NewErrUnknownParams(assertStruct, s)
		return
	}

	return
}`
}
