package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/zendesk-coding-challenge/models"
)

func ListAllData(url string) {
	fmt.Println("\nPlease wait while we fetch the data")
	var tickets models.TicketsList
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error in creating new request : ", err)
	}

	req.SetBasicAuth("yxp200011@utdallas.edu", "Sonata@678")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&tickets)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(tickets.NextPage))
	for _, val := range tickets.Tickets {
		fmt.Println(val.Id, ". Ticket with subject ", val.Subject, " created by ", val.Submitter, " at ", val.CreatedAt)
	}

	if tickets.NextPage != "" {
		fmt.Println("\nThere are more results. Enter your choice.")
		fmt.Println("\n1. View Next Results")
		fmt.Println("2. Main Menu")

		var response int
		_, err := fmt.Scanf("%d", &response)
		if err != nil {
			fmt.Println("Please enter digit inputs")
		}

		switch response {
		case 1:
			fmt.Println("\nOption 1 selected")
			ListAllData(tickets.NextPage)
		case 2:
			return
		default:
			fmt.Println("\nPlease provide valid option")
			fmt.Println("")
		}
	}
}
