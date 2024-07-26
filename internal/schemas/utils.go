package schemas

type FullDatabaseData struct {
	Companies []Company `json:"companies"`
	// CVs                       []CV                      `json:"cvs"`
	// CVProjects                []CVProject               `json:"cv_projects"`
	// CVServices                []CVService               `json:"cv_services"`
	// CVServiceResponsibilities []CVServiceResponsibility `json:"cv_service_responsibilities"`
	Domains                 []Domain                         `json:"domains"`
	Employees               []Employee                       `json:"employees"`
	Projects                []Project                        `json:"projects"`
	ProjectDomains          []ProjectDomainReadable          `json:"project_domains"`
	ProjectServices         []ProjectServiceReadable         `json:"project_services"`
	Responsibilities        []ResponsibilityReadable         `json:"responsibilities"`
	ResponsibilityConflicts []ResponsibilityConflictReadable `json:"responsibility_conflicts"`
	ResponsibilitySynonyms  []ResponsibilitySynonymReadable  `json:"responsibility_synonyms"`
	Skills                  []SkillReadable                  `json:"skills"`
	SkillDomains            []SkillDomainReadable            `json:"skill_domains"`
	SkillConflicts          []SkillConflictReadable          `json:"skill_conflicts"`
	Vacancies               []VacancyReadable                `json:"vacancies"`
	VacancyDomains          []VacancyDomainReadable          `json:"vacancy_domains"`
	VacancySkills           []VacancySkillReadable           `json:"vacancy_skills"`
}
