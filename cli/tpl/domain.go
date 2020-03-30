package tpl

func DomainEntityTemplate() string {
	return `package {{ .LowerName }}

type {{ .Name }} struct {
}`
}

func DomainAggregateTemplate() string {
	return `package {{ .LowerName }}

import (
	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/utils/errbdr"
)

type {{ .AggregateName }} struct {
	mskit.BaseAggregate
	{{ .Name }} *{{ .Name }}
}

func New{{ .AggregateName }}() *{{ .AggregateName }} {
	return &{{ .AggregateName }}{
		{{ .Name }}: &{{ .Name }}{}
	}
}

func ({{ .NameInitial }} *{{ .AggregateName }}) Validate() (errs []error) {
	return
}

func ({{ .NameInitial }} *{{ .AggregateName }}) Process(command interface{}) (mskit.Events, error) {
	switch cmd := command.(type) {
	default:
		return mskit.Events{}, errbdr.NewErrUnknownParams({{ .NameInitial }}.Process, cmd)
	}
}

func ({{ .NameInitial }} *{{ .AggregateName }}) Apply(event interface{}) error {
	switch e := event.(type) {
	default:
		return errbdr.NewErrUnknownParams({{ .NameInitial }}.Apply, e)
	}
}`
}

func DomainCommandsTemplate() string {
	return `package {{ .LowerName }}
`
}

func DomainEventsTemplate() string {
	return `package {{ .LowerName }}
`
}
