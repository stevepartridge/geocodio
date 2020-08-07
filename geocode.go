package geocodio

import (
	"errors"
	"strings"
)

// GeocodeResponse
type GeocodeResult struct {
	Input   Input     `json:"input,omitempty"`
	Results []Address `json:"results"`
	Debug   struct {
		RawResponse  []byte `json:"-"`
		RequestedURL string `json:"requested_url"`
		Status       string `json:"status"`
		StatusCode   int    `json:"status_code"`
	} `json:"-"`
}

func (self *GeocodeResult) ResponseAsString() string {
	return string(self.Debug.RawResponse)
}

// Geocode single address
// See: http://geocod.io/docs/#toc_4
func (g *Geocodio) Geocode(address string) (GeocodeResult, error) {
	if address == "" {
		return GeocodeResult{}, ErrAddressIsEmpty
	}

	results, err := g.get("/geocode", map[string]string{"q": address})
	if err != nil {
		return GeocodeResult{}, err
	}

	return results, nil
}

// Geocode batch addresses
func (g *Geocodio) GeocodeBatch(addresses ...string) (GeocodeResult, error) {
	if len(addresses) == 0 {
		return GeocodeResult{}, ErrBatchAddressesIsEmpty
	}

	// TODO: support limit
	results, err := g.post("/geocode", addresses, nil)
	if err != nil {
		return GeocodeResult{}, err
	}

	return results, nil
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
	if address == "" {
		return GeocodeResult{}, errors.New("address can not be empty")
	}

	fieldsCommaSeparated := strings.Join(fields, ",")

	result, err := g.Call("/geocode",
		map[string]string{
			"q":      address,
			"fields": fieldsCommaSeparated,
		})
	if err != nil {
		return GeocodeResult{}, err
	}

	return result, nil
}
