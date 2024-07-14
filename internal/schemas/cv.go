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

type CVProject struct {
	Id        int `json:"id"`
	CVId      int `json:"cv_id"`
	ProjectId int `json:"project_id"`
	CompanyId int `json:"company_id"`
	Years     int `json:"years"`
}
