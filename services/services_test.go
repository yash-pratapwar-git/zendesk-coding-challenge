package services

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/zendesk-coding-challenge/constants"
)

// Custom type that allows setting the func that our Mock Do func will run while running tests
type MockDoType func(req *http.Request) (*http.Response, error)

// MockClient is the mocked client for testing purpose
type MockClient struct {
	MockDo MockDoType
}

// Overriding the Do function of Client
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

//Testing the success case of GetHttpMethodSuccess function
func TestGetHttpMethodSuccess(t *testing.T) {
	// build server response
	jsonResponse := `[{
   "status": "success"
  }]`
	// create a new reader
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
		t.Error("TestGetHttpMethodSuccess failed.")
		return
	}
}

//Testing the failed request case of GetHttpMethod function
func TestGetHttpMethodUnsuccessfullReq(t *testing.T) {
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
		t.Error("TestGetHttpMethodUnsuccessfullReq failed.")
		return
	}
}

//Testing the case of receiving error from server in GetHttpMethos function
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

//Testing the success case of ListAllData functions
func TestListAllData(t *testing.T) {
	// build our response JSON
	jsonResponse := constants.SampleAllTicketsData
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
		t.Error("TestListAllData failed.")
		return
	}
}

//Testing the fail case of ListAllData function
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

//Testing the failure case of SpecificTicketInfo function
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

//Testing success case of SpecificTicketInfo function
func TestSpecificTicketInfoSuccess(t *testing.T) {
	// build our response JSON
	jsonResponse := constants.SampleSingleTicketResponse
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
		t.Error("TestSpecificTicketInfoSuccess failed.")
		return
	}
}
