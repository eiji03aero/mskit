package account

import (
	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/utils/errbdr"
)

type AccountAggregate struct {
	mskit.BaseAggregate
	Account *Account `json:"account"`
}

func NewAccountAggregate() *AccountAggregate {
	return &AccountAggregate{
		Account: &Account{},
	}
}

func (a *AccountAggregate) Validate() (errs []error) {
	return
}

func (a *AccountAggregate) Process(command interface{}) (mskit.Events, error) {
	switch cmd := command.(type) {
	case CreateAccount:
		return a.processCreateAccount(cmd)
	default:
		return mskit.Events{}, errbdr.NewErrUnknownParams(a.Process, cmd)
	}
}

func (a *AccountAggregate) processCreateAccount(command CreateAccount) (events mskit.Events, err error) {
	events = mskit.NewEventsSingle(
		command.Id,
		AccountAggregate{},
		AccountCreated{
			Id:         command.Id,
			ConsumerId: command.ConsumerId,
		},
	)
	return
}

func (a *AccountAggregate) Apply(event interface{}) error {
	switch e := event.(type) {
	case AccountCreated:
		return a.applyAccountCreated(e)
	default:
		return errbdr.NewErrUnknownParams(a.Apply, e)
	}
}

func (a *AccountAggregate) applyAccountCreated(event AccountCreated) (err error) {
	a.Id = event.Id
	a.Account.ConsumerId = event.ConsumerId
	return
}
