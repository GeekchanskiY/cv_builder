package router

import (
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

const vacancyRoutePrefix = "/vacancies"

func CreateVacancyRoutes(router *httprouter.Router, controller *controllers.VacancyController) {
	router.Handle(http.MethodGet, vacancyRoutePrefix, Wrapper(controller.GetAll))
	router.Handle(http.MethodPost, vacancyRoutePrefix, Wrapper(controller.Create))
	router.Handle(http.MethodPut, vacancyRoutePrefix, Wrapper(controller.Update))
	router.Handle(http.MethodDelete, vacancyRoutePrefix, Wrapper(controller.Delete))
	router.Handle(http.MethodGet, vacancyRoutePrefix+"/:id", Wrapper(controller.Get))

}
