package geocodio

import "errors"

var (
	// ErrReverseGecodeMissingLatLng error when a lat/lng is not provided
	ErrReverseGecodeMissingLatLng = errors.New("Latitude and longitude must not be empty")
	// ErrMissingAPIKey error
	ErrMissingAPIKey = errors.New("Missing or empty API key")
	// ErrAddressIsEmpty error
	ErrAddressIsEmpty = errors.New("Address must not be empty")
	// ErrBatchAddressesIsEmpty error
	ErrBatchAddressesIsEmpty = errors.New("At least one address is required for batch query")
	// ErrReverseBatchMissingCoords error
	ErrReverseBatchMissingCoords = errors.New("Missing minimum coordinates")
	// ErrReverseBatchInvalidCoordsPairs error
	ErrReverseBatchInvalidCoordsPairs = errors.New("Invalid list of coordinate pairs")
	// ErrNoResultsFound
	ErrNoResultsFound = errors.New("No results found")
)
