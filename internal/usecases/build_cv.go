package usecases

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/solndev/cv_builder/internal/repository"
	"github.com/solndev/cv_builder/internal/schemas"
)

const (
	// Experience constant
	juniorExperienceYears = 0
	middleExperienceYears = 3
	seniorExperienceYears = 5
	leadExperienceYears   = 7

	// Responsibilities constants
	leadMaxResponsibilityAmount = 24
	leadMinResponsibilityAmount = 20

	seniorMaxResponsibilityAmount = 20
	seniorMinResponsibilityAmount = 16

	middleMaxResponsibilityAmount = 16
	middleMinResponsibilityAmount = 12

	juniorMaxResponsibilityAmount = 16
	juniorMinResponsibilityAmount = 10

	// Minimal requirements
	companiesRequired        = 1
	employeesRequired        = 1
	projectsRequired         = 8
	vacanciesRequired        = 1
	projectServicesRequired  = 16 // (4 monoliths, 4 with 3 services)
	domainsRequired          = 12
	skillsRequired           = 56 //(24 * 4 = 56)
	responsibilitiesRequired = 224
)

type CVBuilderUseCase struct {
	projectRepo        *repository.ProjectRepository
	domainRepo         *repository.DomainRepository
	cvRepo             *repository.CVRepository
	employeeRepo       *repository.EmployeeRepository
	companyRepo        *repository.CompanyRepository
	responsibilityRepo *repository.ResponsibilityRepository
	skillRepo          *repository.SkillRepository
	vacancyRepo        *repository.VacanciesRepository
}

