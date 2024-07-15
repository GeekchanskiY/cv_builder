package router

import (
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

const vacancyRoutePrefix = "/vacancies"
const vacancySkillRoutePrefix = "/vacancies/skills"

func CreateVacancyRoutes(router *httprouter.Router, controller *controllers.VacancyController) {
	router.Handle(http.MethodGet, vacancyRoutePrefix, Wrapper(controller.GetAll))
	router.Handle(http.MethodPost, vacancyRoutePrefix, Wrapper(controller.Create))
	router.Handle(http.MethodPut, vacancyRoutePrefix, Wrapper(controller.Update))
	router.Handle(http.MethodDelete, vacancyRoutePrefix, Wrapper(controller.Delete))
	router.Handle(http.MethodGet, vacancyRoutePrefix+"/vacancy/:id", Wrapper(controller.Get))

	router.Handle(http.MethodGet, vacancySkillRoutePrefix+"/:id", Wrapper(controller.GetSkills))
	router.Handle(http.MethodPost, vacancySkillRoutePrefix, Wrapper(controller.AddSkill))
	router.Handle(http.MethodDelete, vacancySkillRoutePrefix, Wrapper(controller.DeleteSkill))
	router.Handle(http.MethodPut, vacancySkillRoutePrefix, Wrapper(controller.UpdateSkill))

}
