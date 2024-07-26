package controllers

import (
	"encoding/json"
	"github.com/GeekchanskiY/cv_builder/internal/schemas/requests"
	"github.com/GeekchanskiY/cv_builder/internal/usecases"
	"github.com/GeekchanskiY/cv_builder/internal/utils"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

type CVBuilderController struct {
	useCase *usecases.CVBuilderUseCase
}

func CreateCVBuilderController(
	useCase *usecases.CVBuilderUseCase,
) *CVBuilderController {
	return &CVBuilderController{
		useCase: useCase,
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
	go c.useCase.BuildCV(requestData.EmployeeID, requestData.VacancyID, cvChan)
	newId := <-cvChan
	log.Println(requestData)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("build queued, new CV ID: " + strconv.Itoa(newId)))
	if err != nil {
		log.Println("Error writing response")
	}
}
