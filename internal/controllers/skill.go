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

type SkillController struct {
	skillRepo *repository.SkillRepository
}

func CreateSkillController(repo *repository.SkillRepository) *SkillController {
	return &SkillController{
		skillRepo: repo,
	}
}

func (c *SkillController) GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	skills, err := c.skillRepo.GetAll()
	if err != nil {
		log.Println("error getting skills from db")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(skills)
	if err != nil {
		log.Println("error marshallong skills")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)

}

func (c *SkillController) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	skill := schemas.Skill{}

	err := json.NewDecoder(r.Body).Decode(&skill)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	uid, err := c.skillRepo.Create(skill)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if uid != 0 {
		skill.Id = uid
	}

	w.WriteHeader(http.StatusCreated)
	b, err := json.Marshal(skill)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(b)

}

func (c *SkillController) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	skill := schemas.Skill{}
	err := json.NewDecoder(r.Body).Decode(&skill)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.skillRepo.Update(skill)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	b, err := json.Marshal(skill)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (c *SkillController) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	skill := schemas.Skill{}
	err := json.NewDecoder(r.Body).Decode(&skill)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.skillRepo.Delete(skill)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	b, err := json.Marshal(skill)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (c *SkillController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var skill schemas.Skill
	skill_id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid domain id"))
		return
	}

	skill, err = c.skillRepo.Get(int(skill_id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid domain id"))
		return
	}

	b, err := json.Marshal(skill)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}
