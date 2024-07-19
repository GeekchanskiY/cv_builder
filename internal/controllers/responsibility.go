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

type ResponsibilityController struct {
	responsibilityRepo *repository.ResponsibilityRepository
}

func CreateResponsibilityController(repo *repository.ResponsibilityRepository) *ResponsibilityController {
	return &ResponsibilityController{
		responsibilityRepo: repo,
	}
}

func (c *ResponsibilityController) GetAll(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	schemes, err := c.responsibilityRepo.GetAll()
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
	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *ResponsibilityController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	schema := schemas.Responsibility{}

	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	uid, err := c.responsibilityRepo.Create(schema)
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
	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *ResponsibilityController) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	schema := schemas.Responsibility{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.responsibilityRepo.Update(schema)
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
	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *ResponsibilityController) Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.Responsibility{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.responsibilityRepo.Delete(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *ResponsibilityController) Get(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {

	var schema schemas.Responsibility
	schemaId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	schema, err = c.responsibilityRepo.Get(schemaId)

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
	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *ResponsibilityController) GetConflicts(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	var conflicts []schemas.ResponsibilityConflict
	respId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	conflicts, err = c.responsibilityRepo.GetConflicts(respId)

	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(conflicts)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *ResponsibilityController) CreateConflict(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	conflict := schemas.ResponsibilityConflict{}

	err := json.NewDecoder(r.Body).Decode(&conflict)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	uid, err := c.responsibilityRepo.CreateConflict(conflict)
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

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *ResponsibilityController) UpdateConflict(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	conflict := schemas.ResponsibilityConflict{}
	err := json.NewDecoder(r.Body).Decode(&conflict)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.responsibilityRepo.UpdateConflict(conflict)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(conflict)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *ResponsibilityController) DeleteConflict(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	conflict := schemas.ResponsibilityConflict{}
	err := json.NewDecoder(r.Body).Decode(&conflict)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.responsibilityRepo.DeleteConflict(conflict)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *ResponsibilityController) GetSynonyms(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	var schemes []schemas.ResponsibilitySynonym
	respId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	schemes, err = c.responsibilityRepo.GetSynonyms(respId)

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
	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

}

func (c *ResponsibilityController) CreateSynonym(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.ResponsibilitySynonym{}

	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	uid, err := c.responsibilityRepo.CreateSynonym(schema)
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

	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *ResponsibilityController) UpdateSynonym(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.ResponsibilitySynonym{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.responsibilityRepo.UpdateSynonym(schema)
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
	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *ResponsibilityController) DeleteSynonym(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.ResponsibilitySynonym{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.responsibilityRepo.DeleteSynonym(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
