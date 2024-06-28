package schemas

type Skill struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ParentId    int    `json:"parent_id"`
}

type SkillConflict struct {
	Id       int    `json:"id"`
	Skill1Id int    `json:"skill_1_id"`
	Skill2Id int    `json:"skill_2_id"`
	Comment  string `json:"comment"`
	Priority int    `json:"priority"`
}
