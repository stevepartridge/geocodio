package geocodio_test

import (
	"github.com/stevepartridge/geocodio"
	"testing"
)

func TestReverseGeocodeFullAddress(t *testing.T) {
	Geocodio, err := geocodio.NewGeocodio(ApiKey())

	if err != nil {
		t.Error("Failed with API KEY set.", ApiKey(), err)
	}

	result, err := Geocodio.ReverseGeocode(AddressTestTwoLatitude, AddressTestTwoLongitude)
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
