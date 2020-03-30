package service

import (
	"accounting"

	"github.com/eiji03aero/mskit"
)

type service struct {
	eventRepository *mskit.EventRepository
}

func New(
	er *mskit.EventRepository,
) accounting.Service {
	return &service{
		eventRepository: er,
	}
}
