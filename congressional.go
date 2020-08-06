package geocodio

// Congressional District field
/*
"name": "Congressional District 8",
"district_number": 8,
"congress_number": "116th",
"congress_years": "2019-2021",
"proportion": 1,
"current_legislators": [...]
*/
type CongressionalDistrict struct {
	Name               string       `json:"name"`
	DistrictNumber     int          `json:"district_number"`
	CongressNumber     string       `json:"congress_number"`
	CongressYears      string       `json:"congress_years"`
	Proportion         int          `json:"congress_years"`
	CurrentLegislators []Legislator `json:"current_legislators"` // v1.2+
}

// Legislator field
/*
{
	"type": "representative",
	"bio": {...},
	"contact": {...},
	"social": {...},
	"references": {...},
	"source": "Legislator data is originally collected and aggregated by https://github.com/unitedstates/"
}
*/
type Legislator struct {
	Type       string     `json:"type"`
	Bio        Bio        `json:"bio"`
	Contact    Contact    `json:"contact"`
	Social     Social     `json:"social"`
	References References `json:"references"`
	Source     string     `json:"source"`
}

// Bio field
/*
	"bio": {
		"last_name": "Beyer",
		"first_name": "Donald",
		"birthday": "1950-06-20",
		"gender": "M",
		"party": "Democrat"
	}
*/
type Bio struct {
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Birthday  string `json:"birthday"`
	Gender    string `json:"gender"`
	Party     string `json:"party"`
}

// Contact field
/*
"contact": {
	"url": "https://beyer.house.gov",
	"address": "1119 Longworth House Office Building Washington DC 20515-4608",
	"phone": "(202) 225-4376",
	"contact_form": null
}
*/
type Contact struct {
	URL         string `json:"url"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	ContactForm string `json:"contact_form"`
}

// Social field
/*
"social": {
	"rss_url": null,
	"twitter": "RepDonBeyer",
	"facebook": "RepDonBeyer",
	"youtube": null,
	"youtube_id": "UCPJGVbOVcAVGiBwq8qr_T9w"
}
*/
type Social struct {
	RSSURL    string `json:"rss_url"`
	Twitter   string `json:"twitter"`
	Facebook  string `json:"facebook"`
	YouTube   string `json:"youtube"`
	YouTubeID string `json:"youtube_id"`
}

// References field
/*
"references": {
	"bioguide_id": "B001292",
	"thomas_id": "02272",
	"opensecrets_id": "N00036018",
	"lis_id": null,
	"cspan_id": "21141",
	"govtrack_id": "412657",
	"votesmart_id": "1707",
	"ballotpedia_id": null,
	"washington_post_id": null,
	"icpsr_id": "21554",
	"wikipedia_id": "Don Beyer"
}
*/
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
	DistrictNumber string `json:"district_number"`
}
