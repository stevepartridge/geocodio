package geocodio

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
