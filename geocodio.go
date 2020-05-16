package geocodio

import (
	"errors"
)

const (
	// GeocodioAPIBaseURLv1 is the Geocod.io Base URL
	GeocodioAPIBaseURLv1 = "https://api.geocod.io/v1.5"
)

// NewGeocodio is a helper to create new Geocodio pointer
func NewGeocodio(apiKey string) (*Geocodio, error) {

	if apiKey == "" {
		return nil, errors.New("apiKey is missing")
	}

	newGeocodio := new(Geocodio)
	newGeocodio.APIKey = apiKey

	return newGeocodio, nil
}
