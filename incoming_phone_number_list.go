package twiliolo

import (
	"encoding/json"
	"reflect"

	"github.com/genesor/twiliolo/option"
)

// IncomingPhoneNumberList represents the response of the Twilio API when calling /IncomingPhoneNumbers.json
type IncomingPhoneNumberList struct {
	Page                 int                    `json:"page"`
	PageSize             int                    `json:"page_size"`
	URI                  string                 `json:"uri"`
	FirstPageURI         string                 `json:"first_page_uri"`
	NextPageURI          string                 `json:"next_page_uri"`
	PreviousPageURI      string                 `json:"previous_page_uri"`
	IncomingPhoneNumbers []*IncomingPhoneNumber `json:"incoming_phone_numbers"`
}

// List retrieves the first page of all the Incoming Phone Number owned
// Doc: https://www.twilio.com/docs/api/rest/incoming-phone-numbers#list-get
func (s *IncomingPhoneNumberService) List(requestOptions ...option.RequestOption) (*IncomingPhoneNumberList, error) {
	body, err := s.Client.Get("/IncomingPhoneNumbers.json", requestOptions)
	if err != nil {
		return nil, err
	}

	var incomingPhoneNumberList *IncomingPhoneNumberList
	incomingPhoneNumberList = new(IncomingPhoneNumberList)
	err = json.Unmarshal(body, incomingPhoneNumberList)

	return incomingPhoneNumberList, err
}

// ListNextPage retrieves the next page of a given IncomingPhoneNumberList
// If an empty NextPageURI is present in the struct it'll return an error
// Doc: https://www.twilio.com/docs/api/rest/incoming-phone-numbers#list-get
func (s *IncomingPhoneNumberService) ListNextPage(previousList *IncomingPhoneNumberList, requestOptions ...option.RequestOption) (*IncomingPhoneNumberList, error) {
	if previousList == nil || previousList.NextPageURI == "" {
		return nil, ErrIncomingPhoneListNoNextPage
	}

	newRequestOptions := make([]option.RequestOption, 2)
	newRequestOptions[0] = option.Page(previousList.Page + 1)
	newRequestOptions[1] = option.PageSize(previousList.PageSize)

	for _, option := range requestOptions {
		// Override Page and PageSize options
		if reflect.TypeOf(option).Name() != "Page" || reflect.TypeOf(option).Name() != "PageSize" {
			newRequestOptions = append(newRequestOptions, option)
		}
	}

	body, err := s.Client.Get("/IncomingPhoneNumbers.json", newRequestOptions)
	if err != nil {
		return nil, err
	}

	var incomingPhoneNumberList *IncomingPhoneNumberList
	incomingPhoneNumberList = new(IncomingPhoneNumberList)
	err = json.Unmarshal(body, incomingPhoneNumberList)

	return incomingPhoneNumberList, err
}
