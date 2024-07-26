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

	projectServices, err := c.projectRepo.GetAllServicesReadable()
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
		ProjectServices:         projectServices,
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
	var createdItems = 0
	var createdTotal = 0

	log.Println("Import process started")

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

	log.Printf("Companies: provided %d, created: %d", len(data.Companies), createdItems)
	createdTotal += createdItems
	createdItems = 0

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

	log.Printf("Domains: provided %d, created: %d", len(data.Domains), createdItems)
	createdTotal += createdItems
	createdItems = 0

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

	log.Printf("Employees: provided %d, created: %d", len(data.Employees), createdItems)
	createdTotal += createdItems
	createdItems = 0

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

	log.Printf("Skills: provided %d, created: %d", len(data.Skills), createdItems)
	createdTotal += createdItems
	createdItems = 0

	for _, skillConflict := range data.SkillConflicts {
		created, err = c.skillRepo.CreateConflictIfNotExists(skillConflict)
		if err != nil {
			log.Println(err)
			continue
		}
		if created {
			createdItems++
		}
	}

	log.Printf("SkillConflicts: provided %d, created: %d", len(data.SkillConflicts), createdItems)
	createdTotal += createdItems
	createdItems = 0

	for _, skillDomain := range data.SkillDomains {
		created, err = c.skillRepo.CreateDomainIfNotExists(skillDomain)
		if err != nil {
			log.Println(err)
			continue
		}
		if created {
			createdItems++
		}
	}

	log.Printf("SkillDomains: provided %d, created: %d", len(data.SkillDomains), createdItems)
	createdTotal += createdItems
	createdItems = 0

	for _, responsibility := range data.Responsibilities {
		created, err = c.responsibilityRepo.CreateIfNotExists(responsibility)
		if err != nil {
			log.Println(err)
			continue
		}
		if created {
			createdItems++
		}
	}

	log.Printf("Responsibilities: provided %d, created: %d", len(data.Responsibilities), createdItems)
	createdTotal += createdItems
	createdItems = 0

	for _, responsibilityConflict := range data.ResponsibilityConflicts {
		created, err = c.responsibilityRepo.CreateConflictIfNotExists(responsibilityConflict)
		if err != nil {
			log.Println(err)
			continue
		}
		if created {
			createdItems++
		}
	}

	log.Printf("ResponsibilityConflicts: provided %d, created: %d", len(data.ResponsibilityConflicts), createdItems)
	createdTotal += createdItems
	createdItems = 0

	for _, responsibilitySynonym := range data.ResponsibilitySynonyms {
		created, err = c.responsibilityRepo.CreateSynonymIfNotExists(responsibilitySynonym)
		if err != nil {
			log.Println(err)
			continue
		}
		if created {
			createdItems++
		}
	}

	log.Printf("ResponsibilitySynonyms: provided %d, created: %d", len(data.ResponsibilitySynonyms), createdItems)
	createdTotal += createdItems
	createdItems = 0

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

	log.Printf("Projects: provided %d, created: %d", len(data.Projects), createdItems)
	createdTotal += createdItems
	createdItems = 0

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

	log.Printf("ProjectDomains: provided %d, created: %d", len(data.ProjectDomains), createdItems)

	createdTotal += createdItems
	createdItems = 0

	for _, projectService := range data.ProjectServices {
		created, err = c.projectRepo.CreateServiceIfNotExists(projectService)

		if err != nil {
			log.Println(err)
			continue
		}

		if created {
			createdItems++
		}
	}

	log.Printf("ProjectServices: provided %d, created: %d", len(data.ProjectServices), createdItems)

	createdTotal += createdItems
	createdItems = 0

	for _, vacancy := range data.Vacancies {
		created, err = c.vacancyRepo.CreateIfNotExists(vacancy)

		if err != nil {
			log.Println(err)
			continue
		}

		if created {
			createdItems++
		}
	}

	log.Printf("Vacancies: provided %d, created: %d", len(data.Vacancies), createdItems)
	createdTotal += createdItems
	createdItems = 0

	for _, vacancyDomain := range data.VacancyDomains {
		created, err = c.vacancyRepo.CreateDomainIfNotExists(vacancyDomain)

		if err != nil {
			log.Println(err)
			continue
		}

		if created {
			createdItems++
		}
	}

	log.Printf("VacancyDomains: provided %d, created: %d", len(data.VacancyDomains), createdItems)
	createdTotal += createdItems
	createdItems = 0

	for _, vacancySkill := range data.VacancySkills {
		created, err = c.vacancyRepo.CreateSkillIfNotExists(vacancySkill)

		if err != nil {
			log.Println(err)
			continue
		}

		if created {
			createdItems++
		}
	}

	log.Printf("VacancySkills: provided %d, created: %d", len(data.VacancySkills), createdItems)
	createdTotal += createdItems
	createdItems = 0

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(fmt.Sprintf("Import completed. New items: %d", createdTotal)))
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}
}
