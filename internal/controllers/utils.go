package controllers

import (
	"encoding/json"
	"github.com/GeekchanskiY/cv_builder/internal/repository"
	"github.com/GeekchanskiY/cv_builder/internal/schemas"
	"github.com/GeekchanskiY/cv_builder/internal/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UtilsController struct {
	projectRepo        *repository.ProjectRepository
	domainRepo         *repository.DomainRepository
	cvRepo             *repository.CVRepository
	employeeRepo       *repository.EmployeeRepository
	companyRepo        *repository.CompanyRepository
	responsibilityRepo *repository.ResponsibilityRepository
	skillRepo          *repository.SkillRepository
	vacancyRepo        *repository.VacanciesRepository
}

func CreateUtilsController(
	projectRepo *repository.ProjectRepository,
	domainRepo *repository.DomainRepository,
	cvRepo *repository.CVRepository,
	employeeRepo *repository.EmployeeRepository,
	companyRepo *repository.CompanyRepository,
	responsibilityRepo *repository.ResponsibilityRepository,
	skillRepo *repository.SkillRepository,
	vacancyRepo *repository.VacanciesRepository,
) *UtilsController {
	return &UtilsController{
		projectRepo:        projectRepo,
		domainRepo:         domainRepo,
		cvRepo:             cvRepo,
		employeeRepo:       employeeRepo,
		companyRepo:        companyRepo,
		responsibilityRepo: responsibilityRepo,
		skillRepo:          skillRepo,
		vacancyRepo:        vacancyRepo,
	}
}

func (c *UtilsController) ExportJSON(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	var err error
	var projects []schemas.Project
	projects, err = c.projectRepo.GetAll()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	var data = schemas.FullDatabaseData{
		Projects: projects,
	}

	b, err := json.Marshal(data)
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
