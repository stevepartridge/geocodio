package geocodio

// Fields
type Fields struct {
	Timezone                  Timezone                  `json:"timezone,omitempty"`
	Zip4                      Zip4                      `json:"zip4,omitempty"`
	CongressionalDistrict     CongressionalDistrict     `json:"congressional_district,omitempty"`  // v1.0
	CongressionalDistricts    []CongressionalDistrict   `json:"congressional_districts,omitempty"` // v1.1+
	StateLegislativeDistricts StateLegislativeDistricts `json:"state_legislative_districts,omitempty"`
	SchoolDistricts           SchoolDistricts           `json:"school_districts,omitempty"`
	Census                    CensusResults             `json:"census,omitempty"`
	ACS                       CensusACS                 `json:"acs,omitempty"`
}
