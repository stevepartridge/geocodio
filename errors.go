package geocodio

import "errors"

var (
	// ErrReverseGecodeMissingLatLng error when a lat/lng is not provided
	ErrReverseGecodeMissingLatLng = errors.New("Latitude and longitude must not be empty")
)
