package geocodio

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
	Meta         CensusMeta   `json:"meta"`
	Demographics *Demographic `json:"demographics,omitempty"`
	Economics    *Economics   `json:"economics,omitempty"`
	Families     *Families    `json:"families,omitempty"`
	Housing      *Housing     `json:"housing,omitempty"`
	Social       *Social      `json:"social,omitempty"`
}

type Economics struct {
	NumberOfHouseholds    NumberOfHouseholds    `json:"Number of households"`
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
	Source              string `json:"source,omitempty"`
	SurveyYears         string `json:"survey_years,omitempty"`
	SurveyDurationYears string `json:"survey_duration_years,omitempty"`
	TableID             string `json:"table_id,omitempty"`
	Universe            string `json:"universe,omitempty"`
}

type CensusDataPoint struct {
	Value         float64          `json:"value,omitempty"`
	MarginOfError float64          `json:"margin_of_error,omitempty"`
	Percentage    float64          `json:"percentage,omitempty"`
	TableID       string           `json:"table_id,omitempty"`
	Universe      string           `json:"universe,omitempty"`
	Total         *CensusDataPoint `json:"Total,omitempty"`
}

type Families struct {
	HouseholdTypeByHousehold  map[string]CensusDataPoint `json:"Household type by household"`
	HouseholdTypeByPopulation map[string]CensusDataPoint `json:"Household type by population"`
	MaritalStatus             map[string]CensusDataPoint `json:"Marital status"`
}

type Housing struct {
	NumberOfHousingUnits            map[string]CensusDataPoint `json:"Number of housing units"`
	OccupancyStatus                 map[string]CensusDataPoint `json:"Occupancy status"`
	OwnershipOfOccupiedUnits        map[string]CensusDataPoint `json:"Ownership of occupied units"`
	UnitsInStructure                map[string]CensusDataPoint `json:"Units in structure"`
	MedianValueOfOwnerOccupiedUnits map[string]CensusDataPoint `json:"Median value of owner-occupied housing units"`
	ValueOfOwnerOccupiedUnits       map[string]CensusDataPoint `json:"Value of owner-occupied housing units"`
}

type Social struct {
	PopulationByMinimumLevelOfEducation map[string]CensusDataPoint `json:"Population by minimum level of education"`
	PopulationWithVeteran               map[string]CensusDataPoint `json:"Population with veteran status"`
	PeriodOfMilitaryServiceForVeterans  map[string]CensusDataPoint `json:"Period of military service for veterans"`
}
