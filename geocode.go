package geocodio

import (
	"errors"
	"strings"
)

// BatchResponse
type BatchResponse struct {
	Results []BatchResult `json:"results"`
	Debug   struct {
		RawResponse  []byte `json:"-"`
		RequestedURL string `json:"requested_url"`
		Status       string `json:"status"`
		StatusCode   int    `json:"status_code"`
	} `json:"-"`
}

// BatchResult
type BatchResult struct {
	Query    string          `json:"query"`
	Response BatchResultItem `json:"response"`
}

type BatchResultItem struct {
	Input   Input     `json:"input,omitempty"`
	Results []Address `json:"results"`
	Error   string    `json:"error,omitempty"`
}

// GeocodeResponse
type GeocodeResult struct {
	Input   Input    `json:"input,omitempty"`
	Results []Result `json:"results"`
	Debug   struct {
		RawResponse  []byte `json:"-"`
		RequestedURL string `json:"requested_url"`
		Status       string `json:"status"`
		StatusCode   int    `json:"status_code"`
	} `json:"-"`
}

type ErrorResponse struct {
	Message string    `json:"error"`
	Results []Address `json:"results"`
}

type Result struct {
	Address
	Error *ErrorResponse `json:"response,omitempty"`
}

func (self *GeocodeResult) SaveDebug(requestedURL, status string, statusCode int, body []byte) {
	self.Debug.RequestedURL = requestedURL
	self.Debug.Status = status
	self.Debug.StatusCode = statusCode
	self.Debug.RawResponse = body
}

func (self *GeocodeResult) Error() string {
	if len(self.Results) > 0 {
		if self.Results[0].Error != nil {
			return self.Results[0].Error.Message
		}
	}
	return ""
}

// ResponseAsString helper to return raw response
func (self *GeocodeResult) ResponseAsString() string {
	return string(self.Debug.RawResponse)
}

func (self *BatchResponse) SaveDebug(requestedURL, status string, statusCode int, body []byte) {
	self.Debug.RequestedURL = requestedURL
	self.Debug.Status = status
	self.Debug.StatusCode = statusCode
	self.Debug.RawResponse = body
}

// ResponseAsString helper to return raw response
func (self *BatchResponse) ResponseAsString() string {
	return string(self.Debug.RawResponse)
}

// Geocode single address
// See: http://geocod.io/docs/#toc_4
func (g *Geocodio) Geocode(address string) (GeocodeResult, error) {
	resp := GeocodeResult{}
	if address == "" {
		return resp, ErrAddressIsEmpty
	}

	err := g.get("/geocode", map[string]string{"q": address}, &resp)
	if err != nil {
		return GeocodeResult{}, err
	}

	if len(resp.Results) == 0 {
		return resp, ErrNoResultsFound
	}

	return resp, nil
}

// GeocodeBatch look up addresses
func (g *Geocodio) GeocodeBatch(addresses ...string) (BatchResponse, error) {
	resp := BatchResponse{}
	if len(addresses) == 0 {
		return resp, ErrBatchAddressesIsEmpty
	}

	// TODO: support limit
	err := g.post("/geocode", addresses, nil, &resp)
	if err != nil {
		return BatchResponse{}, err
	}

	if len(resp.Results) == 0 {
		return resp, ErrNoResultsFound
	}

	return resp, nil
}

// GeocodeAndReturnTimezone will geocode and include Timezone in the fields response
func (g *Geocodio) GeocodeAndReturnTimezone(address string) (GeocodeResult, error) {
	return g.GeocodeReturnFields(address, "timezone")
}

// GeocodeAndReturnZip4 will geocode and include zip4 in the fields response
func (g *Geocodio) GeocodeAndReturnZip4(address string) (GeocodeResult, error) {
	return g.GeocodeReturnFields(address, "zip4")
}

// GeocodeAndReturnCongressionalDistrict will geocode and include Congressional District in the fields response
func (g *Geocodio) GeocodeAndReturnCongressionalDistrict(address string) (GeocodeResult, error) {
	return g.GeocodeReturnFields(address, "cd")
}

// GeocodeAndReturnStateLegislativeDistricts will geocode and include State Legislative Districts in the fields response
func (g *Geocodio) GeocodeAndReturnStateLegislativeDistricts(address string) (GeocodeResult, error) {
	return g.GeocodeReturnFields(address, "stateleg")
}

// TODO: School District (school)

// GeocodeReturnFields will geocode and includes additional fields in response
/*
 	See: http://geocod.io/docs/#toc_22
	Note:
		Each field counts as an additional lookup each
*/
func (g *Geocodio) GeocodeReturnFields(address string, fields ...string) (GeocodeResult, error) {
	resp := GeocodeResult{}
	if address == "" {
		return resp, errors.New("address can not be empty")
	}

	fieldsCommaSeparated := strings.Join(fields, ",")

	err := g.get("/geocode",
		map[string]string{
			"q":      address,
			"fields": fieldsCommaSeparated,
		}, &resp)
	if err != nil {
		return resp, err
	}

	if len(resp.Results) == 0 {
		return resp, ErrNoResultsFound
	}

	return resp, nil
}
