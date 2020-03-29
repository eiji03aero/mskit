package kitchen

type TicketLineItem struct {
	Quantity   int    `json:"quantity"`
	MenuItemId string `json:"menu_item_id"`
}
