package router

import (
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

const cvRoutePrefix = "/cvs"
const cvProjectRoutePrefix = cvRoutePrefix + "/projects"

func CreateCVRoutes(router *httprouter.Router, controller *controllers.CVController) {
	// Conflicts
	router.Handle(http.MethodGet, cvProjectRoutePrefix+"/:id", Wrapper(controller.GetProjects))
	router.Handle(http.MethodPost, cvProjectRoutePrefix, Wrapper(controller.CreateProjects))
	router.Handle(http.MethodPut, cvProjectRoutePrefix, Wrapper(controller.UpdateProject))
	router.Handle(http.MethodDelete, cvProjectRoutePrefix, Wrapper(controller.DeleteProject))

	// CV
	router.Handle(http.MethodGet, cvRoutePrefix, Wrapper(controller.GetAll))
	router.Handle(http.MethodPost, cvRoutePrefix, Wrapper(controller.Create))
	router.Handle(http.MethodPut, cvRoutePrefix, Wrapper(controller.Update))
	router.Handle(http.MethodDelete, cvRoutePrefix, Wrapper(controller.Delete))
	router.Handle(http.MethodGet, cvRoutePrefix+"/cv/:id", Wrapper(controller.Get))

}
