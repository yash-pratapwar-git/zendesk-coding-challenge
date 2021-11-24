package services

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
)

// Custom type that allows setting the func that our Mock Do func will run instead
type MockDoType func(req *http.Request) (*http.Response, error)

// MockClient is the mock client
type MockClient struct {
	MockDo MockDoType
}

// Overriding what the Do function should "do" in our MockClient
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}
func TestGetHttpMethodSuccess(t *testing.T) {
	// build our response JSON
	jsonResponse := `[{
   "full_name": "mock-repo"
  }]`
	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}
	_, err := GetHttpMethod("sampleURL")
	if err != nil {
		t.Error("TestGetHttpMethod failed.")
		return
	}
}

func TestGetHttpMethodFailed(t *testing.T) {
	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 404,
				Body:       nil,
			}, nil
		},
	}
	_, err := GetHttpMethod("sampleURL")
	if err == nil {
		t.Error("TestGetHttpMethodFail failed.")
		return
	}
}

func TestGetHttpMethodClientfail(t *testing.T) {
	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 404,
				Body:       nil,
			}, errors.New("Mock Error")
		},
	}
	_, err := GetHttpMethod("sampleURL")
	if err == nil {
		t.Error("TestGetHttpMethodClientfail failed.")
		return
	}
}

func TestListAllData(t *testing.T) {
	// build our response JSON
	jsonResponse := `{"tickets":[
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
	"next_page":"nextPageUrl.com"
	}`
	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}
	err := ListAllData("sampleURL")
	if err != nil {
		t.Error("TestGetHttpMethod failed.")
		return
	}
}

func TestListAllDataFailure(t *testing.T) {
	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 500,
				Body:       nil,
			}, errors.New("Mock Error")
		},
	}
	err := ListAllData("sampleURL")
	if err == nil {
		t.Error("TestListAllDataFailure failed.")
		return
	}
}

func TestSpecificTicketInfoSuccess(t *testing.T) {
	// build our response JSON
	jsonResponse := `
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
	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}
	err := SpecificTicketInfo("sampleURL", 1)
	if err != nil {
		t.Error("TestSpecificTicketInfo failed.")
		return
	}
}

func TestSpecificTicketInfoFailure(t *testing.T) {
	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 500,
				Body:       nil,
			}, errors.New("Mock Error")
		},
	}
	err := SpecificTicketInfo("sampleURL", 1)
	if err == nil {
		t.Error("TestSpecificTicketInfoFailure failed.")
		return
	}
}
