package controllers

import (
	"encoding/json"
	"github.com/GeekchanskiY/cv_builder/internal/repository"
	"github.com/GeekchanskiY/cv_builder/internal/schemas/requests"
	"github.com/GeekchanskiY/cv_builder/internal/usecases"
	"github.com/GeekchanskiY/cv_builder/internal/utils"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

type CVBuilderController struct {
	projectRepo        *repository.ProjectRepository
	domainRepo         *repository.DomainRepository
	cvRepo             *repository.CVRepository
	employeeRepo       *repository.EmployeeRepository
	companyRepo        *repository.CompanyRepository
	responsibilityRepo *repository.ResponsibilityRepository
	skillRepo          *repository.SkillRepository
	vacancyRepo        *repository.VacanciesRepository
}

func CreateCVBuilderController(
	projectRepo *repository.ProjectRepository,
	domainRepo *repository.DomainRepository,
	cvRepo *repository.CVRepository,
	employeeRepo *repository.EmployeeRepository,
	companyRepo *repository.CompanyRepository,
	responsibilityRepo *repository.ResponsibilityRepository,
	skillRepo *repository.SkillRepository,
	vacancyRepo *repository.VacanciesRepository,
) *CVBuilderController {
	return &CVBuilderController{
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

func (c *CVBuilderController) Build(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var err error
	var requestData requests.BuildRequest

	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		utils.HandleInternalError(w, err)
		return
	}

	cvChan := make(chan int)

	// Running building a CV
	go usecases.BuildCV(requestData.EmployeeID, requestData.VacancyID, cvChan)
	newId := <-cvChan
	log.Println(requestData)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("build queued, new CV ID: " + strconv.Itoa(newId)))
	if err != nil {
		log.Println("Error writing response")
	}
}
