package geocodio_test

import (
	"github.com/stevepartridge/geocodio"
	"os"
	"testing"
)

const (
	DefaultAPIKey = "YOUR_API_KEY"

	AddressTestOneFull       = "42370 Bob Hope Dr, Rancho Mirage, CA 92270"
	AddressTestOneWithoutZip = "42370 Bob Hope Dr, Rancho Mirage, CA"
	AddressTestOneNumber     = "42370"
	AddressTestOneStreet     = "Bob Hope Dr"
	AddressTestOneCity       = "Rancho Mirage"
	AddressTestOneState      = "CA"

	AddressTestTwoFull       = "500 H St NE, Washington, DC 20002"
	AddressTestTwoWithoutZip = "500 H St NE, Washington, DC"
	AddressTestTwoNumber     = "500"
	AddressTestTwoStreet     = "H St NE"
	AddressTestTwoCity       = "Washington"
	AddressTestTwoState      = "DC"
	AddressTestTwoLatitude   = 35.9746000
	AddressTestTwoLongitude  = -77.9658000
)

func APIKey() string {
	if DefaultAPIKey != "YOUR_API_KEY" {
		return DefaultAPIKey
	}
	return os.Getenv("API_KEY")
}

func TestGeocodioWithApiKey(t *testing.T) {
	_, err := geocodio.NewGeocodio(APIKey())
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
