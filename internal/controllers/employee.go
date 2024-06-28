package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal/repository"
	"github.com/GeekchanskiY/cv_builder/internal/schemas"
	"github.com/julienschmidt/httprouter"
)

type EmployeeController struct {
	employeeRepo *repository.EmployeeRepository
}

func CreateEmployeeController(repo *repository.EmployeeRepository) *EmployeeController {
	return &EmployeeController{
		employeeRepo: repo,
	}
}

func (c *EmployeeController) GetEmployees(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	employees, err := c.employeeRepo.GetEmployees()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(employees)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)

}

func (c *EmployeeController) CreateEmployee(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	employee := schemas.Employee{}

	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.employeeRepo.CreateEmployee(employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	b, err := json.Marshal(employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)

}
