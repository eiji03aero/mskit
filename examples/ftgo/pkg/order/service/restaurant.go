package service

import (
	errorscommon "common/errors"
	logcommon "common/log"
	orderdmn "order/domain/order"
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

func (s *service) validateMenuItems(restaurant *restaurantdmn.Restaurant, items orderdmn.OrderLineItems) (err error) {
	for _, item := range items.LineItems {
		_, found := restaurant.GetItemById(item.MenuItemId)
		if !found {
			return errorscommon.NewErrDataNotFound(item, item.MenuItemId)
		}
	}
	return nil
}
