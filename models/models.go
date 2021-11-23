package models

type TicketDetails struct {
	Id        int64  `json:"id"`
	Subject   string `json:"subject"`
	Submitter int64  `json:"submitter_id"`
	CreatedAt string `json:"created_at"`
}

type TicketsList struct {
	Tickets  []TicketDetails `json:"tickets"`
	NextPage string          `json:"next_page"`
	Count    int             `json:"count"`
}

type SingleTicketResponse struct {
	Ticket TicketDetails `json:"ticket"`
}
