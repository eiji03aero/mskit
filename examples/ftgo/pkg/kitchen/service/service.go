package service

import (
	"kitchen"
	ticketdmn "kitchen/domain/ticket"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/utils"
)

type service struct {
	eventRepository *mskit.EventRepository
}

func New(r *mskit.EventRepository) kitchen.Service {
	return &service{
		eventRepository: r,
	}
}

func (s *service) CreateTicket(cmd ticketdmn.CreateTicket) (id string, err error) {
	id, err = utils.UUID()
	if err != nil {
		return
	}

	ticket := &ticketdmn.Ticket{}
	cmd.Id = id

	err = s.eventRepository.ExecuteCommand(ticket, cmd)

	return
}

func (s *service) GetTicket(id string) (ticket *ticketdmn.Ticket, err error) {
	ticket = &ticketdmn.Ticket{}

	err = s.eventRepository.Load(id, ticket)
	if err != nil {
		return
	}

	return
}

func (s *service) CancelTicket(cmd ticketdmn.CancelTicket) (err error) {
	ticket, err := s.GetTicket(cmd.Id)
	if err != nil {
		return
	}

	err = s.eventRepository.ExecuteCommand(ticket, cmd)

	return
}

func (s *service) ConfirmTicket(cmd ticketdmn.ConfirmTicket) (err error) {
	ticket, err := s.GetTicket(cmd.Id)
	if err != nil {
		return
	}

	err = s.eventRepository.ExecuteCommand(ticket, cmd)

	return
}

func (s *service) BeginReviseTicket(cmd ticketdmn.BeginReviseTicket) (err error) {
	ticket, err := s.GetTicket(cmd.Id)
	if err != nil {
		return
	}

	err = s.eventRepository.ExecuteCommand(ticket, cmd)

	return
}

func (s *service) UndoBeginReviseTicket(cmd ticketdmn.UndoBeginReviseTicket) (err error) {
	ticket, err := s.GetTicket(cmd.Id)
	if err != nil {
		return
	}

	err = s.eventRepository.ExecuteCommand(ticket, cmd)

	return
}

func (s *service) ConfirmReviseTicket(cmd ticketdmn.ConfirmReviseTicket) (err error) {
	ticket, err := s.GetTicket(cmd.Id)
	if err != nil {
		return
	}

	err = s.eventRepository.ExecuteCommand(ticket, cmd)

	return
}
