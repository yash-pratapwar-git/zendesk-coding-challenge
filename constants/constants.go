package constants

const (
	WelcomeMessage      string = "\n**************   Welcome to the Ticket Viewer Application  *************\n"
	MainMenu            string = "\nPlease select an option :\n\nEnter 1 to List All Tickets\nEnter 2 to View Specific Ticket\nENter 3 to quit the application\n"
	ListAllTicketsURL   string = "https://zccutdallas.zendesk.com/api/v2/tickets.json?page=1&per_page=25"
	ViewSingleTicketURL string = "https://zccutdallas.zendesk.com/api/v2/tickets/{ticketID}.json"
	LoadingMessage      string = "Please wait while we fetch the data"
	ListViewMenu        string = "\nTHERE ARE MORE RESULTS.\n\nPlease enter your choice.\nEnter 1 to view next results\nEnter 2 to go back to main menu\n"
	InValidInputMessage string = "\nPlease provide digit input\n"
)

const (
	UserName string = "yxp200011@utdallas.edu"
	Password string = "Sonata@678"
)
