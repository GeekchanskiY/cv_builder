package schemas

type FullDatabaseData struct {
	Companies                 []*Company                 `json:"companies"`
	CVs                       []*CV                      `json:"cvs"`
	CVProjects                []*CVProject               `json:"cv_projects"`
	CVProjectResponsibilities []*CVProjectResponsibility `json:"cv_responsibilities"`
	Domains                   []*Domain                  `json:"domains"`
	Employees                 []*Employee                `json:"employees"`
	Projects                  []*Project                 `json:"projects"`
	ProjectDomains            []*ProjectDomain           `json:"project_domains"`
	Responsibilities          []*Responsibility          `json:"responsibilities"`
	ResponsibilityConflicts   []*ResponsibilityConflict  `json:"responsibility_conflicts"`
	ResponsibilitySynonyms    []*ResponsibilitySynonym   `json:"responsibility_synonyms"`
	Skills                    []*Skill                   `json:"skills"`
	SkillDomains              []*SkillDomain             `json:"skill_domains"`
	SkillConflicts            []*SkillConflict           `json:"skill_conflicts"`
	Vacancies                 []*Vacancy                 `json:"vacancies"`
	VacancyDomains            []*VacancyDomain           `json:"vacancy_domains"`
	VacancySkills             []*VacancySkill            `json:"vacancy_skills"`
}
