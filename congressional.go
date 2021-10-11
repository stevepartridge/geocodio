package geocodio

// Congressional District field
type CongressionalDistrict struct {
	Name               string       `json:"name"`
	DistrictNumber     int          `json:"district_number"`
	CongressNumber     string       `json:"congress_number"`
	CongressYears      string       `json:"congress_years"`
	Proportion         int          `json:"proportion"`
	CurrentLegislators []Legislator `json:"current_legislators"` // v1.2+
}

// Legislator field
type Legislator struct {
	Type       string              `json:"type"`
	Bio        Bio                 `json:"bio"`
	Contact    Contact             `json:"contact"`
	Social     CongressionalSocial `json:"social"`
	References References          `json:"references"`
	Source     string              `json:"source"`
}

// Bio field
type Bio struct {
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Birthday  string `json:"birthday"`
	Gender    string `json:"gender"`
	Party     string `json:"party"`
}

// Contact field
type Contact struct {
	URL         string `json:"url"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	ContactForm string `json:"contact_form"`
}

// Social field
type CongressionalSocial struct {
	RSSURL    string `json:"rss_url"`
	Twitter   string `json:"twitter"`
	Facebook  string `json:"facebook"`
	YouTube   string `json:"youtube"`
	YouTubeID string `json:"youtube_id"`
}

// References field
type References struct {
	BioguideID       string `json:"bioguide_id"`
	ThomasID         string `json:"thomas_id"`
	OpenSecretsID    string `json:"opensecrets_id"`
	LISID            string `json:"lis_id"`
	CSPANID          string `json:"cspan_id"`
	GovTrackID       string `json:"govtrack_id"`
	VoteSmartID      string `json:"votesmart_id"`
	BallotpediaID    string `json:"ballotpedia_id"`
	WashingtonPostID string `json:"washington_post_id"`
	ICPSRID          string `json:"icpsr_id"`
	WikipediaID      string `json:"wikipedia_id"`
}

type StateLegislativeDistricts struct {
	House  StateLegislativeDistrict `json:"house"`
	Senate StateLegislativeDistrict `json:"senate"`
}

type StateLegislativeDistrict struct {
	Name           string `json:"name"`
	DistrictNumber int    `json:"district_number"`
}
