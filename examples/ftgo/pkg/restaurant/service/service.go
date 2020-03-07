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
	cmd.Id, err = utils.UUID()
	if err != nil {
		return
	}

	restaurant := &restaurantdmn.Restaurant{}
	events, err := restaurant.Process(cmd)
	if err != nil {
		return
	}

	for _, e := range events {
		err = s.repository.Save(restaurant, e)
		if err != nil {
			return
		}

		err = s.publisher.Publish(e.Data)
		if err != nil {
			return
		}
	}

	logcommon.PrintCreated(restaurant)
	return
}
