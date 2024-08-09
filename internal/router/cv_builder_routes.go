package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/solndev/cv_builder/internal/controllers"
)

const cvBuilderRoutePrefix = "/cv_builder"

func CreateCvBuilderRoutes(router *httprouter.Router, controller *controllers.CVBuilderController) {

	router.Handle(http.MethodPost, cvBuilderRoutePrefix+"/build", Wrapper(controller.Build))
	router.Handle(http.MethodGet, cvBuilderRoutePrefix+"/healthcheck", Wrapper(controller.Healthcheck))

}
