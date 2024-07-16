package router

import (
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

const responsibilityRoutePrefix = "/responsibilities"

func CreateResponsibilityRoutes(router *httprouter.Router, controller *controllers.ResponsibilityController) {
	router.Handle(http.MethodGet, responsibilityRoutePrefix, Wrapper(controller.GetAll))
	router.Handle(http.MethodPost, responsibilityRoutePrefix, Wrapper(controller.Create))
	router.Handle(http.MethodPut, responsibilityRoutePrefix, Wrapper(controller.Update))
	router.Handle(http.MethodDelete, responsibilityRoutePrefix, Wrapper(controller.Delete))
	router.Handle(http.MethodGet, responsibilityRoutePrefix+"/:id", Wrapper(controller.Get))

}
