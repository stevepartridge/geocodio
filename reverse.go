package geocodio

import (
	"fmt"
	"strconv"
)

/*
	See: http://geocod.io/docs/#toc_16
*/
// Reverse does a reverse geocode look up for a single coordinate
func (g *Geocodio) Reverse(latitude, longitude float64) (GeocodeResult, error) {
	// if there is an address here, they should probably think about moving
	// regardless, we'll consider it an error
	if latitude == 0.0 && longitude == 0.0 {
		return GeocodeResult{}, ErrReverseGecodeMissingLatLng
	}

	latStr := strconv.FormatFloat(latitude, 'f', 9, 64)
	lngStr := strconv.FormatFloat(longitude, 'f', 9, 64)

	resp := GeocodeResult{}
	err := g.do("GET", "/reverse", map[string]string{"q": latStr + "," + lngStr}, nil, &resp)
	if err != nil {
		return resp, err
	}

	if len(resp.Results) == 0 {
		return resp, ErrNoResultsFound
	}

	return resp, nil
}

// ReverseGeocode is deprecated and will be removed in 2+
func (g *Geocodio) ReverseGeocode(latitude, longitude float64) (GeocodeResult, error) {
	fmt.Printf(`
  ReverseGeocode(%f, %f) is deprecated and will be removed in 2+
  Use Reverse(%f, %f) 
`,
		latitude, longitude,
		latitude, longitude,
	)
	return g.Reverse(latitude, longitude)
}

// ReverseBatch supports a batch lookup by lat/lng coordinate pairs
func (g *Geocodio) ReverseBatch(latlngs ...float64) (BatchResponse, error) {
	resp := BatchResponse{}
	if len(latlngs) == 0 {
		return resp, ErrReverseBatchMissingCoords
	}

	if len(latlngs)%2 == 1 {
		return resp, ErrReverseBatchInvalidCoordsPairs
	}

	var (
		payload = []string{}
		pair    string
	)

	for i := range latlngs {
		coord := strconv.FormatFloat(latlngs[i], 'f', 9, 64)
		if i == 0 {
			pair = coord
			continue
		}
		if i%2 == 0 {
			pair = fmt.Sprintf("%s,%s", pair, coord)
			payload = append(payload, pair)
			continue
		}
		pair = coord
	}

	err := g.do("POST", "/reverse", nil, payload, &resp)
	if err != nil {
		return resp, err
	}

	if len(resp.Results) == 0 {
		return resp, ErrNoResultsFound
	}

	return resp, nil

}
