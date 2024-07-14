package schemas

type Skill struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ParentId    *int   `json:"parent_id"`
}

type SkillConflict struct {
	Id       int    `json:"id"`
	Skill1Id *int   `json:"skill_1_id"`
	Skill2Id *int   `json:"skill_2_id"`
	Comment  string `json:"comment"`
	Priority int    `json:"priority"`
}

type SkillDomain struct {
	Id       int    `json:"id"`
	DomainId int    `json:"domain_id"`
	SkillId  int    `json:"skill_id"`
	Comments string `json:"comments"`
	Priority int    `json:"priority"`
}
