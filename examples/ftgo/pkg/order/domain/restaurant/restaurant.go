package restaurant

type Restaurant struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	RestaurantMenu RestaurantMenu `json:"restaurant_menu"`
}

type RestaurantMenu struct {
	MenuItems []RestaurantMenuItem `json:"menu_items"`
}

type RestaurantMenuItem struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
