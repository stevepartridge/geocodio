package geocodio

import (
	"strconv"
)

/*
	See: http://geocod.io/docs/#toc_16
*/
// ReverseGeocode does a reverse geocode look up for a single coordinate
func (g *Geocodio) ReverseGeocode(latitude, longitude float64) (GeocodeResult, error) {
	// if there is an address here, they should probably think about moving
	// regardless, we'll consider it an error
	if latitude == 0.0 && longitude == 0.0 {
		return GeocodeResult{}, ErrReverseGecodeMissingLatLng
	}

	latStr := strconv.FormatFloat(latitude, 'f', 9, 64)
	lngStr := strconv.FormatFloat(longitude, 'f', 9, 64)

	result, err := g.Call("/reverse", map[string]string{"q": latStr + "," + lngStr})
	if err != nil {
		return GeocodeResult{}, err
	}

	// fmt.Println(result.Debug.RequestedURL)

	return result, nil
}
