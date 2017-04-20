package twiliolo

import (
	"encoding/json"
	"net/url"
)

// AvailablePhoneNumberServiceInterface is the interface of a IncomingPhoneNumberService
type AvailablePhoneNumberServiceInterface interface {
	Local(string, ...RequestOption) ([]AvailablePhoneNumber, error)
}

// AvailablePhoneNumberService handles communication with the Incoming Phone Number related methods.
type AvailablePhoneNumberService service

// AvailablePhoneNumber represents a Twilio Incoming Phone Number.
type AvailablePhoneNumber struct {
	FriendlyName        string      `json:"friendly_name"`
	PhoneNumber         string      `json:"phone_number"`
	ISOCountry          string      `json:"iso_country"`
	AddressRequirements string      `json:"address_requirements"`
	Capabilities        Capabilites `json:"capabilities"`
	Beta                bool        `json:"beta"`
	// US & CA Only
	Lata       string `json:"lata"`
	RateCenter string `json:"rate_center"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	Region     string `json:"region"`
	PostalCode string `json:"postal_code"`
}

type SearchAvailablePhoneNumber struct {
	AvailablePhoneNumbers []AvailablePhoneNumber `json:"available_phone_numbers"`
	URI                   string                 `json:"uri"`
}

// Local performs a call to the twilio API to retrieve Incoming Phone Numbers
// available with the given params
func (s *AvailablePhoneNumberService) Local(countryCode string, requestOptions ...RequestOption) (*IncomingPhoneNumber, error) {
	var incomingPhoneNumber *IncomingPhoneNumber

	res, err := s.Client.Get("/AvailablePhoneNumber/"+countryCode+"/Local.json", requestOptions)
	if err != nil {
		return nil, err
	}

	incomingPhoneNumber = new(IncomingPhoneNumber)
	err = json.Unmarshal(res, incomingPhoneNumber)

	return incomingPhoneNumber, err
}

// Buy performs the update of the differents attributes of an Incoming Phone Number.
func (s *AvailablePhoneNumberService) Buy(availablePhoneNumber *AvailablePhoneNumber, requestOptions ...RequestOption) (*IncomingPhoneNumber, error) {

	updates := url.Values{}
	updates.Set("PhoneNumber", availablePhoneNumber.PhoneNumber)
	updates.Set("FriendlyName", availablePhoneNumber.FriendlyName)

	body, err := s.Client.Post("/IncomingPhoneNumbers.json", requestOptions, updates)
	if err != nil {
		return nil, err
	}

	var incomingPhoneNumber *IncomingPhoneNumber

	err = json.Unmarshal(body, incomingPhoneNumber)
	if err != nil {
		return nil, err
	}

	return incomingPhoneNumber, nil
}