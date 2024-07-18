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

type DomainController struct {
	domainRepo *repository.DomainRepository
}

func CreateDomainController(repo *repository.DomainRepository) *DomainController {
	return &DomainController{
		domainRepo: repo,
	}
}

func (c *DomainController) GetAll(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	domains, err := c.domainRepo.GetAll()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
	b, err := json.Marshal(domains)
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

func (c *DomainController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	domain := schemas.Domain{}

	err := json.NewDecoder(r.Body).Decode(&domain)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	uid, err := c.domainRepo.Create(domain)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	if uid != 0 {
		domain.Id = uid
	}

	b, err := json.Marshal(domain)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *DomainController) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	domain := schemas.Domain{}
	err := json.NewDecoder(r.Body).Decode(&domain)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.domainRepo.Update(domain)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(domain)
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

func (c *DomainController) Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	domain := schemas.Domain{}
	err := json.NewDecoder(r.Body).Decode(&domain)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.domainRepo.Delete(domain)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *DomainController) Get(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {

	var domain schemas.Domain
	domainId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	domain, err = c.domainRepo.Get(domainId)

	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(domain)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
	}
	w.WriteHeader(http.StatusOK)
}
