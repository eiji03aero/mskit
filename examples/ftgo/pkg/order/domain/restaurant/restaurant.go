package restaurant

type Restaurant struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	RestaurantMenu RestaurantMenu `json:"restaurant_menu"`
}

func (r *Restaurant) GetItemById(id string) (item RestaurantMenuItem, found bool) {
	return r.RestaurantMenu.GetItemById(id)
}
