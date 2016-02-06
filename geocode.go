package geocodio

import (
	"errors"
	"strings"
)

// Geocode single address
// See: http://geocod.io/docs/#toc_4
func (g *Geocodio) Geocode(address string) (GeocodeResult, error) {
	if address == "" {
		return GeocodeResult{}, errors.New("address must not be empty")
	}

	results, err := g.Call("/geocode", map[string]string{"q": address})
	if err != nil {
		return GeocodeResult{}, err
	}

	return results, nil
}

// GeocodeAndReturnTimezone will geocode and include Timezone in the fields response
func (g *Geocodio) GeocodeAndReturnTimezone(address string) (GeocodeResult, error) {
	return g.GeocodeReturnFields(address, "timezone")
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

	results, err := g.Call("/geocode",
		map[string]string{
			"q":      address,
			"fields": fieldsCommaSeparated,
		})
	if err != nil {
		return GeocodeResult{}, err
	}

	return results, nil
}
