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

type VacancyReadable struct {
	Name        string `json:"name"`
	CompanyName string `json:"company_name"`
	Link        string `json:"link"`
	Description string `json:"description"`
	PublishedAt string `json:"published_at"`
	Experience  int    `json:"experience"`
}

type VacancySkill struct {
	Id        int `json:"id"`
	SkillId   int `json:"skill_id"`
	VacancyId int `json:"vacancy_id"`

	// Used to check if this skill is required in a CV
	Priority int `json:"priority"`
}

type VacancySkillReadable struct {
	SkillName   string `json:"skill_name"`
	VacancyName string `json:"vacancy_name"`
	Priority    int    `json:"priority"`
}

type VacancyDomain struct {
	Id        int `json:"id"`
	VacancyId int `json:"vacancy_id"`
	DomainId  int `json:"domain_id"`

	// Used to check if this domain is required in a CV
	Priority int `json:"priority"`
}

type VacancyDomainReadable struct {
	VacancyName string `json:"vacancy_name"`
	DomainName  string `json:"domain_name"`

	// Used to check if this domain is required in a CV
	Priority int `json:"priority"`
}
