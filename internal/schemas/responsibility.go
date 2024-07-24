package schemas

type Responsibility struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Priority   int    `json:"priority"`
	SkillId    int    `json:"skill_id"`
	Experience int    `json:"experience"`
	Comments   string `json:"comments"`
}

type ResponsibilityReadable struct {
	Name       string `json:"name"`
	Priority   int    `json:"priority"`
	SkillName  string `json:"skill_name"`
	Experience int    `json:"experience"`
	Comments   string `json:"comments"`
}

type ResponsibilitySynonym struct {
	Id               int    `json:"id"`
	ResponsibilityId int    `json:"responsibility_id"`
	Name             string `json:"name"`
}

type ResponsibilitySynonymReadable struct {
	ResponsibilityName string `json:"responsibility_name"`
	Name               string `json:"name"`
}

type ResponsibilityConflict struct {
	Id                int    `json:"id"`
	Responsibility1Id int    `json:"responsibility_1_id"`
	Responsibility2Id int    `json:"responsibility_2_id"`
	Comment           string `json:"comment"`
	Priority          int    `json:"priority"`
}

type ResponsibilityConflictReadable struct {
	Responsibility1Name string `json:"responsibility_1_name"`
	Responsibility2Name string `json:"responsibility_2_name"`
	Comment             string `json:"comment"`
	Priority            int    `json:"priority"`
}
