package geocodio

import (
	"os"
	"strings"
)

const (
	// GeocodioAPIBaseURLv1 is the Geocod.io Base URL
	GeocodioAPIBaseURLv1 = "https://api.geocod.io/v1.6"
)

// Geocodio is the base struct
type Geocodio struct {
	APIKey string
}

type Input struct {
	AddressComponents Components `json:"address_components"`
	FormattedAddress  string     `json:"formatted_address"`
}

// New creates a new Geocodio instance based on an API key in either the environment
// or passed in as the first string value
func New(apiKey ...string) (*Geocodio, error) {

	key := os.Getenv(EnvGeocodioAPIKey)
	if strings.TrimSpace(key) == "" {
		key = os.Getenv(EnvOldAPIKey)
	}
	if len(apiKey) == 0 && strings.TrimSpace(key) == "" {
		return nil, ErrMissingApiKey
	}

	if len(apiKey) == 1 {
		key = apiKey[0]
	}

	if strings.TrimSpace(key) == "" {
		return nil, ErrMissingApiKey
	}

	return NewGeocodio(key)
}

// NewGeocodio is a helper to create new Geocodio reference
// since 1.6+ this is kept for backwards compatiblity
// after 2+ this will be deprecated
func NewGeocodio(apiKey string) (*Geocodio, error) {

	if apiKey == "" {
		return nil, ErrMissingApiKey
	}

	g := Geocodio{
		APIKey: apiKey,
	}

	return &g, nil
}
