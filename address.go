package geocodio

type Address struct {
	Query        string     `json:"query"`
	Components   Components `json:"address_components"`
	Formatted    string     `json:"formatted_address"`
	Location     Location   `json:"location"`
	Accuracy     float64    `json:"accuracy"`
	AccuracyType string     `json:"accuracy_type"`
	Source       string     `json:"source"`
	Fields       Fields     `json:"fields,omitempty"`
}

// Components
/*
	"address_components": {
		"number": "1109",
		"predirectional": "N",
		"street": "Highland",
		"suffix": "St",
		"formatted_street": "N Highland St",
		"city": "Arlington",
		"county": "Arlington County",
		"state": "VA",
		"zip": "22201",
		"country": "US"
	},
*/
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
	PreDirectional  string `json:"predirectional"`
	Prefix          string `json:"prefix"`
}
