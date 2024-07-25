package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/GeekchanskiY/cv_builder/internal/repository"
	"github.com/GeekchanskiY/cv_builder/internal/schemas"
	"github.com/GeekchanskiY/cv_builder/internal/utils"
	"github.com/julienschmidt/httprouter"
	"log"
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

	employees, err := c.employeeRepo.GetAll()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	projectDomains, err := c.projectRepo.GetAllDomainsReadable()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	responsibilities, err := c.responsibilityRepo.GetAllReadable()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	responsibilityConflicts, err := c.responsibilityRepo.GetAllConflictsReadable()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	responsibilitySynonyms, err := c.responsibilityRepo.GetAllSynonymsReadable()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	skills, err := c.skillRepo.GetAllReadable()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	skillDomains, err := c.skillRepo.GetAllDomainsReadable()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	skillConflicts, err := c.skillRepo.GetAllConflictsReadable()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	vacancies, err := c.vacancyRepo.GetAllReadable()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	vacancySkills, err := c.vacancyRepo.GetAllSkillsReadable()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	vacancyDomains, err := c.vacancyRepo.GetAllDomainsReadable()
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	var data = schemas.FullDatabaseData{
		Projects:                projects,
		Domains:                 domains,
		Companies:               companies,
		Employees:               employees,
		ProjectDomains:          projectDomains,
		Responsibilities:        responsibilities,
		ResponsibilitySynonyms:  responsibilitySynonyms,
		ResponsibilityConflicts: responsibilityConflicts,
		Skills:                  skills,
		SkillDomains:            skillDomains,
		SkillConflicts:          skillConflicts,
		Vacancies:               vacancies,
		VacancyDomains:          vacancyDomains,
		VacancySkills:           vacancySkills,
	}

	b, err := json.Marshal(data)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(b); err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}

func (c *UtilsController) ImportJSON(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// No point in importing new CVs. Firstly - unsafe. Secondly - it's easier to recreate them.
	data := schemas.FullDatabaseData{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	// Used to count new items
	var created bool
	var createdItems int

	for _, company := range data.Companies {
		created, err = c.companyRepo.CreateIfNotExists(company)

		if err != nil {
			log.Println(err)
			continue
		}

		if created {
			createdItems++
		}
	}

	for _, domain := range data.Domains {
		created, err = c.domainRepo.CreateIfNotExists(domain)
		if err != nil {
			log.Println(err)
			continue
		}
		if created {
			createdItems++
		}
	}

	for _, employee := range data.Employees {
		created, err = c.employeeRepo.CreateIfNotExists(employee)

		if err != nil {
			log.Println(err)
			continue
		}

		if created {
			createdItems++
		}
	}

	for _, skill := range data.Skills {
		created, err = c.skillRepo.CreateIfNotExists(skill)

		if err != nil {
			log.Println(err)
			continue
		}
		if created {
			createdItems++
		}
	}

	for _, skillConflict := range data.SkillConflicts {
		created, err := c.skillRepo.CreateConflictIfNotExists(skillConflict)
		if err != nil {
			log.Println(err)
			continue
		}
		if created {
			createdItems++
		}
	}

	for _, project := range data.Projects {
		created, err = c.projectRepo.CreateIfNotExists(project)

		if err != nil {
			log.Println(err)
			continue
		}

		if created {
			createdItems++
		}
	}

	for _, projectDomain := range data.ProjectDomains {
		created, err = c.projectRepo.CreateDomainsIfNotExists(projectDomain)

		if err != nil {
			log.Println(err)
			continue
		}

		if created {
			createdItems++
		}
	}

	//for _, responsibility := range data.Responsibilities {
	//	created, err := c.responsibilityRepo.CreateIfNotExists(responsibility)
	//}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(fmt.Sprintf("Import completed. New items: %d", createdItems)))
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}
