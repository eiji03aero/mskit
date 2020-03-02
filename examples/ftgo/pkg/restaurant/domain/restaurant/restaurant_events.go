package restaurant

type RestaurantCreated struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	RestaurantMenu RestaurantMenu `json:"restaurant_menu"`
}
