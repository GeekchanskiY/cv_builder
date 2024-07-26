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
	"time"
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
	// This endpoint creates a build request and returns it's ID
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

	select {
	// returning CV build request ID to be able to track status of the building in future
	case cv := <-cvChan:
		log.Println("New CV id:", cv)
		w.Header().Set("Content-Type", "application/json")

		// BuildCV returns 0 if there's an error in building request
		if cv == 0 {
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte("Bad request"))
			if err != nil {
				log.Println("Error writing response")
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("build queued, new CV build status id: " + strconv.Itoa(cv)))
		if err != nil {
			log.Println("Error writing response")
		}

	// Timeout to avoid unexpected crash in buildCV method
	case <-time.After(10 * time.Second):
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("Error creating CV"))
		if err != nil {
			log.Println("Error writing response")
		}

	}

}
