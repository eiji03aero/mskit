package accounting

import (
	accountdmn "accounting/domain/account"
)

type Service interface {
	CreateAccount(cmd accountdmn.CreateAccount) (id string, err error)
	Authorize(consumerId string) (err error)
}
