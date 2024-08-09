package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/solndev/cv_builder/internal/controllers"
)

const cvRoutePrefix = "/cvs"
const cvProjectRoutePrefix = cvRoutePrefix + "/projects"
const cvServiceRoutePrefix = cvRoutePrefix + "/services"
const cvServiceResponsibilityRoutePrefix = cvRoutePrefix + "/service-responsibility"

func CreateCVRoutes(router *httprouter.Router, controller *controllers.CVController) {
	// CV Projects
	router.Handle(http.MethodGet, cvProjectRoutePrefix+"/:id", Wrapper(controller.GetProjects))
	router.Handle(http.MethodPost, cvProjectRoutePrefix, Wrapper(controller.CreateProjects))
	router.Handle(http.MethodPut, cvProjectRoutePrefix, Wrapper(controller.UpdateProject))
	router.Handle(http.MethodDelete, cvProjectRoutePrefix, Wrapper(controller.DeleteProject))

	// CV Services
	router.Handle(http.MethodGet, cvServiceRoutePrefix+"/:id", Wrapper(controller.GetCVService))
	router.Handle(http.MethodPost, cvServiceRoutePrefix, Wrapper(controller.CreateCVService))
	router.Handle(http.MethodPut, cvServiceRoutePrefix, Wrapper(controller.UpdateCVService))
	router.Handle(http.MethodDelete, cvServiceRoutePrefix, Wrapper(controller.DeleteCVService))

	// CV Service Responsibilities
	router.Handle(http.MethodGet, cvServiceResponsibilityRoutePrefix+"/:id", Wrapper(controller.GetCVServiceResponsibility))
	router.Handle(http.MethodPost, cvServiceResponsibilityRoutePrefix, Wrapper(controller.CreateCVServiceResponsibility))
	router.Handle(http.MethodPut, cvServiceResponsibilityRoutePrefix, Wrapper(controller.UpdateCVServiceResponsibility))
	router.Handle(http.MethodDelete, cvServiceResponsibilityRoutePrefix, Wrapper(controller.DeleteCVServiceResponsibility))

	// CV
	router.Handle(http.MethodGet, cvRoutePrefix, Wrapper(controller.GetAll))
	router.Handle(http.MethodPost, cvRoutePrefix, Wrapper(controller.Create))
	router.Handle(http.MethodPut, cvRoutePrefix, Wrapper(controller.Update))
	router.Handle(http.MethodDelete, cvRoutePrefix, Wrapper(controller.Delete))
	router.Handle(http.MethodGet, cvRoutePrefix+"/cv/:id", Wrapper(controller.Get))

}
