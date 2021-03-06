package service

import (
	restaurantroot "restaurant"
	restaurantdmn "restaurant/domain/restaurant"

	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/utils"
)

type service struct {
	eventRepository *mskit.EventRepository
	publisher       mskit.EventPublisher
}

func New(r *mskit.EventRepository, p mskit.EventPublisher) restaurantroot.Service {
	return &service{
		eventRepository: r,
		publisher:       p,
	}
}

func (s *service) CreateRestaurant(cmd restaurantdmn.CreateRestaurant) (id string, err error) {
	id, err = utils.UUID()
	if err != nil {
		return
	}

	restaurant := &restaurantdmn.Restaurant{}
	cmd.Id = id

	s.eventRepository.ExecuteCommand(restaurant, cmd)
	if err != nil {
		return
	}

	return
}
