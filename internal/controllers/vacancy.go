package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/solndev/cv_builder/internal/utils"

	"github.com/julienschmidt/httprouter"
	"github.com/solndev/cv_builder/internal/repository"
	"github.com/solndev/cv_builder/internal/schemas"
)

type VacancyController struct {
	vacancyRepo *repository.VacanciesRepository
}

func CreateVacancyController(repo *repository.VacanciesRepository) *VacancyController {
	return &VacancyController{
		vacancyRepo: repo,
	}
}

func (c *VacancyController) GetAll(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	domains, err := c.vacancyRepo.GetAll()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
	b, err := json.Marshal(domains)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *VacancyController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	schema := schemas.Vacancy{}

	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	uid, err := c.vacancyRepo.Create(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	if uid != 0 {
		schema.Id = uid
	}

	b, err := json.Marshal(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *VacancyController) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	schema := schemas.Vacancy{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.vacancyRepo.Update(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *VacancyController) Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	w.WriteHeader(http.StatusNoContent)
}

func (c *VacancyController) Get(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {

	var schema schemas.Vacancy
	vacancyId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	schema, err = c.vacancyRepo.Get(vacancyId)

	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *VacancyController) GetSkills(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	var schemes []schemas.VacancySkill
	vacancyId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	schemes, err = c.vacancyRepo.GetSkills(vacancyId)

	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(schemes)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *VacancyController) DeleteSkill(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.VacancySkill{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil || schema.Id == 0 {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.vacancyRepo.DeleteSkill(schema)

	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *VacancyController) AddSkill(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.VacancySkill{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	newId, err := c.vacancyRepo.AddSkill(schema)

	if err != nil || newId == 0 {
		log.Println(err)
		utils.HandleInternalError(w, err)
		return
	}

	schema.Id = newId

	b, err := json.Marshal(schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *VacancyController) UpdateSkill(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.VacancySkill{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.vacancyRepo.UpdateSkill(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *VacancyController) GetDomains(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	var schemes []schemas.VacancyDomain
	vacancyId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	schemes, err = c.vacancyRepo.GetDomains(vacancyId)

	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(schemes)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *VacancyController) DeleteDomain(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.VacancyDomain{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil || schema.Id == 0 {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.vacancyRepo.DeleteDomain(schema)

	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *VacancyController) AddDomain(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.VacancyDomain{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	newId, err := c.vacancyRepo.AddDomain(schema)

	if err != nil || newId == 0 {
		log.Println(err)
		utils.HandleInternalError(w, err)
		return
	}

	schema.Id = newId

	b, err := json.Marshal(schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *VacancyController) UpdateDomain(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.VacancyDomain{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.vacancyRepo.UpdateDomain(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}
