package service

import (
	ticketdmn "kitchen/domain/ticket"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/utils"
	"github.com/eiji03aero/mskit/utils/logger"
)

type service struct {
	repository *mskit.EventRepository
}

type Service interface {
	CreateTicket(cmd ticketdmn.CreateTicket) (id string, err error)
}

func New(r *mskit.EventRepository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateTicket(cmd ticketdmn.CreateTicket) (id string, err error) {
	id, err = utils.UUID()
	if err != nil {
		return id, err
	}

	ticket := &ticketdmn.Ticket{}
	cmd.Id = id

	err = s.repository.ExecuteCommand(ticket, cmd)

	logger.PrintResourceCreated(ticket)
	return
}
