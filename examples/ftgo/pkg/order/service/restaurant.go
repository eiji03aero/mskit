package service

import (
	orderdmn "order/domain/order"
	restaurantdmn "order/domain/restaurant"

	"github.com/eiji03aero/mskit/utils/errbdr"
	"github.com/eiji03aero/mskit/utils/logger"
)

func (s *service) CreateRestaurant(restaurant restaurantdmn.Restaurant) (err error) {
	err = s.restaurantRepository.Save(restaurant)
	if err != nil {
		return
	}

	logger.PrintResourceCreated(restaurant)
	return
}

func (s *service) GetRestaurant(id string) (restaurant *restaurantdmn.Restaurant, err error) {
	restaurant, err = s.restaurantRepository.FindById(id)
	if err != nil {
		return
	}

	logger.PrintResourceGet(restaurant)
	return
}

func (s *service) validateMenuItems(restaurant *restaurantdmn.Restaurant, items orderdmn.OrderLineItems) (err error) {
	for _, item := range items.LineItems {
		_, found := restaurant.GetItemById(item.MenuItemId)
		if !found {
			return errbdr.NewErrDataNotFound(item, item.MenuItemId)
		}
	}
	return nil
}
