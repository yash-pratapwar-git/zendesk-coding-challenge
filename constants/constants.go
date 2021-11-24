package constants

//Authorization variables
const (
	UserName string = "yxp200011@utdallas.edu"
	Password string = "Sonata@678"
)

//constants used in all packages are maintained here
const (
	WelcomeMessage      string = "\n**************   Welcome to the Ticket Viewer Application  *************\n"
	MainMenu            string = "\nPlease select an option :\n\nEnter 1 to List All Tickets\nEnter 2 to View Specific Ticket\nENter 3 to quit the application\n"
	ListAllTicketsURL   string = "https://zccutdallas.zendesk.com/api/v2/tickets.json?page=1&per_page=25"
	ViewSingleTicketURL string = "https://zccutdallas.zendesk.com/api/v2/tickets/{ticketID}.json"
	LoadingMessage      string = "Please wait while we fetch the data"
	ListViewMenu        string = "\nTHERE ARE MORE RESULTS.\n\nPlease enter your choice.\nEnter 1 to view next results\nEnter 2 to go back to main menu\n"
	InValidInputMessage string = "\nPlease provide digit input\n"
)

//constants used in unit tests
const (
	SampleAllTicketsData string = `{"tickets":[
		{
			"url": "sampleurl.com",
            "id": 1,
            "external_id": null,
            "via": {
                "channel": "api",
                "source": {
                    "from": {},
                    "to": {},
                    "rel": null
                }
            },
            "created_at": "2021-11-23T17:16:32Z",
            "updated_at": "2021-11-23T17:16:32Z",
            "type": null,
            "subject": "sample subject",
            "raw_subject": "sample subject",
            "description": "sample description",
            "priority": null,
            "status": "open",
            "recipient": null,
            "requester_id": 1,
            "submitter_id": 2,
            "assignee_id": 3,
            "organization_id": 4,
            "group_id": 5,
            "collaborator_ids": [],
            "follower_ids": [],
            "email_cc_ids": [],
            "forum_topic_id": null,
            "problem_id": null,
            "has_incidents": false,
            "is_public": true,
            "due_at": null,
            "tags": [
                "deserunt",
                "enim",
                "est"
            ],
            "custom_fields": [],
            "satisfaction_rating": null,
            "sharing_agreement_ids": [],
            "fields": [],
            "followup_ids": [],
            "ticket_form_id": 6,
            "brand_id": 7,
            "allow_channelback": false,
            "allow_attachments": true
		}
	],
	"next_page":null
	}`

	SampleSingleTicketResponse string = `
	{
		"url": "sampleurl.com",
		"id": 1,
		"external_id": null,
		"via": {
			"channel": "api",
			"source": {
				"from": {},
				"to": {},
				"rel": null
			}
		},
		"created_at": "2021-11-23T17:16:32Z",
		"updated_at": "2021-11-23T17:16:32Z",
		"type": null,
		"subject": "sample subject",
		"raw_subject": "sample subject",
		"description": "sample description",
		"priority": null,
		"status": "open",
		"recipient": null,
		"requester_id": 1,
		"submitter_id": 2,
		"assignee_id": 3,
		"organization_id": 4,
		"group_id": 5,
		"collaborator_ids": [],
		"follower_ids": [],
		"email_cc_ids": [],
		"forum_topic_id": null,
		"problem_id": null,
		"has_incidents": false,
		"is_public": true,
		"due_at": null,
		"tags": [
			"deserunt",
			"enim",
			"est"
		],
		"custom_fields": [],
		"satisfaction_rating": null,
		"sharing_agreement_ids": [],
		"fields": [],
		"followup_ids": [],
		"ticket_form_id": 6,
		"brand_id": 7,
		"allow_channelback": false,
		"allow_attachments": true
	}`
)
