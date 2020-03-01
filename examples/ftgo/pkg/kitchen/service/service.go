package service

import (
	"log"

	logcommon "common/log"
	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/utils"
	kitchendmn "kitchen/domain/kitchen"
)

type service struct {
	repository *mskit.Repository
}

type Service interface {
	CreateTicket(cmd kitchendmn.CreateTicket) (id string, err error)
}

func New(r *mskit.Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateTicket(cmd kitchendmn.CreateTicket) (id string, err error) {
	id, err = utils.UUID()
	if err != nil {
		return id, err
	}

	cmd.Id = id

	ticket := &kitchendmn.Ticket{}
	events, err := ticket.Process(cmd)
	if err != nil {
		return
	}

	for _, event := range events {
		err = s.repository.Save(ticket, event)
		if err != nil {
			return
		}
	}

	log.Println("ticket created: ")
	logcommon.PrintJsonln(ticket)
	return
}
