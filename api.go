package geocodio

import (
	"encoding/json"
	"errors"
	// "fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Basic call (GET) method to API
func (g *Geocodio) Call(path string, query map[string]string) (GeocodeResult, error) {

	if strings.Index(path, "/") != 0 {
		return GeocodeResult{}, errors.New("Path must start with a forward slash: ' / ' ")
	}

	_url := GeocodioAPIBaseURLv1 + path + "?api_key=" + g.ApiKey

	for k, v := range query {
		_url = _url + "&" + k + "=" + url.QueryEscape(v)
	}

	resp, err := http.Get(_url)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return GeocodeResult{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return GeocodeResult{}, err
	}

	result := GeocodeResult{}

	result.Debug.RequestedURL = _url
	result.Debug.Status = resp.Status
	result.Debug.StatusCode = resp.StatusCode

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, err
}
