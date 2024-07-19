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

type ProjectController struct {
	projectRepo *repository.ProjectRepository
}

func CreateProjectController(repo *repository.ProjectRepository) *ProjectController {
	return &ProjectController{
		projectRepo: repo,
	}
}

func (c *ProjectController) GetAll(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	schemes, err := c.projectRepo.GetAll()
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

func (c *ProjectController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	schema := schemas.Project{}

	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	uid, err := c.projectRepo.Create(schema)
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

func (c *ProjectController) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	schema := schemas.Project{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.projectRepo.Update(schema)
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

func (c *ProjectController) Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.Project{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.projectRepo.Delete(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *ProjectController) Get(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {

	var schema schemas.Project
	schemaId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	schema, err = c.projectRepo.Get(schemaId)

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

func (c *ProjectController) GetDomains(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	var schemes []schemas.ProjectDomain
	schemaId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	schemes, err = c.projectRepo.GetDomains(schemaId)

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

func (c *ProjectController) CreateDomains(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.ProjectDomain{}

	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	uid, err := c.projectRepo.CreateDomains(schema)
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

func (c *ProjectController) UpdateDomain(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.ProjectDomain{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.projectRepo.UpdateDomains(schema)
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

func (c *ProjectController) DeleteDomain(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.ProjectDomain{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.projectRepo.DeleteDomains(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
