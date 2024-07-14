package schemas

type Vacancy struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	CompanyId   int    `json:"company_id"`
	Link        string `json:"link"`
	Description string `json:"description"`
	PublishedAt string `json:"published_at"`

	// Required experience in years
	Experience int `json:"experience"`
}

type VacancySkill struct {
	Id        int `json:"id"`
	SkillId   int `json:"skill_id"`
	VacancyId int `json:"vacancy_id"`

	// Used to check if this skill is required in a CV
	Priority int `json:"priority"`
}

type VacancyDomain struct {
	Id        int `json:"id"`
	VacancyId int `json:"vacancy_id"`
	DomainId  int `json:"domain_id"`
	Priority  int `json:"priority"`
}
