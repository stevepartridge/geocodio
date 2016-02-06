package geocodio_test

import (
	"github.com/stevepartridge/geocodio"
	"testing"
)

func TestReverseGeocodeFullAddress(t *testing.T) {
	Geocodio, err := geocodio.NewGeocodio(API_KEY)

	if err != nil {
		t.Error("Failed with API KEY set.", API_KEY, err)
	}

	result, err := Geocodio.ReverseGeocode(TEST_ADDRESS_2_LATITUDE, TEST_ADDRESS_2_LONGITUDE)
	if err != nil {
		t.Error(err)
	}

	t.Log(result.Debug)

	if len(result.Results) == 0 {
		t.Error("Results length is 0")
	}

	if len(result.Results) < 3 {
		t.Error("Results found length is less than 3", len(result.Results))
	}

	if result.Results[0].Location.Latitude != 38.900203 {
		t.Error("Location latitude does not match", result.Results[0].Location.Latitude, 38.900203)
	}

	// if result.Results[0].Location.Longitude != ​-76.999507 {
	//    t.Error("Location longitude does not match", result.Results[0].Location.Longitude, ​"-76.999507")
	//  }
}
