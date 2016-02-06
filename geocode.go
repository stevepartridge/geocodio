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

func (g *Geocodio) GeocodeAndReturnTimezone(address string) (GeocodeResult, error) {
	return g.GeocodeReturnFields(address, "timezone")
}

func (g *Geocodio) GeocodeAndReturnCongressionalDistrict(address string) (GeocodeResult, error) {
	return g.GeocodeReturnFields(address, "cd")
}

func (g *Geocodio) GeocodeAndReturnStateLegislativeDistricts(address string) (GeocodeResult, error) {
	return g.GeocodeReturnFields(address, "stateleg")
}

// TODO: School District (school)

// Include additional fields
// Note: each field counts as an additional lookup each
// See: http://geocod.io/docs/#toc_22
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
