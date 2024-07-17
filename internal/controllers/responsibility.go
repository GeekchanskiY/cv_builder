package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GeekchanskiY/cv_builder/internal/repository"
	"github.com/GeekchanskiY/cv_builder/internal/schemas"
	"github.com/julienschmidt/httprouter"
)

type ResponsibilityController struct {
	responsibilityRepo *repository.ResponsibilityRepository
}

func CreateResponsibilityController(repo *repository.ResponsibilityRepository) *ResponsibilityController {
	return &ResponsibilityController{
		responsibilityRepo: repo,
	}
}

func (c *ResponsibilityController) GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	schemes, err := c.responsibilityRepo.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(schemes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)

}

func (c *ResponsibilityController) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	schema := schemas.Responsibility{}

	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	uid, err := c.responsibilityRepo.Create(schema)
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

func (c *ResponsibilityController) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	schema := schemas.Responsibility{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.responsibilityRepo.Update(schema)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	b, err := json.Marshal(schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (c *ResponsibilityController) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	schema := schemas.Responsibility{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.responsibilityRepo.Delete(schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *ResponsibilityController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var schema schemas.Responsibility
	schema_id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid id"))
		return
	}

	schema, err = c.responsibilityRepo.Get(schema_id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (c *ResponsibilityController) GetConflicts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var conflicts []schemas.ResponsibilityConflict
	resp_id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid id"))
		return
	}

	conflicts, err = c.responsibilityRepo.GetConflicts(resp_id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid id"))
		return
	}

	b, err := json.Marshal(conflicts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (c *ResponsibilityController) CreateConflict(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	conflict := schemas.ResponsibilityConflict{}

	err := json.NewDecoder(r.Body).Decode(&conflict)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	uid, err := c.responsibilityRepo.CreateConflict(conflict)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if uid != 0 {
		conflict.Id = uid
	}

	b, err := json.Marshal(conflict)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func (c *ResponsibilityController) UpdateConflict(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	conflict := schemas.ResponsibilityConflict{}
	err := json.NewDecoder(r.Body).Decode(&conflict)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.responsibilityRepo.UpdateConflict(conflict)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(conflict)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func (c *ResponsibilityController) DeleteConflict(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	conflict := schemas.ResponsibilityConflict{}
	err := json.NewDecoder(r.Body).Decode(&conflict)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.responsibilityRepo.DeleteConflict(conflict)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *ResponsibilityController) GetSynonyms(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var schemes []schemas.ResponsibilitySynonym
	resp_id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid id"))
		return
	}

	schemes, err = c.responsibilityRepo.GetSynonyms(resp_id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid id"))
		return
	}

	b, err := json.Marshal(schemes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (c *ResponsibilityController) CreateSynonym(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	schema := schemas.ResponsibilitySynonym{}

	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	uid, err := c.responsibilityRepo.CreateSynonym(schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if uid != 0 {
		schema.Id = uid
	}

	b, err := json.Marshal(schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func (c *ResponsibilityController) UpdateSynonym(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	schema := schemas.ResponsibilitySynonym{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.responsibilityRepo.UpdateSynonym(schema)
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

func (c *ResponsibilityController) DeleteSynonym(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	schema := schemas.ResponsibilitySynonym{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.responsibilityRepo.DeleteSynonym(schema)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}