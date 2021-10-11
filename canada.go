package geocodio

type Riding struct {
	Code        string `json:"code"`
	NameFrench  string `json:"name_french"`
	NameEnglish string `json:"name_english"`
}

type Statcan struct {
	Division                CanadaDivision                `json:"division"`
	ConsolidatedSubdivision CanadaConsolidatedSubdivision `json:"consolidated_subdivision"`
	Subdivision             CanadaSubdivision             `json:"subdivision"`
	EconomicRegion          string                        `json:"economic_region"`
	StatisticalArea         CanadaStatisticalArea         `json:"statistical_area"`
	CensusMetroArea         CanadaCensusMetroArea         `json:"cma_ca"`
	Tract                   string                        `json:"tract"`
	CensusYear              int                           `json:"census_year"`
}

type CanadaDivision struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Type            string `json:"type"`
	TypeDescription string `json:"type_description"`
}

type CanadaConsolidatedSubdivision struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CanadaSubdivision struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	TypeCode        string `json:"type"`
	TypeDescription string `json:"type_description"`
}

type CanadaStatisticalArea struct {
	Code            string `json:"code"`
	CodeDescription string `json:"code_description"`
	Type            string `json:"type"`
	TypeDescription string `json:"type_description"`
}

type CanadaCensusMetroArea struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Type            string `json:"type"`
	TypeDescription string `json:"type_description"`
}
