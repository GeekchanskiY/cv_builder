package controllers

import (
	"encoding/json"
	"log"
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

func (c *VacancyController) GetSkills(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var skills []schemas.Skill
	vacancy_id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid domain id"))
		return
	}

	skills, err = c.vacancyRepo.GetSkills(int(vacancy_id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid skill id"))
		return
	}

	b, err := json.Marshal(skills)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (c *VacancyController) DeleteSkill(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vacancy_id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid vacancy id"))
		return
	}

	schema := schemas.Id{}
	err = json.NewDecoder(r.Body).Decode(&schema)
	if err != nil || schema.Id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid skill id"))
		return
	}

	err = c.vacancyRepo.DeleteSkill(int(vacancy_id), schema.Id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No such skill in this vacancy"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Skill deleted from vacancy"))
}

func (c *VacancyController) AddSkill(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vacancy_id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid domain id"))
		return
	}
	schema := schemas.Id{}
	err = json.NewDecoder(r.Body).Decode(&schema)
	if err != nil || schema.Id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid skill id"))
		return
	}

	_, err = c.vacancyRepo.AddSkill(int(vacancy_id), schema.Id)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Cant add skill to this vacancy"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Skill added to vacancy"))
}
