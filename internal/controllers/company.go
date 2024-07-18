package controllers

import (
	"encoding/json"
	"github.com/GeekchanskiY/cv_builder/internal/utils"
	"net/http"
	"strconv"

	"github.com/GeekchanskiY/cv_builder/internal/repository"
	"github.com/GeekchanskiY/cv_builder/internal/schemas"
	"github.com/julienschmidt/httprouter"
)

type CompanyController struct {
	companyRepo *repository.CompanyRepository
}

func CreateCompanyController(repo *repository.CompanyRepository) *CompanyController {
	return &CompanyController{
		companyRepo: repo,
	}
}

func (c *CompanyController) GetAll(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	companies, err := c.companyRepo.GetAll()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(companies)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *CompanyController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	company := schemas.Company{}

	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	uid, err := c.companyRepo.Create(company)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	if uid != 0 {
		company.Id = uid
	}

	b, err := json.Marshal(company)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *CompanyController) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	company := schemas.Company{}
	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.companyRepo.Update(company)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(company)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *CompanyController) Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	w.WriteHeader(http.StatusNoContent)
}

func (c *CompanyController) Get(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {

	var company schemas.Company
	companyId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	company, err = c.companyRepo.Get(companyId)

	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(company)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
