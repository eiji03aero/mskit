package service

import (
	consumerdmn "consumer/domain/consumer"

	"github.com/eiji03aero/mskit"
)

type service struct {
	eventRepository *mskit.EventRepository
}

type Service interface {
	CreateConsumer(cmd consumerdmn.CreateConsumer) (id string, err error)
	GetConsumer(id string) (consumer *consumerdmn.Consumer, err error)
	ValidateOrder(cmd consumerdmn.ValidateOrder) error
}

func New(r *mskit.EventRepository) Service {
	return &service{
		eventRepository: r,
	}
}
