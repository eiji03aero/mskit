package service

import (
	logcommon "common/log"
	restaurantdmn "order/domain/restaurant"
)

func (s *service) CreateRestaurant(restaurant restaurantdmn.Restaurant) (err error) {
	err = s.restaurantRepository.Save(restaurant)
	if err != nil {
		return
	}

	logcommon.PrintCreated(restaurant)
	return
}

func (s *service) GetRestaurant(id string) (restaurant *restaurantdmn.Restaurant, err error) {
	restaurant, err = s.restaurantRepository.FindById(id)
	if err != nil {
		return
	}
	return
}
