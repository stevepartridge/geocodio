package geocodio_test

import (
	"github.com/stevepartridge/geocodio"
	"os"
	"testing"
)

const (
	DefaultApiKey = "YOUR_API_KEY"

	AddressTestOne_Full       = "42370 Bob Hope Dr, Rancho Mirage, CA 92270"
	AddressTestOne_WithoutZip = "42370 Bob Hope Dr, Rancho Mirage, CA"
	AddressTestOne_Number     = "42370"
	AddressTestOne_Street     = "Bob Hope Dr"
	AddressTestOne_City       = "Rancho Mirage"
	AddressTestOne_State      = "CA"

	AddressTestTwo_Full       = "500 H St NE, Washington, DC 20002"
	AddressTestTwo_WithoutZip = "500 H St NE, Washington, DC"
	AddressTestTwo_Number     = "500"
	AddressTestTwo_Street     = "H St NE"
	AddressTestTwo_City       = "Washington"
	AddressTestTwo_State      = "DC"
	AddressTestTwo_Latitude   = 38.9002898
	AddressTestTwo_Longitude  = -76.9990361
)

func ApiKey() string {
	if DefaultApiKey != "YOUR_API_KEY" {
		return DefaultApiKey
	}
	return os.Getenv("API_KEY")
}

func TestGeocodioWithApiKey(t *testing.T) {
	_, err := geocodio.NewGeocodio(ApiKey())
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
}

func TestGeocodioWithoutApiKey(t *testing.T) {
	_, err := geocodio.NewGeocodio("")
	if err == nil {
		t.Error("Did not through error when omitting API KEY")
	}
}
