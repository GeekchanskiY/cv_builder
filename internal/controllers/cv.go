package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GeekchanskiY/cv_builder/internal/repository"
	"github.com/GeekchanskiY/cv_builder/internal/schemas"
	"github.com/GeekchanskiY/cv_builder/internal/utils"
	"github.com/julienschmidt/httprouter"
)

type CVController struct {
	cvRepo *repository.CVRepository
}

func CreateCVController(repo *repository.CVRepository) *CVController {
	return &CVController{
		cvRepo: repo,
	}
}

func (c *CVController) GetAll(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	schemes, err := c.cvRepo.GetAll()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
	b, err := json.Marshal(schemes)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *CVController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	schema := schemas.CV{}

	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	uid, err := c.cvRepo.Create(schema)
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

	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *CVController) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	schema := schemas.CV{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.cvRepo.Update(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *CVController) Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.CV{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.cvRepo.Delete(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *CVController) Get(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {

	var schema schemas.CV
	schemaId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	schema, err = c.cvRepo.Get(schemaId)

	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *CVController) GetProjects(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	var schemes []schemas.CVProject
	cvId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	schemes, err = c.cvRepo.GetProjects(cvId)

	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(schemes)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *CVController) CreateProjects(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.CVProject{}

	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	uid, err := c.cvRepo.CreateProject(schema)
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

	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *CVController) UpdateProject(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.CVProject{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.cvRepo.UpdateProject(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *CVController) DeleteProject(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.CVProject{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.cvRepo.DeleteProject(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *CVController) GetProjectResponsibilities(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	var schemes []schemas.CVProjectResponsibility
	cvProjectId, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	schemes, err = c.cvRepo.GetProjectsResponsibilities(cvProjectId)

	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(schemes)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *CVController) CreateProjectsResponsibilities(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.CVProjectResponsibility{}

	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	uid, err := c.cvRepo.CreateProjectResponsibility(schema)
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

	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func (c *CVController) UpdateProjectResponsibility(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.CVProjectResponsibility{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.cvRepo.UpdateProjectResponsibility(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	b, err := json.Marshal(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *CVController) DeleteProjectResponsibility(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	schema := schemas.CVProjectResponsibility{}
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	err = c.cvRepo.DeleteProjectResponsibility(schema)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
