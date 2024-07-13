package router

import (
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

const projectRoutePrefix = "/projects"

func CreateProjectRoutes(router *httprouter.Router, controller *controllers.ProjectController) {
	router.Handle(http.MethodGet, projectRoutePrefix, Wrapper(controller.GetAll))
	router.Handle(http.MethodPost, projectRoutePrefix, Wrapper(controller.Create))
	router.Handle(http.MethodPut, projectRoutePrefix, Wrapper(controller.Update))
	router.Handle(http.MethodDelete, projectRoutePrefix, Wrapper(controller.Delete))
	router.Handle(http.MethodGet, projectRoutePrefix+"/:id", Wrapper(controller.Get))

}
