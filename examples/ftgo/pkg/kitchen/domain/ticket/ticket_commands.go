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

type BeginReviseTicket struct {
	Id              string          `json:"id"`
	TicketLineItems TicketLineItems `json:"ticket_line_items"`
}

type UndoBeginReviseTicket struct {
	Id string `json:"id"`
}

type ConfirmReviseTicket struct {
	Id              string          `json:"id"`
	TicketLineItems TicketLineItems `json:"ticket_line_items"`
}
