package schemas

import "time"

type CV struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	VacancyId  int    `json:"vacancy_id"`
	EmployeeId int    `json:"employee_id"`

	// Need to confirm real employee experience
	IsReal bool `json:"is_real"`
}

type CVReadable struct {
	Name         string `json:"name"`
	VacancyName  string `json:"vacancy_name"`
	EmployeeName string `json:"employee_name"`

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

// CVService is used to build imaginary architecture to insert all data without conflicts
type CVService struct {
	Id          int    `json:"id"`
	CVProjectId int    `json:"cv_project_id"`
	Name        string `json:"name"`
}

// CVServiceResponsibility is used to add responsibility to the CVService => to CV actually
type CVServiceResponsibility struct {
	Id               int `json:"id"`
	CVServiceId      int `json:"cv_service_id"`
	ResponsibilityId int `json:"responsibility_id"`

	// Order used to keep the same order of cv responsibilities between document generations
	Order int `json:"order"`
}
