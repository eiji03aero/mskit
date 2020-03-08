package service

import (
	logcommon "common/log"
	restaurantdmn "restaurant/domain/restaurant"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/utils"
)

type service struct {
	repository *mskit.Repository
	publisher  mskit.DomainEventPublisher
}

type Service interface {
	CreateRestaurant(cmd restaurantdmn.CreateRestaurant) (id string, err error)
}

func New(r *mskit.Repository, p mskit.DomainEventPublisher) Service {
	return &service{
		repository: r,
		publisher:  p,
	}
}

func (s *service) CreateRestaurant(cmd restaurantdmn.CreateRestaurant) (id string, err error) {
	id, err = utils.UUID()
	if err != nil {
		return
	}

	restaurant := &restaurantdmn.Restaurant{}
	cmd.Id = id

	s.repository.ExecuteCommand(restaurant, cmd)
	if err != nil {
		return
	}

	logcommon.PrintCreated(restaurant)
	return
}
