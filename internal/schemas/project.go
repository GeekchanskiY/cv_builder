package schemas

type Project struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CompanyId   int    `json:"company_id"`
}

type ProjectSkill struct {
	Id        int `json:"id"`
	SkillId   int `json:"skill_id"`
	ProjectId int `json:"project_id"`
	Years     int `json:"years"`
}
