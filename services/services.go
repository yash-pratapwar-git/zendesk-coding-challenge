package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

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
		return nil, errors.New(zendeskerrors.ERR104)
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

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, constants.TableHeaders)

	//displaying the tickets list
	for _, val := range tickets.Tickets {
		fmt.Fprintln(w, val.Id, "\t", val.Subject, "\t", val.Submitter, "\t", val.CreatedAt, "\t", val.Status, "\t")
	}
	w.Flush()

	//If there are more results available (pagination support)
	if tickets.NextPage != "" {
		fmt.Println(constants.ListViewMenu)

		var response int
		_, err := fmt.Scanf("%d", &response)
		if err != nil {
			response = 3
		}

		switch response {
		case 1:
			ListAllData(tickets.NextPage)
		case 2:
			return nil
		default:
			fmt.Println(constants.InValidInputMessage)
		}

	} else {
		fmt.Println(constants.EndOfResultMessage)
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

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	//Displaying the single ticket information
	fmt.Fprintln(w, "ID : \t", ticket.Ticket.Id)
	fmt.Fprintln(w, "SUBJECT : \t", ticket.Ticket.Subject)
	fmt.Fprintln(w, "SUBMITTED BY : \t", ticket.Ticket.Submitter)
	fmt.Fprintln(w, "CREATED AT : \t", ticket.Ticket.CreatedAt)
	fmt.Fprintln(w, "STATUS : \t", ticket.Ticket.Status)
	w.Flush()

	return nil
}
