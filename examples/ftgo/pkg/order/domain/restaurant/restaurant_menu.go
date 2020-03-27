package restaurant

type RestaurantMenu struct {
	MenuItems []RestaurantMenuItem `json:"menu_items"`
}

func (rm *RestaurantMenu) GetItemById(id string) (item RestaurantMenuItem, found bool) {
	found = false

	for _, menuItem := range rm.MenuItems {
		if menuItem.Id == id {
			item = menuItem
			found = true
			return
		}
	}

	return
}

func (rm *RestaurantMenu) Len() int {
	return len(rm.MenuItems)
}
