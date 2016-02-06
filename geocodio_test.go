package geocodio_test

import (
	"github.com/stevepartridge/geocodio"
	"os"
	"testing"
)

const (
	API_KEY = "YOUR_API_KEY"

	TEST_ADDRESS_1_FULL        = "42370 Bob Hope Dr, Rancho Mirage, CA 92270"
	TEST_ADDRESS_1_WITHOUT_ZIP = "42370 Bob Hope Dr, Rancho Mirage, CA"
	TEST_ADDRESS_1_NUMBER      = "42370"
	TEST_ADDRESS_1_STREET      = "Bob Hope Dr"
	TEST_ADDRESS_1_CITY        = "Rancho Mirage"
	TEST_ADDRESS_1_STATE       = "CA"
	TEST_ADDRESS_2_FULL        = "500 H St NE, Washington, DC 20002"
	TEST_ADDRESS_2_WITHOUT_ZIP = "500 H St NE, Washington, DC"
	TEST_ADDRESS_2_NUMBER      = "500"
	TEST_ADDRESS_2_STREET      = "H St NE"
	TEST_ADDRESS_2_CITY        = "Washington"
	TEST_ADDRESS_2_STATE       = "DC"
	TEST_ADDRESS_2_LATITUDE    = 38.9002898
	TEST_ADDRESS_2_LONGITUDE   = -76.9990361
)

func ApiKey() string {
	if API_KEY != "YOUR_API_KEY" {
		return API_KEY
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
