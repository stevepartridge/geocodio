package geocodio_test

import (
	"os"
	"testing"

	"github.com/stevepartridge/geocodio"
)

const (
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

	AddressTestTwoFull       = "100 Legends Way, Boston, MA 02114"
	AddressTestTwoWithoutZip = "100 Legends Way, Boston, MA"
	AddressTestTwoNumber     = "100"
	AddressTestTwoStreet     = "Legends Way"
	AddressTestTwoCity       = "Boston"
	AddressTestTwoState      = "MA"
	AddressTestTwoLatitude   = 42.36629
	AddressTestTwoLongitude  = -71.0622

	AddressTestThreeFull       = "19 Tony Gwynn Drive, San Diego, CA 02114"
	AddressTestThreeWithoutZip = "19 Tony Gwynn Drive, San Diego, CA"
	AddressTestThreeNumber     = "19"
	AddressTestThreeStreet     = "Tony Gwynn Drive"
	AddressTestThreeCity       = "San Diego"
	AddressTestThreeState      = "CA"
	AddressTestThreeLatitude   = 32.708343
	AddressTestThreeLongitude  = -117.158124
)

func TestGeocodioWithApiKey(t *testing.T) {

	_, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
}

func TestGeocodioWithoutApiKey(t *testing.T) {

	key := os.Getenv(geocodio.EnvGeocodioAPIKey)
	if key == "" && os.Getenv(geocodio.EnvOldAPIKey) != "" {
		key = os.Getenv(geocodio.EnvOldAPIKey)
	}

	os.Setenv(geocodio.EnvGeocodioAPIKey, "")
	os.Setenv(geocodio.EnvOldAPIKey, "")

	_, err := geocodio.New()
	if err == nil {
		t.Error("Did not throw error with no API Key set")
	}

	os.Setenv(geocodio.EnvGeocodioAPIKey, key)
}

func TestGeocodioWithoutApiKeyEnvAndEmptyString(t *testing.T) {

	key := os.Getenv(geocodio.EnvGeocodioAPIKey)
	if key == "" && os.Getenv(geocodio.EnvOldAPIKey) != "" {
		key = os.Getenv(geocodio.EnvOldAPIKey)
	}

	os.Setenv(geocodio.EnvGeocodioAPIKey, "")
	os.Setenv(geocodio.EnvOldAPIKey, "")

	_, err := geocodio.New("")
	if err == nil {
		t.Error("Did not throw error with no API Key set")
	}

	os.Setenv(geocodio.EnvGeocodioAPIKey, key)
}

func TestGeocodioDeprecatedWithApiKey(t *testing.T) {
	_, err := geocodio.NewGeocodio(os.Getenv(geocodio.EnvGeocodioAPIKey))
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
}

func TestGeocodioDeprecatedWithoutApiKey(t *testing.T) {
	_, err := geocodio.NewGeocodio("")
	if err == nil {
		t.Error("Expected error:", geocodio.ErrMissingAPIKey, " but did not see any error")
	}
}

func TestGeocodioWithOldApiKey(t *testing.T) {
	os.Setenv(geocodio.EnvOldAPIKey, os.Getenv(geocodio.EnvGeocodioAPIKey))
	os.Setenv(geocodio.EnvGeocodioAPIKey, "")

	_, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}

	os.Setenv(geocodio.EnvGeocodioAPIKey, os.Getenv(geocodio.EnvOldAPIKey))
	os.Setenv(geocodio.EnvOldAPIKey, "")
}
