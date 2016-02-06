package geocodio

import (
	"errors"
)

const (
	GEOCODIO_API_V1_BASE_URL = "https://api.geocod.io/v1"
)

// Helper to create new Geocodio pointer
func NewGeocodio(apiKey string) (*Geocodio, error) {

	if apiKey == "" {
		return nil, errors.New("apiKey is missing")
	}

	newGeocodio := new(Geocodio)
	newGeocodio.ApiKey = apiKey

	return newGeocodio, nil
}
