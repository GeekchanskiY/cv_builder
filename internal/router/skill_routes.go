package router

import (
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

const skillRoutePrefix = "/skills"
const skillConflictRoutePrefix = skillRoutePrefix + "/conflicts"

func CreateSkillRoutes(router *httprouter.Router, skillController *controllers.SkillController) {
	// Conflicts
	router.Handle(http.MethodGet, skillConflictRoutePrefix+"/:id", Wrapper(skillController.GetConflicts))
	router.Handle(http.MethodPost, skillConflictRoutePrefix, Wrapper(skillController.CreateConflict))
	router.Handle(http.MethodPut, skillConflictRoutePrefix, Wrapper(skillController.UpdateConflict))
	router.Handle(http.MethodDelete, skillConflictRoutePrefix, Wrapper(skillController.DeleteConflict))

	// Skills
	router.Handle(http.MethodGet, skillRoutePrefix, Wrapper(skillController.GetAll))
	router.Handle(http.MethodPost, skillRoutePrefix, Wrapper(skillController.Create))
	router.Handle(http.MethodPut, skillRoutePrefix, Wrapper(skillController.Update))
	router.Handle(http.MethodDelete, skillRoutePrefix, Wrapper(skillController.Delete))
	router.Handle(http.MethodGet, skillRoutePrefix+"/skill/:id", Wrapper(skillController.Get))

	router.Handle(http.MethodGet, skillRoutePrefix+"/vacancy/:id", Wrapper(skillController.GetByVacancyId))

}
