package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

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

func (c *EmployeeController) GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	employees, err := c.employeeRepo.GetAll()
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

func (c *EmployeeController) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	employee := schemas.Employee{}

	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	uid, err := c.employeeRepo.Create(employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if uid != 0 {
		employee.Id = uid
	}

	w.WriteHeader(http.StatusCreated)
	b, err := json.Marshal(employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(b)

}

func (c *EmployeeController) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	employee := schemas.Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.employeeRepo.Update(employee)
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

func (c *EmployeeController) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	employee := schemas.Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.employeeRepo.Delete(employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	b, err := json.Marshal(employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (c *EmployeeController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var employee schemas.Employee
	employee_id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid employee id"))
		return
	}

	employee, err = c.employeeRepo.Get(int(employee_id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid employee id"))
		return
	}

	b, err := json.Marshal(employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}
