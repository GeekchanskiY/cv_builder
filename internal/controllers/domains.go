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

func (c *DomainController) GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	domains, err := c.domainRepo.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(domains)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (c *DomainController) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	domain := schemas.Domain{}

	err := json.NewDecoder(r.Body).Decode(&domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	uid, err := c.domainRepo.Create(domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if uid != 0 {
		domain.Id = uid
	}

	b, err := json.Marshal(domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func (c *DomainController) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	domain := schemas.Domain{}
	err := json.NewDecoder(r.Body).Decode(&domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.domainRepo.Update(domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func (c *DomainController) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	domain := schemas.Domain{}
	err := json.NewDecoder(r.Body).Decode(&domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.domainRepo.Delete(domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write(b)
}

func (c *DomainController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var domain schemas.Domain
	domain_id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid domain id"))
		return
	}

	domain, err = c.domainRepo.Get(int(domain_id))

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

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
