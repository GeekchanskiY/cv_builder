package requests

type BuildRequest struct {
	EmployeeID int `json:"employee_id"`
	VacancyID  int `json:"vacancy_id"`
}
