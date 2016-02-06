package geocodio_test

import (
	"github.com/stevepartridge/geocodio"
	"testing"
)

const (
	API_KEY = "c00b7df5dcb85daf6701f552b5f78c8225826b6"

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

func TestGeocodioWithApiKey(t *testing.T) {
	_, err := geocodio.NewGeocodio(API_KEY)
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
