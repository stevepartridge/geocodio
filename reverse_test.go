package geocodio_test

import (
	"fmt"
	"testing"

	"github.com/stevepartridge/geocodio"
)

func TestReverseGeocodeLookup(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}

	result, err := gc.Reverse(AddressTestTwoLatitude, AddressTestTwoLongitude)
	if err != nil {
		t.Error(err)
	}

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

	if result.Results[0].Formatted != AddressTestTwoFull {
		t.Error("Location latitude does not match", result.Results[0].Formatted, AddressTestTwoFull)
	}
}

func TestDeprecatedReverseGeocodeLookup(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}

	result, err := gc.ReverseGeocode(AddressTestTwoLatitude, AddressTestTwoLongitude)
	if err != nil {
		t.Error(err)
	}

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

	if result.Results[0].Formatted != AddressTestTwoFull {
		t.Error("Location latitude does not match", result.Results[0].Formatted, AddressTestTwoFull)
	}
}

func TestReverseWithZeroLatLng(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
	_, err = gc.Reverse(0.0, 0.0)
	if err != geocodio.ErrReverseGecodeMissingLatLng {
		t.Errorf("Error should be '%s' not '%s'", geocodio.ErrReverseGecodeMissingLatLng.Error(), err.Error())
	}
}

func TestReverseBatchLookup(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}

	result, err := gc.ReverseBatch(
		AddressTestOneLatitude, AddressTestOneLongitude,
		AddressTestTwoLatitude, AddressTestTwoLongitude,
		AddressTestThreeLatitude, AddressTestThreeLongitude,
	)
	fmt.Println(result.ResponseAsString())
	if err != nil {
		t.Error(err)
	}

	if len(result.Results) == 0 {
		t.Error("Results length is 0")
	}

}

func TestReverseBatchWithoutLatLng(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
	_, err = gc.ReverseBatch()
	if err != geocodio.ErrReverseBatchMissingCoords {
		t.Errorf("Error should be '%s' not '%s'", geocodio.ErrReverseGecodeMissingLatLng.Error(), err.Error())
	}
}

func TestReverseBatchWithInvalidLatLngPairs(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}

	_, err = gc.ReverseBatch(AddressTestOneLatitude)
	if err != geocodio.ErrReverseBatchInvalidCoordsPairs {
		t.Errorf("Error should be '%s' not '%s'", geocodio.ErrReverseGecodeMissingLatLng.Error(), err.Error())
	}

	_, err = gc.ReverseBatch(AddressTestOneLatitude, AddressTestOneLongitude, AddressTestTwoLatitude)
	if err != geocodio.ErrReverseBatchInvalidCoordsPairs {
		t.Errorf("Error should be '%s' not '%s'", geocodio.ErrReverseGecodeMissingLatLng.Error(), err.Error())
	}
}

func TestReverseWithInvalidLatLng(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}

	_, err = gc.Reverse(1.234, 5.678)
	if err == nil {
		t.Error("Expected to see an error")
		return
	}
	if err != geocodio.ErrNoResultsFound {
		t.Error("Expected error", geocodio.ErrNoResultsFound.Error(), "but saw", err.Error())
	}

	_, err = gc.ReverseBatch(1.234, 5.678)
	if err == nil {
		t.Error("Expected to see an error")
		return
	}
	if err != geocodio.ErrNoResultsFound {
		t.Error("Expected error", geocodio.ErrNoResultsFound.Error(), "but saw", err.Error())
	}

}
