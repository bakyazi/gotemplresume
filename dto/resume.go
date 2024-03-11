package dto

type Resume struct {
	WorkExperiences []ExperienceItem `json:"workExperiences,omitempty"`
	Education       []ExperienceItem `json:"education,omitempty"`
	Skills          Skills           `json:"skills"`
}

type ExperienceItem struct {
	StartDate   string `json:"startDate,omitempty"`
	EndDate     string `json:"endDate,omitempty"`
	Header      string `json:"header,omitempty"`
	Location    string `json:"location,omitempty"`
	Description string `json:"description,omitempty"`
}

type Skills struct {
	ProgrammingLanguages []Skill `json:"programmingLanguages,omitempty"`
	Frameworks           []Skill `json:"frameworks,omitempty"`
	OperatingSystems     []Skill `json:"operatingSystems,omitempty"`
	CiCd                 []Skill `json:"ciCd,omitempty"`
	Tools                []Skill `json:"tools,omitempty"`
}

type Skill struct {
	Name string  `json:"name,omitempty"`
	Rank float64 `json:"rank,omitempty"`
	Key  string  `json:"key,omitempty"`
}
