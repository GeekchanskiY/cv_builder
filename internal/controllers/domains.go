package controllers

import (
	"encoding/json"
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

func (c *DomainController) GetDomains(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	domains, err := c.domainRepo.GetDomains()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(domains)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)

}

func (c *DomainController) CreateDomain(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	domain := schemas.Domain{}

	err := json.NewDecoder(r.Body).Decode(&domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	uid, err := c.domainRepo.CreateDomain(domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if uid != 0 {
		domain.Id = uid
	}

	w.WriteHeader(http.StatusCreated)
	b, err := json.Marshal(domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(b)

}

func (c *DomainController) UpdateDomain(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	domain := schemas.Domain{}
	err := json.NewDecoder(r.Body).Decode(&domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.domainRepo.UpdateDomain(domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	b, err := json.Marshal(domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (c *DomainController) DeleteDomain(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	domain := schemas.Domain{}
	err := json.NewDecoder(r.Body).Decode(&domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.domainRepo.DeleteDomain(domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	b, err := json.Marshal(domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (c *DomainController) GetDomain(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var domain schemas.Domain
	domain_id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid domain id"))
		return
	}

	domain, err = c.domainRepo.GetDomainById(int(domain_id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid domain id"))
		return
	}

	b, err := json.Marshal(domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}
