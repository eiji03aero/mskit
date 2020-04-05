package service

import (
	consumerroot "consumer"

	"github.com/eiji03aero/mskit"
)

type service struct {
	eventRepository *mskit.EventRepository
}

func New(r *mskit.EventRepository) consumerroot.Service {
	return &service{
		eventRepository: r,
	}
}
