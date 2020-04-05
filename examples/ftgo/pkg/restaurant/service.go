package restaurant

import (
	restaurantdmn "restaurant/domain/restaurant"
)

type Service interface {
	CreateRestaurant(cmd restaurantdmn.CreateRestaurant) (id string, err error)
}
