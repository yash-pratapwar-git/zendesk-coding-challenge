package models

type Ticket struct {
	Id        int64  `json:"id"`
	Subject   string `json:"subject"`
	Submitter int64  `json:"submitter_id"`
	CreatedAt string `json:"created_at"`
}

type TicketsList struct {
	Tickets  []Ticket `json:"tickets"`
	NextPage string   `json:"next_page"`
	Count    int      `json:"count"`
}
