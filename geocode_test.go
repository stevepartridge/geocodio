package geocodio_test

import (
	"testing"

	"github.com/stevepartridge/geocodio"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGeocodeWithEmptyAddress(t *testing.T) {
	gc, err := geocodio.New()
	require.NoError(t, err)

	_, err = gc.Geocode("")
	assert.Error(t, err)
}

func TestGeocodeFullAddress(t *testing.T) {
	gc, err := geocodio.New()
	require.NoError(t, err)

	result, err := gc.Geocode(AddressTestOneFull)
	require.NoError(t, err)
	require.True(t, len(result.Results) > 0)
	assert.Equal(t, AddressTestOneLatitude, result.Results[0].Location.Latitude)
	assert.Equal(t, AddressTestOneLongitude, result.Results[0].Location.Longitude)
}

func TestGeocodeFullAddressReturningTimezone(t *testing.T) {
	gc, err := geocodio.New()
	require.NoError(t, err)

	result, err := gc.GeocodeAndReturnTimezone(AddressTestOneFull)
	assert.NoError(t, err)
	require.True(t, len(result.Results) > 0)

	assert.Equal(t, AddressTestOneLatitude, result.Results[0].Location.Latitude)
	assert.Equal(t, AddressTestOneLongitude, result.Results[0].Location.Longitude)
	assert.NotEmpty(t, result.Results[0].Fields.Timezone.Name)
	assert.True(t, result.Results[0].Fields.Timezone.ObservesDST)
}

func TestGeocodeFullAddressReturningZip4(t *testing.T) {
	gc, err := geocodio.New()
	require.NoError(t, err)

	result, err := gc.GeocodeAndReturnZip4(AddressTestOneFull)
	require.NoError(t, err)
	require.True(t, len(result.Results) > 0)

	assert.Equal(t, AddressTestOneLatitude, result.Results[0].Location.Latitude)
	assert.Equal(t, AddressTestOneLongitude, result.Results[0].Location.Longitude)
	assert.True(t, len(result.Results[0].Fields.Zip4.Plus4) > 0)
}

func TestGeocodeFullAddressReturningCongressionalDistrict(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
		t.Fail()
	}
	result, err := gc.GeocodeAndReturnCongressionalDistrict(AddressTestOneFull)
	if err != nil {
		t.Error(err)
	}

	if len(result.Results) == 0 {
		t.Error("Results length is 0")
		t.Fail()
	}

	if result.Results[0].Location.Latitude != AddressTestOneLatitude {
		t.Error("Location latitude does not match", result.Results[0].Location.Latitude, AddressTestOneLatitude)
		t.Fail()
	}

	if result.Results[0].Location.Longitude != AddressTestOneLongitude {
		t.Error("Location longitude does not match", result.Results[0].Location.Longitude, AddressTestOneLongitude)
		t.Fail()
	}

	if len(result.Results[0].Fields.CongressionalDistricts) == 0 {
		t.Error("Congressional District field not found", result.Results[0].Fields.CongressionalDistrict)
		t.Fail()
	}

	if result.Results[0].Fields.CongressionalDistricts[0].Name == "" {
		t.Error("Congressional District field not found", result.Results[0].Fields.CongressionalDistricts[0])
		t.Fail()
	}

	if result.Results[0].Fields.CongressionalDistricts[0].DistrictNumber != 8 {
		t.Error("Congressional District field does not match", result.Results[0].Fields.CongressionalDistrict)
		t.Fail()
	}
}

