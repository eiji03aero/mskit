package restaurant

type CreateRestaurant struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	RestaurantMenu RestaurantMenu `json:"restaurant_menu"`
}
