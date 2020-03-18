package service

import (
	logcommon "common/log"
	kitchendmn "kitchen/domain/kitchen"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/utils"
)

type service struct {
	repository *mskit.EventRepository
}

type Service interface {
	CreateTicket(cmd kitchendmn.CreateTicket) (id string, err error)
}

func New(r *mskit.EventRepository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateTicket(cmd kitchendmn.CreateTicket) (id string, err error) {
	id, err = utils.UUID()
	if err != nil {
		return id, err
	}

	ticket := &kitchendmn.Ticket{}
	cmd.Id = id

	err = s.repository.ExecuteCommand(ticket, cmd)

	logcommon.PrintCreated(ticket)
	return
}
