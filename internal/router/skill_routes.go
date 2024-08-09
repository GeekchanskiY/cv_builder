package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/solndev/cv_builder/internal/controllers"
)

const skillRoutePrefix = "/skills"
const skillConflictRoutePrefix = skillRoutePrefix + "/conflicts"
const skillDomainRoutePrefix = skillRoutePrefix + "/domains"

func CreateSkillRoutes(router *httprouter.Router, skillController *controllers.SkillController) {
	// Conflicts
	router.Handle(http.MethodGet, skillConflictRoutePrefix+"/:id", Wrapper(skillController.GetConflicts))
	router.Handle(http.MethodPost, skillConflictRoutePrefix, Wrapper(skillController.CreateConflict))
	router.Handle(http.MethodPut, skillConflictRoutePrefix, Wrapper(skillController.UpdateConflict))
	router.Handle(http.MethodDelete, skillConflictRoutePrefix, Wrapper(skillController.DeleteConflict))

	// Domains
	router.Handle(http.MethodGet, skillDomainRoutePrefix+"/:id", Wrapper(skillController.GetDomains))
	router.Handle(http.MethodPost, skillDomainRoutePrefix, Wrapper(skillController.CreateDomains))
	router.Handle(http.MethodPut, skillDomainRoutePrefix, Wrapper(skillController.UpdateDomain))
	router.Handle(http.MethodDelete, skillDomainRoutePrefix, Wrapper(skillController.DeleteDomain))

	// Skills
	router.Handle(http.MethodGet, skillRoutePrefix, Wrapper(skillController.GetAll))
	router.Handle(http.MethodPost, skillRoutePrefix, Wrapper(skillController.Create))
	router.Handle(http.MethodPut, skillRoutePrefix, Wrapper(skillController.Update))
	router.Handle(http.MethodDelete, skillRoutePrefix, Wrapper(skillController.Delete))
	router.Handle(http.MethodGet, skillRoutePrefix+"/skill/:id", Wrapper(skillController.Get))

}