func CreateCVBuilderUseCase(
	projectRepo *repository.ProjectRepository,
	domainRepo *repository.DomainRepository,
	cvRepo *repository.CVRepository,
	employeeRepo *repository.EmployeeRepository,
	companyRepo *repository.CompanyRepository,
	responsibilityRepo *repository.ResponsibilityRepository,
	skillRepo *repository.SkillRepository,
	vacancyRepo *repository.VacanciesRepository,
) *CVBuilderUseCase {
	return &CVBuilderUseCase{
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

type CVNameHash struct {
	EmployeeId   int       `json:"employee_id"`
	VacancyId    int       `json:"vacancy_id"`
	CreationTime time.Time `json:"creation_time"`
}

func (uc CVBuilderUseCase) BuildCV(employeeID, vacancyID, microservices int, cvChan chan int) {
	log.Println("Building CV")
	log.Println(uc.CheckDataAvailability())
	vacancyData, err := uc.vacancyRepo.Get(vacancyID)
	if err != nil {
		log.Println("Error getting vacancy data")
		cvChan <- 0
		return
	}

	employeeData, err := uc.employeeRepo.Get(employeeID)
	if err != nil {
		log.Println("Error getting employee data")
		cvChan <- 0
		return
	}

	// Hash generation for CV name:
	nameStruct := CVNameHash{
		EmployeeId:   employeeID,
		VacancyId:    vacancyID,
		CreationTime: time.Now(),
	}
	data, err := json.Marshal(nameStruct)
	if err != nil {
		log.Println("Error marshalling name")
		cvChan <- 0
		return
	}

	nameBytes := sha256.Sum256(data)
	nameString := hex.EncodeToString(nameBytes[0:])

	// Generation new CV base
	newCV := schemas.CV{
		Name:       nameString,
		VacancyId:  vacancyData.Id,
		EmployeeId: employeeData.Id,
		IsReal:     false,
	}
	cvId, err := uc.cvRepo.Create(newCV)
	if err != nil {
		log.Println("Error creating CV")
		cvChan <- 0
		return
	}
	log.Println(fmt.Sprintf("New CV build status ID: %d", cvId))

	// Generating cv build status
	cvBuildStatus := schemas.CVBuildStatus{
		CVId:      cvId,
		Status:    "queued",
		Logs:      fmt.Sprintf("CV Build init: %s", newCV.Name),
		StartTime: time.Now(),
		EndTime: sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		},
	}
	cvStatusId, err := uc.cvRepo.CreateCVBuildStatus(cvBuildStatus)
	if err != nil {
		log.Println("Error creating CV build status")
		cvChan <- 0
		return
	}

	// returning cv build status id for further creation
	cvChan <- cvStatusId

	// Getting basic required data
	requiredVacancySkills, err := uc.vacancyRepo.GetSkills(vacancyData.Id)
	if err != nil {
		log.Println("Error getting required skills")
		return
	}
	log.Println(fmt.Sprintf("Required skills in vacancy: %d", len(requiredVacancySkills)))

	requiredVacancyDomains, err := uc.vacancyRepo.GetDomains(vacancyData.Id)
	if err != nil {
		log.Println("Error getting required domains")
		return
	}
	log.Println(fmt.Sprintf("Required domains in vacancy: %d", len(requiredVacancyDomains)))

	// Getting all child skills from tree and adding them to the single array
	var predictedSkills []schemas.Skill
	for _, vacancySkill := range requiredVacancySkills {
		predictedSkillsPart, err := uc.skillRepo.GetAllChildrenSkills(vacancySkill.SkillId)
		if err != nil {
			log.Println(fmt.Sprintf("Error getting children skills for skill: %d", vacancySkill.SkillId))
			continue
		}
		for _, skill := range predictedSkillsPart {
			predictedSkills = append(predictedSkills, skill)
		}
	}

	// Calculating amount of projects
	var projectsAmount int

	if vacancyData.Experience >= juniorExperienceYears {
		projectsAmount = 1
	}

	if vacancyData.Experience >= middleExperienceYears {
		projectsAmount = 2
	}

	if vacancyData.Experience >= seniorExperienceYears {
		projectsAmount = 3
	}

	if vacancyData.Experience >= leadExperienceYears {
		projectsAmount = 4
	}

	// defining amount of projects with different architectures
	var monolithProjectsAmount = projectsAmount
	var microserviceProjectsAmount = 0
	if microservices == 1 {
		microserviceProjectsAmount = projectsAmount / 2
		monolithProjectsAmount = projectsAmount - microserviceProjectsAmount
	}

	if microservices == 2 {
		if projectsAmount >= 3 {
			monolithProjectsAmount = 1
			microserviceProjectsAmount = projectsAmount - monolithProjectsAmount
		}
	}

	// searching for projects
	var projects []schemas.Project
	for i := 0; i < microserviceProjectsAmount; i++ {
		log.Println(projects)
	}

	for i := 0; i < monolithProjectsAmount; i++ {
		log.Println(projects)
	}

	log.Println(fmt.Sprintf("Amount of skills to work with: %d", len(predictedSkills)))
	var vacancyDomainIds []int
	for _, domainID := range requiredVacancyDomains {
		vacancyDomainIds = append(vacancyDomainIds, domainID.DomainId)

	}
	predictedProjects, err := uc.projectRepo.GetMicroservicesByDomains(vacancyDomainIds)
	if err != nil {
		log.Println("Error predicting projects")
		return
	}
	log.Println(predictedProjects)
	log.Println(len(predictedProjects))
	log.Println("CV Build Process finished")
}

// CheckDataAvailability is used for handler to check if minimal required amount of data presents in the database
// to create a cv for a senior developer. Remember: it can't be less, but should be more
func (uc CVBuilderUseCase) CheckDataAvailability() (available bool, err error) {
	var amount int

	// Companies

	amount, err = uc.companyRepo.Count()
	if err != nil {
		return false, err
	}
	if amount < companiesRequired {
		return false, nil
	}

	// Employees

	amount, err = uc.employeeRepo.Count()
	if err != nil {
		return false, err
	}
	if amount < employeesRequired {
		return false, nil
	}

	// Projects

	amount, err = uc.projectRepo.Count()
	if err != nil {
		return false, err
	}
	if amount < projectsRequired {
		return false, nil
	}

	// Project services

	amount, err = uc.projectRepo.CountServices()
	if err != nil {
		return false, err
	}
	if amount < projectServicesRequired {
		return false, nil
	}

	// Vacancies

	amount, err = uc.vacancyRepo.Count()
	if err != nil {
		return false, err
	}
	if amount < vacanciesRequired {
		return false, nil
	}

	// Domains

	amount, err = uc.domainRepo.Count()
	if err != nil {
		return false, err
	}
	if amount < domainsRequired {
		return false, nil
	}

	// Skills

	amount, err = uc.skillRepo.Count()
	if err != nil {
		return false, err
	}
	if amount < skillsRequired {
		return false, nil
	}

	// Responsibilities

	amount, err = uc.responsibilityRepo.Count()
	if err != nil {
		return false, err
	}
	if amount < responsibilitiesRequired {
		return false, nil
	}

	return true, nil
}
