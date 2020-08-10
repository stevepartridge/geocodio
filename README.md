# geocodio

[![GoDoc](https://godoc.org/github.com/stevepartridge/geocodio?status.svg)](https://godoc.org/github.com/stevepartridge/geocodio)
[![Go Report Card](https://goreportcard.com/badge/github.com/stevepartridge/geocodio)](https://goreportcard.com/report/github.com/stevepartridge/geocodio)
[![Coverage](http://gocover.io/_badge/github.com/stevepartridge/service)](http://gocover.io/github.com/stevepartridge/geocodio)

Go client for [Geocodio](http://geocod.io) API v1

## Usage

### Geocode

```go
import(
  "github.com/stevepartridge/geocodio"
  "fmt"
)

func main() {
	gc, err := geocodio.New("YOUR_API_KEY")
	// or if env var GEOCODIO_API_KEY is set
	// gc, err := geocodio.New()
	if err != nil {
		panic(err)
	}
	result, err := gc.Geocode("42370 Bob Hope Dr, Rancho Mirage, CA")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Geocode Result %v", result)
}
```

### Reverse Geocode

```go
import(
  "github.com/stevepartridge/geocodio"
  "fmt"
)

func main() {
	gc, err := geocodio.NewGeocodio("YOUR_API_KEY")
	// or if env var GEOCODIO_API_KEY is set
	// gc, err := geocodio.New()
	if err != nil {
		panic(err)
	}
	result, err := gc.ReverseGeocode(38.9002898, -76.9990361)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Reverse Geocode Result %v", result)
}
```

## Tests

You can run the tests leveraging your API key as an enviroment variable from terminal (\*nix).

```
API_KEY=<YOUR_API_KEY> go test -v -cover
```
