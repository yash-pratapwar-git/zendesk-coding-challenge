package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/zendesk-coding-challenge/constants"
	"github.com/zendesk-coding-challenge/models"
	"github.com/zendesk-coding-challenge/zendeskerrors"
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

//GetHttpMethod : function to make HTTP calls with GET method
func GetHttpMethod(url string) (body io.ReadCloser, err error) {
	//creating new http request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error in creating new request : ", err)
		return nil, errors.New(zendeskerrors.ERR102)
	}

	//setting authorization headers
	req.SetBasicAuth(constants.UserName, constants.Password)

	//making the actual request
	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	//checking if the user is authorized
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, errors.New(zendeskerrors.ERR101)
	}

	//handling all the status codes other than 200 OK
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Could not connect to the API")
		fmt.Println("Reason : ", resp.Status)

		return nil, errors.New(zendeskerrors.ERR103)
	}

	return resp.Body, nil
}

//ListAllData : function to print all the available tickets with pagination
func ListAllData(url string) (err error) {
	fmt.Println("\n", constants.LoadingMessage)
	var tickets models.TicketsList

	body, err := GetHttpMethod(url)
	if err != nil {
		fmt.Println(" Error Occurred : ", err)
		return err
	}
	defer body.Close()

	//unmarshalling the response into Go structs
	err = json.NewDecoder(body).Decode(&tickets)
	if err != nil {
		return err
	}
	fmt.Println("\nTOTAL NUMBER OF TICKETS : ", tickets.Count)
	fmt.Println()

	//displaying the tickets list
	for _, val := range tickets.Tickets {
		fmt.Println(val.Id, ". Ticket with subject ", val.Subject, " created by ", val.Submitter, " at ", val.CreatedAt)
	}

	//If there are more results available (pagination support)
	if tickets.NextPage != "" {

		var exitFlag bool

		for !exitFlag {
			fmt.Println(constants.ListViewMenu)

			var response int
			_, err := fmt.Scanf("%d", &response)
			if err != nil {
				fmt.Println("Please enter digit inputs")
			}

			switch response {
			case 1:
				ListAllData(tickets.NextPage)
			case 2:
				return nil
			default:
				fmt.Println(constants.InValidInputMessage)
			}
		}

	}
	return nil
}

// SpecificTicketInfo : function to print the information of single ticket
func SpecificTicketInfo(url string, ticketId int) (err error) {
	url = strings.Replace(url, "{ticketID}", strconv.Itoa(ticketId), 1)

	body, err := GetHttpMethod(url)
	if err != nil {
		fmt.Println(" Error Occurred : ", err)
		return
	}
	defer body.Close()

	var ticket models.SingleTicketResponse

	//unmarshalling the response into Go structs
	err = json.NewDecoder(body).Decode(&ticket)
	if err != nil {
		return err
	}

	//Displaying the single ticket information
	fmt.Println(ticket.Ticket.Id, ". Ticket with subject ", ticket.Ticket.Subject, " created by ", ticket.Ticket.Submitter, " at ", ticket.Ticket.CreatedAt)

	return nil
}
