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
