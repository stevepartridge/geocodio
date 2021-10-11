package geocodio

import (
	"errors"
	"fmt"
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

// BatchResultItem
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
	resp := GeocodeResult{}
	if address == "" {
		return resp, ErrAddressIsEmpty
	}

	return g.geocode(map[string]string{"q": address})
}

// Geocode single address with full component list
// See: https://www.geocod.io/docs/#single-address
func (g *Geocodio) GeocodeComponents(address InputAddress) (GeocodeResult, error) {
	if address.Street == "" && address.City == "" && address.State == "" && address.PostalCode == "" && address.Country == "" {
		return GeocodeResult{}, ErrAddressIsEmpty
	}

	return g.geocode(map[string]string{
		"street":      address.Street,
		"city":        address.City,
		"state":       address.State,
		"postal_code": address.PostalCode,
		"country":     address.Country})
}

func (g *Geocodio) geocode(params map[string]string) (GeocodeResult, error) {
	resp := GeocodeResult{}
	err := g.do("GET", "/geocode", params, nil, &resp)
	if err != nil {
		return resp, err
	}

	if len(resp.Results) == 0 {
		return resp, ErrNoResultsFound
	}

	return resp, nil
}

// GeocodeBatch lookup list of addresses (either string or InputAddress)
func (g *Geocodio) GeocodeBatch(addresses ...interface{}) (BatchResponse, error) {
	resp := BatchResponse{}
	if len(addresses) == 0 {
		return resp, ErrBatchAddressesIsEmpty
	}
	if err := verifyValidAddresses(addresses); err != nil {
		return resp, err
	}

	// TODO: support limit
	err := g.do("POST", "/geocode", nil, addresses, &resp)
	if err != nil {
		return resp, err
	}

	if len(resp.Results) == 0 {
		return resp, ErrNoResultsFound
	}

	return resp, nil
}

func verifyValidAddresses(addresses []interface{}) error {
	var builder strings.Builder
	for i := range addresses {
		switch addresses[i].(type) {
		case string, InputAddress:
		default:
			builder.WriteString(fmt.Sprintf("address[%d]: %t ", i, addresses[i]))
		}
	}
	if builder.Len() > 0 {
		return fmt.Errorf("all addresses must be of type string or InputAddress: %s", builder.String())
	}
	return nil
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

	err := g.do("GET", "/geocode", map[string]string{
		"q":      address,
		"fields": fieldsCommaSeparated,
	}, nil, &resp)
	if err != nil {
		return resp, err
	}

	if len(resp.Results) == 0 {
		return resp, ErrNoResultsFound
	}

	return resp, nil
}
