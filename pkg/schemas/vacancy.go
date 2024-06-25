package schemas

type Vacancy struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	CompanyId   int    `json:"company_id"`
	Link        string `json:"link"`
	Description string `json:"description"`
	PublishedAt string `json:"published_at"`
	Experience  int    `json:"experience"`
}

type VacancySkill struct {
	Id        int `json:"id"`
	SkillId   int `json:"skill_id"`
	VacancyId int `json:"vacancy_id"`
}
