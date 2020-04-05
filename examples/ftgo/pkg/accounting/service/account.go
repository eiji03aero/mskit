package service

import (
	"fmt"
	"math/rand"
	"time"

	accountdmn "accounting/domain/account"

	"github.com/eiji03aero/mskit/utils"
)

func (s *service) CreateAccount(cmd accountdmn.CreateAccount) (id string, err error) {
	id, err = utils.UUID()
	if err != nil {
		return
	}

	cmd.Id = id
	accountAggregate := accountdmn.NewAccountAggregate()

	err = s.eventRepository.ExecuteCommand(accountAggregate, cmd)
	if err != nil {
		return
	}

	return
}

func (s *service) Authorize(consumerId string) (err error) {
	// simulate authorize logic with random number
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(10)

	if i < 2 {
		return fmt.Errorf("authorize failed")
	}

	return
}
