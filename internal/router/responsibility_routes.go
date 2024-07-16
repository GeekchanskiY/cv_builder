package router

import (
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

const responsibilityRoutePrefix = "/responsibilities"
const responsibilityConflictRoutePrefix = responsibilityRoutePrefix + "/conflicts"

func CreateResponsibilityRoutes(router *httprouter.Router, controller *controllers.ResponsibilityController) {
	// Conflicts
	router.Handle(http.MethodGet, responsibilityConflictRoutePrefix+"/:id", Wrapper(controller.GetConflicts))
	router.Handle(http.MethodPost, responsibilityConflictRoutePrefix, Wrapper(controller.CreateConflict))
	router.Handle(http.MethodPut, responsibilityConflictRoutePrefix, Wrapper(controller.UpdateConflict))
	router.Handle(http.MethodDelete, responsibilityConflictRoutePrefix, Wrapper(controller.DeleteConflict))

	router.Handle(http.MethodGet, responsibilityRoutePrefix, Wrapper(controller.GetAll))
	router.Handle(http.MethodPost, responsibilityRoutePrefix, Wrapper(controller.Create))
	router.Handle(http.MethodPut, responsibilityRoutePrefix, Wrapper(controller.Update))
	router.Handle(http.MethodDelete, responsibilityRoutePrefix, Wrapper(controller.Delete))
	router.Handle(http.MethodGet, responsibilityRoutePrefix+"/responsibility/:id", Wrapper(controller.Get))

}
