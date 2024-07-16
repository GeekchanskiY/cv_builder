package router

import (
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

const cvRoutePrefix = "/cvs"
const cvProjectRoutePrefix = cvRoutePrefix + "/projects"
const cvProjectResponsibilitiesRoutePrefix = cvRoutePrefix + "/project_responsibilities"

func CreateCVRoutes(router *httprouter.Router, controller *controllers.CVController) {
	// Conflicts
	router.Handle(http.MethodGet, cvProjectRoutePrefix+"/:id", Wrapper(controller.GetProjects))
	router.Handle(http.MethodPost, cvProjectRoutePrefix, Wrapper(controller.CreateProjects))
	router.Handle(http.MethodPut, cvProjectRoutePrefix, Wrapper(controller.UpdateProject))
	router.Handle(http.MethodDelete, cvProjectRoutePrefix, Wrapper(controller.DeleteProject))

	// Project Responsibilities
	router.Handle(http.MethodGet, cvProjectResponsibilitiesRoutePrefix+"/:id", Wrapper(controller.GetProjectResponsibilities))
	router.Handle(http.MethodPost, cvProjectResponsibilitiesRoutePrefix, Wrapper(controller.CreateProjectsResponsibilities))
	router.Handle(http.MethodPut, cvProjectResponsibilitiesRoutePrefix, Wrapper(controller.UpdateProjectResponsibility))
	router.Handle(http.MethodDelete, cvProjectResponsibilitiesRoutePrefix, Wrapper(controller.DeleteProjectResponsibility))

	// CV
	router.Handle(http.MethodGet, cvRoutePrefix, Wrapper(controller.GetAll))
	router.Handle(http.MethodPost, cvRoutePrefix, Wrapper(controller.Create))
	router.Handle(http.MethodPut, cvRoutePrefix, Wrapper(controller.Update))
	router.Handle(http.MethodDelete, cvRoutePrefix, Wrapper(controller.Delete))
	router.Handle(http.MethodGet, cvRoutePrefix+"/cv/:id", Wrapper(controller.Get))

}
