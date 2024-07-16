package router

import (
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

const projectRoutePrefix = "/projects"
const projectDomainRoutePrefix = projectRoutePrefix + "/domains"

func CreateProjectRoutes(router *httprouter.Router, controller *controllers.ProjectController) {
	// Domains
	router.Handle(http.MethodGet, projectDomainRoutePrefix+"/:id", Wrapper(controller.GetDomains))
	router.Handle(http.MethodPost, projectDomainRoutePrefix, Wrapper(controller.CreateDomains))
	router.Handle(http.MethodPut, projectDomainRoutePrefix, Wrapper(controller.UpdateDomain))
	router.Handle(http.MethodDelete, projectDomainRoutePrefix, Wrapper(controller.DeleteDomain))

	// Projects
	router.Handle(http.MethodGet, projectRoutePrefix, Wrapper(controller.GetAll))
	router.Handle(http.MethodPost, projectRoutePrefix, Wrapper(controller.Create))
	router.Handle(http.MethodPut, projectRoutePrefix, Wrapper(controller.Update))
	router.Handle(http.MethodDelete, projectRoutePrefix, Wrapper(controller.Delete))
	router.Handle(http.MethodGet, projectRoutePrefix+"/project/:id", Wrapper(controller.Get))

}
