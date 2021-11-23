package main

import (
	"fmt"

	"github.com/zendesk-coding-challenge/services"
)

func main() {
	fmt.Println("\n**************   Welcome to the Ticket Viewer Application  *************")

	var exitFlag bool

	for !exitFlag {

		fmt.Println("\nSelect an option :")
		fmt.Println("\n1. List All Tickets")
		fmt.Println("2. View Specific Ticket")
		fmt.Println("3. Quit Appliaction")
		fmt.Println("")
		var response int
		_, err := fmt.Scanf("%d", &response)
		if err != nil {
			fmt.Println("Please enter digit inputs")
		}

		switch response {
		case 1:
			fmt.Println("\nOption 1 selected")
			services.ListAllData("https://zccutdallas.zendesk.com/api/v2/tickets.json?page=1&per_page=25")
		case 2:
			fmt.Println("\nOption 2 selected")
		case 3:
			exitFlag = true
		default:
			fmt.Println("\nPlease provide valid option")
			fmt.Println("")
		}
	}
}
