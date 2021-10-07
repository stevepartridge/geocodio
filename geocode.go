package geocodio

import (
	"errors"
	"strings"
)

// BatchResponse
type BatchResponse struct {
	Results []BatchResult `json:"results"`
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
}

type ErrorResponse struct {
	Message string    `json:"error"`
	Results []Address `json:"results"`
}

type Result struct {
	Address
	Error *ErrorResponse `json:"response,omitempty"`
}

// Geocode single address
// See: http://geocod.io/docs/#toc_4
func (g *Geocodio) Geocode(address string) (GeocodeResult, error) {
	res := GeocodeResult{}
	if address == "" {
		return res, ErrAddressIsEmpty
	}

	err := g.do("GET", "/geocode", map[string]string{"q": address}, nil, &res)
	if err != nil {
		return res, err
	}

	if len(res.Results) == 0 {
		return res, ErrNoResultsFound
	}

	return res, nil
}

// GeocodeBatch look up addresses
func (g *Geocodio) GeocodeBatch(addresses ...string) (BatchResponse, error) {
	res := BatchResponse{}
	if len(addresses) == 0 {
		return res, ErrBatchAddressesIsEmpty
	}

	// TODO: support limit
	err := g.do("POST", "/geocode", nil, addresses, &res)
	if err != nil {
		return res, err
	}

	if len(res.Results) == 0 {
		return res, ErrNoResultsFound
	}

	return res, nil
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
	res := GeocodeResult{}
	if address == "" {
		return res, errors.New("address can not be empty")
	}

	fieldsCommaSeparated := strings.Join(fields, ",")

	err := g.do("GET", "/geocode", map[string]string{
		"q":      address,
		"fields": fieldsCommaSeparated,
	}, nil, &res)
	if err != nil {
		return res, err
	}

	if len(res.Results) == 0 {
		return res, ErrNoResultsFound
	}

	return res, nil
}
