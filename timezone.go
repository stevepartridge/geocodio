package geocodio

// Timezone based on this payload
/*
	"timezone": {
    "name": "America/New_York",
    "utc_offset": -5,
    "observes_dst": true,
    "abbreviation": "EST",
    "source": "© OpenStreetMap contributors"
  }
*/

type Timezone struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"` // v1.3+
	UTCOffset    int    `json:"utc_offset"`
	ObservesDST  bool   `json:"observes_dst"`
	Source       string `json:"source"`
}
