package geocodio

import "errors"

var (
	// ErrReverseGecodeMissingLatLng error when a lat/lng is not provided
	ErrReverseGecodeMissingLatLng = errors.New("Latitude and longitude must not be empty")
	ErrMissingApiKey              = errors.New("Missing or empty API key")
	ErrAddressIsEmpty             = errors.New("Address must not be empty")
	ErrBatchAddressesIsEmpty      = errors.New("At least one address is required for batch query")
)
