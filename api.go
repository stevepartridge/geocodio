package geocodio

import (
	"encoding/json"
	"errors"
	"fmt"

	// "fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	MethodGet  = "GET"
	MethodPost = "POST"
)

// Call uses basic (GET) method to make a request to the API
func (g *Geocodio) Call(path string, query map[string]string) (GeocodeResult, error) {

	if strings.Index(path, "/") != 0 {
		return GeocodeResult{}, errors.New("Path must start with a forward slash: ' / ' ")
	}

	_url := GeocodioAPIBaseURLv1 + path + "?api_key=" + g.APIKey

	for k, v := range query {
		_url = _url + "&" + k + "=" + url.QueryEscape(v)
	}

	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(_url)

	if err != nil {
		return GeocodeResult{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return GeocodeResult{}, err
	}

	result := GeocodeResult{}

	result.Debug.RequestedURL = _url
	result.Debug.Status = resp.Status
	result.Debug.StatusCode = resp.StatusCode
	result.Debug.RawResponse = body

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, err
}

func (g *Geocodio) get(path string, query map[string]string) (GeocodeResult, error) {
	return g.call(MethodGet, path, nil, query)
}
func (g *Geocodio) post(path string, payload interface{}, query map[string]string) (GeocodeResult, error) {
	return g.call(MethodPost, path, payload, query)
}

func (g *Geocodio) call(method, path string, payload interface{}, query map[string]string) (GeocodeResult, error) {

	if strings.Index(path, "/") != 0 {
		return GeocodeResult{}, errors.New("Path must start with a forward slash: ' / ' ")
	}

	rawURL := GeocodioAPIBaseURLv1 + path + "?api_key=" + g.APIKey

	if query != nil {
		for k, v := range query {
			rawURL = fmt.Sprintf("%s&%s=%s", rawURL, k, url.QueryEscape(v))
		}
	}

	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return GeocodeResult{}, nil
	}

	req := http.Request{
		Method: method,
		URL:    u,
	}

	resp, err := client.Do(&req)

	if err != nil {
		return GeocodeResult{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return GeocodeResult{}, err
	}

	result := GeocodeResult{}

	result.Debug.RequestedURL = u.String()
	result.Debug.Status = resp.Status
	result.Debug.StatusCode = resp.StatusCode
	result.Debug.RawResponse = body

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, err
}
