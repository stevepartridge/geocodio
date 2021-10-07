package geocodio

type SchoolDistricts struct {
	Unified    SchoolDistrict `json:"unified"`
	Elementary SchoolDistrict `json:"elementary"`
	Secondar   SchoolDistrict `json:"secondary"`
}

// SchoolDistrict field
type SchoolDistrict struct {
	Name      string `json:"name"`
	LEACode   string `json:"lea_code"`
	GradeLow  string `json:"grade_low"`
	GradeHigh string `json:"grade_high"`
}
