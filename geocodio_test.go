package geocodio_test

import (
	"os"
	"testing"

	"github.com/stevepartridge/geocodio"
)

const (
	DefaultAPIKey = "YOUR_API_KEY"

	// AddressTestOneFull       = "42370 Bob Hope Dr, Rancho Mirage, CA 92270"
	// AddressTestOneWithoutZip = "42370 Bob Hope Dr, Rancho Mirage, CA"
	// AddressTestOneNumber     = "42370"
	// AddressTestOneStreet     = "Bob Hope Dr"
	// AddressTestOneCity       = "Rancho Mirage"
	// AddressTestOneState      = "CA"
	// AddressTestOneLatitude   = 33.738645
	// AddressTestOneLongitude  = -116.407141

	AddressTestOneFull       = "1109 N Highland St, Arlington, VA 22201"
	AddressTestOneWithoutZip = "1109 N Highland St, Arlington, VA"
	AddressTestOneNumber     = "1109"
	AddressTestOneStreet     = "N Highland St"
	AddressTestOneCity       = "Arlington"
	AddressTestOneState      = "VA"
	AddressTestOneLatitude   = 38.886672
	AddressTestOneLongitude  = -77.094735

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

// 1109+N+Highland+St%2c+Arlington+VA
