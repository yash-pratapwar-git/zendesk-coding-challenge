package models

//struct for the ticket attributes
type TicketDetails struct {
	Id        int64  `json:"id"`
	Subject   string `json:"subject"`
	Submitter int64  `json:"submitter_id"`
	CreatedAt string `json:"created_at"`
}

//struct for the response of fetching all tickets
type TicketsList struct {
	Tickets  []TicketDetails `json:"tickets"`
	NextPage string          `json:"next_page"`
	Count    int             `json:"count"`
}

//struct for the response of fetching the single ticket
type SingleTicketResponse struct {
	Ticket TicketDetails `json:"ticket"`
}
