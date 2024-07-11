package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GeekchanskiY/cv_builder/internal/repository"
	"github.com/GeekchanskiY/cv_builder/internal/schemas"
	"github.com/julienschmidt/httprouter"
)

type VacancyController struct {
	vacancyRepo *repository.VacanciesRepository
}

func CreateVacancyController(repo *repository.VacanciesRepository) *VacancyController {
	return &VacancyController{
		vacancyRepo: repo,
	}
}

func (c *VacancyController) GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	domains, err := c.vacancyRepo.GetAll()
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

func (c *VacancyController) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	schema := schemas.Vacancy{}

	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	uid, err := c.vacancyRepo.Create(schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if uid != 0 {
		schema.Id = uid
	}

	w.WriteHeader(http.StatusCreated)
	b, err := json.Marshal(schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(b)

}

func (c *VacancyController) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	schema := schemas.Vacancy{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.vacancyRepo.Update(schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func (c *VacancyController) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	schema := schemas.Vacancy{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.vacancyRepo.Delete(schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write(b)
}

func (c *VacancyController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var schema schemas.Vacancy
	vacancy_id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid vacancy id"))
		return
	}

	schema, err = c.vacancyRepo.Get(int(vacancy_id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid vacancy id"))
		return
	}

	b, err := json.Marshal(schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
