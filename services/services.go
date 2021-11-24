package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/zendesk-coding-challenge/models"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

func GetHttpMethod(url string) (body io.ReadCloser, err error) {
	// client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error in creating new request : ", err)
		return nil, err
	}

	req.SetBasicAuth("yxp200011@utdallas.edu", "Sonata@678")
	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Could not connect to the API")
		fmt.Println("Reason : ", resp.Status)

		return nil, errors.New("unsucessfull http call")
	}

	return resp.Body, nil
}

func ListAllData(url string) (err error) {
	fmt.Println("\nPlease wait while we fetch the data")
	var tickets models.TicketsList

	body, err := GetHttpMethod(url)
	if err != nil {
		fmt.Println(" Error Occurred : ", err)
		return err
	}
	defer body.Close()
	err = json.NewDecoder(body).Decode(&tickets)
	if err != nil {
		return err
	}
	fmt.Println("\nTOTAL NUMBER OF TICKETS : ", tickets.Count)
	fmt.Println()

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
			return nil
		default:
			fmt.Println("\nPlease provide valid option")
			fmt.Println("")
		}
	}
	return nil
}

func SpecificTicketInfo(url string, ticketId int) (err error) {
	url = strings.Replace(url, "{ticketID}", strconv.Itoa(ticketId), 1)

	body, err := GetHttpMethod(url)
	if err != nil {
		fmt.Println(" Error Occurred : ", err)
		return
	}
	defer body.Close()

	var ticket models.SingleTicketResponse

	err = json.NewDecoder(body).Decode(&ticket)
	if err != nil {
		return err
	}

	fmt.Println(ticket.Ticket.Id, ". Ticket with subject ", ticket.Ticket.Subject, " created by ", ticket.Ticket.Submitter, " at ", ticket.Ticket.CreatedAt)

	return nil
}
