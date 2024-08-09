package requests

import "github.com/solndev/cv_builder/internal/schemas/interfaces"

type BuildRequest struct {
	interfaces.Validatable
	EmployeeID int `json:"employee_id"`
	VacancyID  int `json:"vacancy_id"`

	// Value from 0 to 2.
	// 0 - no microservice architecture at all
	// 1 - latest 1 or 2 projects may be microservice
	// 2 - all projects will be microservices
	Microservices int `json:"microservices"`
}

func (r *BuildRequest) Validate() (errors []string) {

	if r.Microservices < 0 || r.Microservices > 2 {
		errors = append(errors, "microservices")
	}

	return errors
}
