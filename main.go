package main

import (
	"fmt"

	"github.com/zendesk-coding-challenge/constants"
	"github.com/zendesk-coding-challenge/services"
)

func main() {
	fmt.Println(constants.WelcomeMessage)

	var exitFlag bool

	//Maintaining the main menu message operation
	for !exitFlag {

		fmt.Println(constants.MainMenu)
		var response int
		fmt.Scanf("%d", &response)

		//switch case to handle the user input
		switch response {
		case 1:
			services.ListAllData(constants.ListAllTicketsURL)
		case 2:
			fmt.Println("Enter the Ticket ID ")
			var ticketId int
			_, err := fmt.Scanf("%d", &ticketId)
			if err != nil {
				fmt.Println(constants.InValidInputMessage)
				continue
			}
			services.SpecificTicketInfo(constants.ViewSingleTicketURL, ticketId)
		case 3:
			exitFlag = true
		default:
			fmt.Println("\nPlease provide valid option")
		}
	}
}
