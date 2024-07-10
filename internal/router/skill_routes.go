package router

import (
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

const skillRoutePrefix = "/skills"

func CreateSkillRoutes(router *httprouter.Router, skillController *controllers.SkillController) {
	router.Handle(http.MethodGet, skillRoutePrefix, Wrapper(skillController.GetAll))
	router.Handle(http.MethodPost, skillRoutePrefix, Wrapper(skillController.Create))
	router.Handle(http.MethodPut, skillRoutePrefix, Wrapper(skillController.Update))
	router.Handle(http.MethodDelete, skillRoutePrefix, Wrapper(skillController.Delete))
	router.Handle(http.MethodGet, skillRoutePrefix+"/:id", Wrapper(skillController.Get))
}
