package schemas

import "time"

type CV struct {
	Id         int `json:"id"`
	VacancyId  int `json:"vacancy_id"`
	EmployeeId int `json:"employee_id"`

	// Need to confirm real employee experience
	IsReal bool `json:"is_real"`
}

type CVProject struct {
	Id        int `json:"id"`
	CVId      int `json:"cv_id"`
	ProjectId int `json:"project_id"`
	CompanyId int `json:"company_id"`

	// Work experience on project
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// Responsibility that will be used in CV
type CVProjectRespobsibility struct {
	Id               int `json:"id"`
	ResponsibilityId int `json:"responsibility_id"`
	CVProjectId      int `json:"cv_project_id"`
	Priority         int `json:"priority"`
}
