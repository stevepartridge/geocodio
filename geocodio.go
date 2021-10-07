package geocodio

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dghubble/sling"
)

const (
	// GeocodioAPIBaseURLv1 is the Geocod.io Base URL
	GeocodioAPIBaseURLv1 = "https://api.geocod.io/v1.6"
	EnvGeocodioAPIKey    = "GEOCODIO_API_KEY"
)

// Geocodio is the base struct
type Geocodio struct {
	APIKey string `url:"api_key"`
	client sling.Doer
}

type Input struct {
	AddressComponents Components `json:"address_components"`
	FormattedAddress  string     `json:"formatted_address"`
}

type Address struct {
	Query        string     `json:"query"`
	Components   Components `json:"address_components"`
	Formatted    string     `json:"formatted_address"`
	Location     Location   `json:"location"`
	Accuracy     float64    `json:"accuracy"`
	AccuracyType string     `json:"accuracy_type"`
	Source       string     `json:"source"`
	Fields       Fields     `json:"fields,omitempty"`
}

// Components
type Components struct {
	Number          string `json:"number"`
	Street          string `json:"street"`
	Suffix          string `json:"suffix"`
	SecondaryNumber string `json:"secondarynumber"`
	SecondaryUnit   string `json:"secondaryunit"`
	PostDirectional string `json:"postdirectional"`
	FormattedStreet string `json:"formatted_street"`
	City            string `json:"city"`
	State           string `json:"state"`
	Zip             string `json:"zip"`
	County          string `json:"county"`
	Country         string `json:"country"`
	PreDirectional  string `json:"predirectional"`
	Prefix          string `json:"prefix"`
}

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

// Fields
type Fields struct {
	Timezone                  Timezone                  `json:"timezone,omitempty"`
	Zip4                      Zip4                      `json:"zip4,omitempty"`
	CongressionalDistrict     CongressionalDistrict     `json:"congressional_district,omitempty"`  // v1.0
	CongressionalDistricts    []CongressionalDistrict   `json:"congressional_districts,omitempty"` // v1.1+
	StateLegislativeDistricts StateLegislativeDistricts `json:"state_legislative_districts,omitempty"`
	SchoolDistricts           SchoolDistricts           `json:"school_districts,omitempty"`
	Census                    CensusResults             `json:"census,omitempty"`
	ACS                       CensusACS                 `json:"acs,omitempty"`
}

// New creates a Geocodio instance based on an API key in either the environment
// or passed in as the first string value
func New(apiKey ...string) (*Geocodio, error) {
	key := os.Getenv(EnvGeocodioAPIKey)
	if len(apiKey) == 0 && strings.TrimSpace(key) == "" {
		return nil, ErrMissingAPIKey
	}

	if len(apiKey) == 1 {
		key = apiKey[0]
	}

	if strings.TrimSpace(key) == "" {
		return nil, ErrMissingAPIKey
	}

	timeout := time.Duration(10 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	g := Geocodio{
		APIKey: key,
		client: client,
	}

	return &g, nil
}

func (g *Geocodio) do(method, path string, params map[string]string, bodyJSON, result interface{}) error {
	s := sling.New().
		Doer(g.client).
		Base(GeocodioAPIBaseURLv1).
		QueryStruct(g).
		Set("Content-Type", "application/json").
		Path(path).
		BodyJSON(bodyJSON)

	req, err := s.Request()
	if err != nil {
		return err
	}
	req.Method = method
	req.URL.RawQuery = getQueryString(req, params)

	_, err = s.Do(req, result, nil)
	return err
}

func getQueryString(req *http.Request, params map[string]string) string {
	query := req.URL.Query()
	for key, value := range params {
		query.Add(key, value)
	}
	return query.Encode()
}
