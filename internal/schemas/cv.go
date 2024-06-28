package schemas

type CV struct {
	Id         int `json:"id"`
	VacancyId  int `json:"vacancy_id"`
	EmployeeId int `json:"employee_id"`
}

type CvSkill struct {
	Id      int `json:"id"`
	SkillId int `json:"skill_id"`
	CvId    int `json:"cv_id"`
	Years   int `json:"years"`
}

type CvDomain struct {
	Id       int `json:"id"`
	DomainId int `json:"domain_id"`
	CvId     int `json:"cv_id"`
}
