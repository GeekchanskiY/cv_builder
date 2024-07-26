package usecases

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/GeekchanskiY/cv_builder/internal/repository"
	"github.com/GeekchanskiY/cv_builder/internal/schemas"
	"log"
	"time"
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

func (uc CVBuilderUseCase) BuildCV(employeeID, vacancyID int, cvChan chan int) {
	log.Println("Building CV")

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

	log.Println(fmt.Sprintf("Amount of skills to work with: %d", len(predictedSkills)))

	log.Println("CV Build Process finished")
}
