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

type SkillController struct {
	skillRepo *repository.SkillRepository
}

func CreateSkillController(repo *repository.SkillRepository) *SkillController {
	return &SkillController{
		skillRepo: repo,
	}
}

func (c *SkillController) GetAll(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	skills, err := c.skillRepo.GetAll()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
	b, err := json.Marshal(skills)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *SkillController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	skill := schemas.Skill{}

	err := json.NewDecoder(r.Body).Decode(&skill)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	uid, err := c.skillRepo.Create(skill)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	if uid != 0 {
		skill.Id = uid
	}

	b, err := json.Marshal(skill)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *SkillController) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	skill := schemas.Skill{}
	err := json.NewDecoder(r.Body).Decode(&skill)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.skillRepo.Update(skill)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(skill)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *SkillController) Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	skill := schemas.Skill{}
	err := json.NewDecoder(r.Body).Decode(&skill)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.skillRepo.Delete(skill)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *SkillController) Get(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {

	var skill schemas.Skill
	skillId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	skill, err = c.skillRepo.Get(skillId)

	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(skill)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *SkillController) GetConflicts(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	var conflicts []schemas.SkillConflict
	skillId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	conflicts, err = c.skillRepo.GetConflicts(skillId)

	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(conflicts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *SkillController) CreateConflict(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	conflict := schemas.SkillConflict{}

	err := json.NewDecoder(r.Body).Decode(&conflict)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	uid, err := c.skillRepo.CreateConflict(conflict)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	if uid != 0 {
		conflict.Id = uid
	}

	b, err := json.Marshal(conflict)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *SkillController) UpdateConflict(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	conflict := schemas.SkillConflict{}
	err := json.NewDecoder(r.Body).Decode(&conflict)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.skillRepo.UpdateConflict(conflict)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(conflict)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *SkillController) DeleteConflict(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	conflict := schemas.SkillConflict{}
	err := json.NewDecoder(r.Body).Decode(&conflict)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.skillRepo.DeleteConflict(conflict)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *SkillController) GetDomains(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	var schemes []schemas.SkillDomain
	skillId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	schemes, err = c.skillRepo.GetDomains(skillId)

	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(schemes)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *SkillController) CreateDomains(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.SkillDomain{}

	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	uid, err := c.skillRepo.CreateDomains(schema)
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

	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *SkillController) UpdateDomain(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.SkillDomain{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.skillRepo.UpdateDomains(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *SkillController) DeleteDomain(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.SkillDomain{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.skillRepo.DeleteDomains(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
