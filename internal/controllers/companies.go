package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GeekchanskiY/cv_builder/internal/repository"
	"github.com/GeekchanskiY/cv_builder/internal/schemas"
	"github.com/julienschmidt/httprouter"
)

type CompanyController struct {
	companyRepo *repository.CompanyRepository
}

func CreateComapnyController(repo *repository.CompanyRepository) *CompanyController {
	return &CompanyController{
		companyRepo: repo,
	}
}

func (c *CompanyController) GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	companies, err := c.companyRepo.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(companies)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (c *CompanyController) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	company := schemas.Company{}

	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	uid, err := c.companyRepo.Create(company)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if uid != 0 {
		company.Id = uid
	}

	b, err := json.Marshal(company)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func (c *CompanyController) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	comapny := schemas.Company{}
	err := json.NewDecoder(r.Body).Decode(&comapny)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.companyRepo.Update(comapny)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(comapny)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func (c *CompanyController) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	company := schemas.Company{}
	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.companyRepo.Delete(company)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(company)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write(b)
}

func (c *CompanyController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var company schemas.Company
	comapny_id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid employee id"))
		return
	}

	company, err = c.companyRepo.Get(int(comapny_id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid employee id"))
		return
	}

	b, err := json.Marshal(company)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}