func TestGeocodeFullAddressReturningStateLegislativeDistricts(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
		t.Fail()
	}

	result, err := gc.GeocodeAndReturnStateLegislativeDistricts(AddressTestOneFull)
	if err != nil {
		t.Error(err)
	}

	// t.Log(result.ResponseAsString())

	if len(result.Results) == 0 {
		t.Error("Results length is 0", result)
		t.Fail()
	}

	if result.Results[0].Location.Latitude != AddressTestOneLatitude {
		t.Errorf("Location latitude %f does not match %f", result.Results[0].Location.Latitude, AddressTestOneLatitude)
		t.Fail()
	}

	if result.Results[0].Location.Longitude != AddressTestOneLongitude {
		t.Errorf("Location longitude %f does not match %f", result.Results[0].Location.Longitude, AddressTestOneLongitude)
		t.Fail()
	}

	if result.Results[0].Fields.StateLegislativeDistricts.House.DistrictNumber != "47" {
		t.Error("State Legislative Districts house does not match", result.Results[0].Fields.StateLegislativeDistricts.House)
		t.Fail()
	}

	if result.Results[0].Fields.StateLegislativeDistricts.Senate.DistrictNumber != "31" {
		t.Error("State Legislative Districts senate does not match", result.Results[0].Fields.StateLegislativeDistricts.Senate)
		t.Fail()
	}
}

func TestGeocodeFullAddressReturningMultipleFields(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
	result, err := gc.GeocodeReturnFields(AddressTestOneFull, "timezone", "cd")
	if err != nil {
		t.Error(err)
	}

	// fmt.Println(result.Debugc.RequestedURL)

	if len(result.Results) == 0 {
		t.Error("Results length is 0")
	}

	if result.Results[0].Location.Latitude != AddressTestOneLatitude {
		t.Error("Location latitude does not match", result.Results[0].Location.Latitude, AddressTestOneLatitude)
	}

	if result.Results[0].Location.Longitude != AddressTestOneLongitude {
		t.Error("Location longitude does not match", result.Results[0].Location.Longitude, AddressTestOneLongitude)
	}

	if result.Results[0].Fields.Timezone.Name == "" {
		t.Error("Timezone field not found")
	}

	if !result.Results[0].Fields.Timezone.ObservesDST {
		t.Error("Timezone field does not match", result.Results[0].Fields.Timezone)
	}

	congressionalDistrict := geocodio.CongressionalDistrict{}

	// check congressional district
	if result.Results[0].Fields.CongressionalDistrict.Name != "" {
		congressionalDistrict = result.Results[0].Fields.CongressionalDistrict
	} else if len(result.Results[0].Fields.CongressionalDistricts) > 0 {
		congressionalDistrict = result.Results[0].Fields.CongressionalDistricts[0]
	}

	if congressionalDistrict.Name == "" {
		t.Error("Congressional District field not found", congressionalDistrict)
	}

	if congressionalDistrict.DistrictNumber != 8 {
		t.Error("Congressional District field does not match", result.Results[0].Fields.CongressionalDistrict)
	}

}

func TestGeocodeBatchGeocode(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
	resp, err := gc.GeocodeBatch(
		AddressTestOneFull,
		AddressTestTwoFull,
		AddressTestThreeFull,
	)

	if err != nil {
		t.Error(err.Error())
	}

	if len(resp.Results) == 0 {
		t.Error("Results length is 0")
	}

	if len(resp.Results) < 3 {
		t.Error("Expected 3 or more results but saw", len(resp.Results))
	}
}

func TestGeocodeBatchEmptyListGeocode(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}

	_, err = gc.GeocodeBatch()
	if err == nil {
		t.Error("Expected to see an error")
	}

}

func TestGeocodeInvalidNoResults(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}

	_, err = gc.Geocode("123 Nonsense Ln, Nowhere, XX")
	if err == nil {
		t.Error("Expected to see an error")
		return
	}
	if err != geocodio.ErrNoResultsFound {
		t.Error("Expected error", geocodio.ErrNoResultsFound.Error(), "but saw", err.Error())
	}
}

func TestGeocodeBatchInvalidNoResults(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}

	resp, err := gc.GeocodeBatch("123 Nonsense Ln, Nowhere, XX")
	if err != nil {
		t.Error("Expected success", err)
	}
	if resp.Results[0].Response.Error == "" {
		t.Error("Expected to see an error")
	}
}

// TODO: School District (school)
