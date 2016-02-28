package geocodio

// Geocodio is the base struct
type Geocodio struct {
	APIKey string
}

type Address struct {
	Components   Components `json:"address_components"`
	Formatted    string     `json:"formatted_address"`
	Location     Location   `json:"location"`
	Accuracy     float64    `json:"accuracy"`
	AccuracyType string     `json:"accuracy_type"`
	Source       string     `json:"source"`
	Fields       Fields     `json:"fields,omitempty"`
}

type Components struct {
	Number          string `json:"number"`
	Street          string `json:"street"`
	Suffix          string `json:"suffix"`
	SecondaryNumber string `json:"secondarynumber"`
	SecondaryUnit   string `json:"secondaryunit"`
	PostDirectional string `json:"postdirectional"`
	FormattedStreet string `json:"formatted_street"`
	City            string `json:"city"`
	State           string `json:"state"`
	Zip             string `json:"zip"`
	County          string `json:"county"`
	Country         string `json:"country"`
}

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

// Fields
type Fields struct {
	Timezone                  Timezone                  `json:"timezone,omitempty"`
	CongressionalDistrict     CongressionalDistrict     `json:"congressional_district,omitempty"`
	StateLegislativeDistricts StateLegislativeDistricts `json:"state_legislative_districts,omitempty"`
}

type Timezone struct {
	Name        string `json:"name"`
	UTCOffset   int    `json:"utc_offset"`
	ObservesDST bool   `json:"observes_dst"`
}

type CongressionalDistrict struct {
	Name           string `json:"name"`
	DistrictNumber int    `json:"district_number"`
	CongressNumber string `json:"congress_number"`
	CongressYears  string `json:"congress_years"`
}

type StateLegislativeDistricts struct {
	House  StateLegislativeDistrict `json:"house"`
	Senate StateLegislativeDistrict `json:"senate"`
}

type StateLegislativeDistrict struct {
	Name           string `json:"name"`
	DistrictNumber string `json:"district_number"`
}

// GeocodeResponse
type GeocodeResult struct {
	Input struct {
		AddressComponents Components `json:"address_components"`
	} `json:"input,omitempty"`
	Results []Address `json:"results"`
	Debug   struct {
		RawResponse  []byte `json:"-"`
		RequestedURL string `json:"requested_url"`
		Status       string `json:"status"`
		StatusCode   int    `json:"status_code"`
	} `json:"-"`
}
