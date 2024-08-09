package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/solndev/cv_builder/internal/utils"

	"github.com/julienschmidt/httprouter"
	"github.com/solndev/cv_builder/internal/repository"
	"github.com/solndev/cv_builder/internal/schemas"
)

type EmployeeController struct {
	employeeRepo *repository.EmployeeRepository
}

func CreateEmployeeController(repo *repository.EmployeeRepository) *EmployeeController {
	return &EmployeeController{
		employeeRepo: repo,
	}
}

func (c *EmployeeController) GetAll(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	employees, err := c.employeeRepo.GetAll()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
	b, err := json.Marshal(employees)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *EmployeeController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	employee := schemas.Employee{}

	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	uid, err := c.employeeRepo.Create(employee)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	if uid != 0 {
		employee.Id = uid
	}

	b, err := json.Marshal(employee)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *EmployeeController) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	employee := schemas.Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.employeeRepo.Update(employee)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(employee)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *EmployeeController) Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	employee := schemas.Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.employeeRepo.Delete(employee)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *EmployeeController) Get(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {

	var employee schemas.Employee
	employeeId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	employee, err = c.employeeRepo.Get(employeeId)

	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(employee)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}
