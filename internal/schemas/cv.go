package schemas

type CV struct {
	Id         int `json:"id"`
	VacancyId  int `json:"vacancy_id"`
	EmployeeId int `json:"employee_id"`

	// Need to confirm real employee experience
	IsReal bool `json:"is_real"`
}

type CvDomain struct {
	Id       int `json:"id"`
	DomainId int `json:"domain_id"`
	CvId     int `json:"cv_id"`
}

type CVSkillExtension struct {
	SkillId     int    `json:"skill_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ParentId    *int   `json:"parent_id"`
	Years       int    `json:"years"`
}
