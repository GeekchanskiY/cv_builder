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
	var domains []schemas.Domain

	projects, err = c.projectRepo.GetAll()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	domains, err = c.domainRepo.GetAll()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	companies, err := c.companyRepo.GetAll()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	cvs, err := c.cvRepo.GetAll()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	cvProjects, err := c.cvRepo.GetAllProjects()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	cvProjectResponsibilities, err := c.cvRepo.GetAllProjectResponsibilities()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	employees, err := c.employeeRepo.GetAll()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	projectDomains, err := c.projectRepo.GetAllDomains()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	responsibilities, err := c.responsibilityRepo.GetAll()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	responsibilityConflicts, err := c.responsibilityRepo.GetAllConflicts()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	responsibilitySynonyms, err := c.responsibilityRepo.GetAllSynonyms()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	skills, err := c.skillRepo.GetAll()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	skillDomains, err := c.skillRepo.GetAllDomains()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	skillConflicts, err := c.skillRepo.GetAllConflicts()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	vacancies, err := c.vacancyRepo.GetAll()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	vacancySkills, err := c.vacancyRepo.GetAllSkills()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	vacancyDomains, err := c.vacancyRepo.GetAllDomains()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	var data = schemas.FullDatabaseData{
		Projects:                  projects,
		Domains:                   domains,
		Companies:                 companies,
		CVs:                       cvs,
		CVProjects:                cvProjects,
		CVProjectResponsibilities: cvProjectResponsibilities,
		Employees:                 employees,
		ProjectDomains:            projectDomains,
		Responsibilities:          responsibilities,
		ResponsibilitySynonyms:    responsibilitySynonyms,
		ResponsibilityConflicts:   responsibilityConflicts,
		Skills:                    skills,
		SkillDomains:              skillDomains,
		SkillConflicts:            skillConflicts,
		Vacancies:                 vacancies,
		VacancyDomains:            vacancyDomains,
		VacancySkills:             vacancySkills,
	}

	b, err := json.Marshal(data)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
