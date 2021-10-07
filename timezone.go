package geocodio

// Timezone based on this payload
type Timezone struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"` // v1.3+
	UTCOffset    int    `json:"utc_offset"`
	ObservesDST  bool   `json:"observes_dst"`
	Source       string `json:"source"`
}
