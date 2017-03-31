package geocodio_test

import (
	"github.com/stevepartridge/geocodio"
	"testing"
)

func TestReverseGeocodeFullAddress(t *testing.T) {
	Geocodio, err := geocodio.NewGeocodio(APIKey())

	if err != nil {
		t.Error("Failed with API KEY set.", APIKey(), err)
	}

	result, err := Geocodio.ReverseGeocode(AddressTestTwoLatitude, AddressTestTwoLongitude)
	if err != nil {
		t.Error(err)
	}

	// t.Log(result.ResponseAsString())

	if len(result.Results) == 0 {
		t.Error("Results length is 0")
	}

	if len(result.Results) < 3 {
		t.Error("Results found length is less than 3", len(result.Results))
	}

	if len(result.Results) == 0 {
		t.Error("No results were found.")
		return
	}

	if result.Results[0].Formatted != "101 State Hwy 58, Nashville, NC 27856" {
		t.Error("Location latitude does not match", result.Results[0].Formatted, "101 State Hwy 58, Nashville, NC 27856")
	}

}
