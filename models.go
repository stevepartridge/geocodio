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

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

// Fields
type Fields struct {
	Timezone                  Timezone                  `json:"timezone,omitempty"`
	CongressionalDistrict     CongressionalDistrict     `json:"congressional_district,omitempty"`  // v1.0
	CongressionalDistricts    []CongressionalDistrict   `json:"congressional_districts,omitempty"` // v1.1+
	StateLegislativeDistricts StateLegislativeDistricts `json:"state_legislative_districts,omitempty"`
	SchoolDistricts           SchoolDistricts           `json:"school_districts,omitempty"`
	Census                    Census                    `json:"census,omitempty"`
	ACS                       CensusACS                 `json:"acs,omitempty"`
}

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

type SchoolDistricts struct {
	Unified    SchoolDistrict `json:"unified"`
	Elementary SchoolDistrict `json:"elementary"`
	Secondar   SchoolDistrict `json:"secondary"`
}

// SchoolDistrict field
/*
{
	"name": "Desert Sands Unified School District",
	"lea_code": "11110",
	"grade_low": "KG",
	"grade_high": "12"
}
*/
type SchoolDistrict struct {
	Name      string `json:"name"`
	LEACode   string `json:"lea_code"`
	GradeLow  string `json:"grade_low"`
	GradeHigh string `json:"grade_high"`
}

type CensusResults struct {
	Census2010 *Census `json:"2010,omitempty"`
	Census2011 *Census `json:"2011,omitempty"`
	Census2012 *Census `json:"2012,omitempty"`
	Census2013 *Census `json:"2013,omitempty"`
	Census2014 *Census `json:"2014,omitempty"`
	Census2015 *Census `json:"2015,omitempty"`
	Census2016 *Census `json:"2016,omitempty"`
	Census2017 *Census `json:"2017,omitempty"`
	Census2018 *Census `json:"2018,omitempty"`
	Census2019 *Census `json:"2019,omitempty"`
	Census2020 *Census `json:"2020,omitempty"`
}

// Census field
/*
{
	"census_year": 2010,
	"state_fips": "51",
	"county_fips": "51013",
	"tract_code": "101801",
	"block_code": "1004",
	"block_group": "1",
	"full_fips": "510131018011004",
	"place": {
		"name": "Arlington",
		"fips": "5103000"
	},
	"metro_micro_statistical_area": {
		"name": "Washington-Arlington-Alexandria, DC-VA-MD-WV",
		"area_code": "47900",
		"type": "metropolitan"
	},
	"combined_statistical_area": {
		"name": "Washington-Baltimore-Northern Virginia, DC-MD-VA-WV",
		"area_code": "51548"
	},
	"metropolitan_division": {
		"name": "Washington-Arlington-Alexandria, DC-VA-MD-WV",
		"area_code": "47894"
	},
	"source": "US Census Bureau"
}
*/
type Census struct {
	Year                      int             `json:"census_year"`
	StateFIPS                 string          `json:"state_fips"`
	CountyFIPS                string          `json:"county_fips"`
	TractCode                 string          `json:"tract_code"`
	BlockCode                 string          `json:"block_code"`
	BlockGroup                string          `json:"block_group"`
	FullFIPS                  string          `json:"full_fips"`
	Place                     Place           `json:"place"`
	MetroMicroStatisticalArea StatisticalArea `json:"metro_micro_statistical_area"`
	CombinedStatisticalArea   StatisticalArea `json:"combined_statistical_area"`
	MetropolitanDivision      StatisticalArea `json:"metropolitan_division"`
	Source                    string          `json:"source"`
}

type Place struct {
	Name string `json:"name"`
	FIPS string `json:"fips"`
}

type StatisticalArea struct {
	Name     string `json:"name"`
	AreaCode string `json:"area_code"`
	Type     string `json:"type,omitempty"`
}

type CensusACS struct {
	Meta         CensusMeta  `json:"meta"`
	Demographics Demographic `json:"demographics"`
	Economics    Economics   `json:"economics`
}

type Economics struct {
	NumberOfHouseholds    struct{}              `json:"Number of households"`
	MedianHouseholdIncome MedianHouseholdIncome `json:"Median household income"`
	HouseholdIncome       HouseholdIncome       `json:"Household income"`
}

type NumberOfHouseholds struct {
	Meta  CensusMeta      `json:"meta"`
	Total CensusDataPoint `json:"Total"`
}

type MedianHouseholdIncome struct {
	Meta  CensusMeta      `json:"meta"`
	Total CensusDataPoint `json:"Total"`
}

type HouseholdIncome struct {
	Meta                 CensusMeta      `json:"meta"`
	LessThan10000        CensusDataPoint `json:"Less than $10,000"`
	Income10000to14999   CensusDataPoint `json:"$10,000 to $14,999"`
	Income15000to19999   CensusDataPoint `json:"$15,000 to $19,999"`
	Income20000to24999   CensusDataPoint `json:"$20,000 to $24,999"`
	Income25000to29999   CensusDataPoint `json:"$25,000 to $29,999"`
	Income30000to34999   CensusDataPoint `json:"$30,000 to $34,999"`
	Income35000to39999   CensusDataPoint `json:"$35,000 to $39,999"`
	Income40000to44999   CensusDataPoint `json:"$40,000 to $44,999"`
	Income45000to49999   CensusDataPoint `json:"$45,000 to $49,999"`
	Income50000to59000   CensusDataPoint `json:"$50,000 to $59,999"`
	Income60000to74999   CensusDataPoint `json:"$60,000 to $74,999"`
	Income75000to99999   CensusDataPoint `json:"$75,000 to $99,999"`
	Income100000to124999 CensusDataPoint `json:"$100,000 to $124,999"`
	Income125000to149000 CensusDataPoint `json:"$125,000 to $149,999"`
	Income150000to199999 CensusDataPoint `json:"$150,000 to $199,999"`
	Income200000orMore   CensusDataPoint `json:"$200,000 or more"`
}

type Demographic struct {
	MedianAge            map[string]CensusDataPoint `json:"Median age"`
	PopulationByAgeRange map[string]CensusDataPoint `json:"Population by age range"`
	Sex                  map[string]CensusDataPoint `json:"Sex"`
	RaceAndEthnicity     map[string]CensusDataPoint `json:"Race and ethnicity"`
}

type CensusMeta struct {
	Source              string `json:"source"`
	SurveyYears         string `json:"survey_years"`
	SurveyDurationYears string `json:"survey_duration_years"`
	TableID             string `json:"table_id"`
	Universe            string `json:"universe"`
}

type CensusDataPoint struct {
	Value         float64 `json:"value"`
	MarginOfError float64 `json:"margin_of_error"`
	Percentage    float64 `json:"percentage"`
	TableID       string  `json:"table_id"`
	Universe      string  `json:"universe"`
}

type Input struct {
	AddressComponents Components `json:"address_components"`
	FormattedAddress  string     `json:"formatted_address"`
}

// GeocodeResponse
type GeocodeResult struct {
	Input   Input     `json:"input,omitempty"`
	Results []Address `json:"results"`
	Debug   struct {
		RawResponse  []byte `json:"-"`
		RequestedURL string `json:"requested_url"`
		Status       string `json:"status"`
		StatusCode   int    `json:"status_code"`
	} `json:"-"`
}

func (self *GeocodeResult) ResponseAsString() string {
	return string(self.Debug.RawResponse)
}
