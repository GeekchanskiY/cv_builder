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
