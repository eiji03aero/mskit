package kitchen

type CreateTicket struct {
	Id              string          `json:"id"`
	RestaurantId    string          `json:"restaurant_id"`
	TicketLineItems TicketLineItems `json:"ticket_line_items"`
}

type CancelTicket struct {
	Id string `json:"id"`
}

type ConfirmTicket struct {
	Id string `json:"id"`
}
