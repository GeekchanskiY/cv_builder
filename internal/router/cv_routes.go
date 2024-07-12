package router

import (
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

const cvRoutePrefix = "/cvs"

func CreateCVRoutes(router *httprouter.Router, controller *controllers.CVController) {
	router.Handle(http.MethodGet, cvRoutePrefix, Wrapper(controller.GetAll))
	router.Handle(http.MethodPost, cvRoutePrefix, Wrapper(controller.Create))
	router.Handle(http.MethodPut, cvRoutePrefix, Wrapper(controller.Update))
	router.Handle(http.MethodDelete, cvRoutePrefix, Wrapper(controller.Delete))
	router.Handle(http.MethodGet, cvRoutePrefix+"/cv/:id", Wrapper(controller.Get))

	router.Handle(http.MethodGet, cvRoutePrefix+"/skill/:id", Wrapper(controller.GetSkills))
	router.Handle(http.MethodPost, cvRoutePrefix+"/skill", Wrapper(controller.AddSkill))
	router.Handle(http.MethodDelete, cvRoutePrefix+"/skill", Wrapper(controller.DeleteSkill))

}
