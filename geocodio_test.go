package geocodio_test

import (
	"os"
	"testing"

	"github.com/mypricehealth/jsonassert"
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
	os.Setenv(geocodio.EnvGeocodioAPIKey, "")

	_, err := geocodio.New()
	if err == nil {
		t.Error("Did not throw error with no API Key set")
	}

	os.Setenv(geocodio.EnvGeocodioAPIKey, key)
}

func TestGeocodioWithoutApiKeyEnvAndEmptyString(t *testing.T) {
	key := os.Getenv(geocodio.EnvGeocodioAPIKey)
	os.Setenv(geocodio.EnvGeocodioAPIKey, "")

	_, err := geocodio.New("")
	if err == nil {
		t.Error("Did not throw error with no API Key set")
	}

	os.Setenv(geocodio.EnvGeocodioAPIKey, key)
}

func TestVerifyStructs(t *testing.T) {
	jsonassert.StructCheck(t, "testdata/batchResponse.json", &geocodio.BatchResponse{})
	jsonassert.StructCheck(t, "testdata/fields-acs-demographics.json", &geocodio.Fields{})
	jsonassert.StructCheck(t, "testdata/fields-acs-economics.json", &geocodio.Fields{})
	jsonassert.StructCheck(t, "testdata/fields-acs-families.json", &geocodio.Fields{})
	jsonassert.StructCheck(t, "testdata/fields-acs-housing.json", &geocodio.Fields{})
	jsonassert.StructCheck(t, "testdata/fields-acs-social.json", &geocodio.Fields{})
	jsonassert.StructCheck(t, "testdata/fields-cd.json", &geocodio.Fields{})
	jsonassert.StructCheck(t, "testdata/fields-census.json", &geocodio.Fields{})
	jsonassert.StructCheck(t, "testdata/fields-riding.json", &geocodio.Fields{})
	jsonassert.StructCheck(t, "testdata/fields-school (unified).json", &geocodio.Fields{})
	jsonassert.StructCheck(t, "testdata/fields-school.json", &geocodio.Fields{})
	jsonassert.StructCheck(t, "testdata/fields-statcan.json", &geocodio.Fields{})
	jsonassert.StructCheck(t, "testdata/fields-stateleg.json", &geocodio.Fields{})
	jsonassert.StructCheck(t, "testdata/fields-timezone.json", &geocodio.Fields{})
	jsonassert.StructCheck(t, "testdata/fields-zip4.json", &geocodio.Fields{})
	jsonassert.StructCheck(t, "testdata/singleResponse.json", &geocodio.GeocodeResult{})
	jsonassert.StructCheck(t, "testdata/singleResponseWithAllFields.json", &geocodio.GeocodeResult{})
}
