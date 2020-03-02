package service

import (
	logcommon "common/log"
	"github.com/eiji03aero/mskit"
	"github.com/eiji03aero/mskit/utils"
	restaurantdmn "restaurant/domain/restaurant"
)

type service struct {
	repository *mskit.Repository
}

type Service interface {
	CreateRestaurant(cmd restaurantdmn.CreateRestaurant) (id string, err error)
}

func New(r *mskit.Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateRestaurant(cmd restaurantdmn.CreateRestaurant) (id string, err error) {
	id, err = utils.UUID()
	if err != nil {
		return
	}

	restaurant := &restaurantdmn.Restaurant{}
	restaurant.Id = id

	events, err := restaurant.Process(cmd)
	if err != nil {
		return
	}

	for _, e := range events {
		err = s.repository.Save(restaurant, e)
		if err != nil {
			return
		}
	}

	logcommon.PrintJsonln("Restaurant created: ")
	logcommon.PrintJsonln(restaurant)
	return
}
