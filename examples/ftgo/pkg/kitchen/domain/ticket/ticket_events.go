package kitchen

type TicketCreated struct {
	Id              string          `json:"id"`
	RestaurantId    string          `json:"restaurant_id"`
	TicketLineItems TicketLineItems `json:"ticket_line_items"`
}

type TicketCancelled struct {
	Id string `json:"id"`
}
