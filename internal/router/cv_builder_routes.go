package router

import (
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

const cvBuilderRoutePrefix = "/utils"

func CreateCvBuilderRoutes(router *httprouter.Router, controller *controllers.CVBuilderController) {
	router.Handle(http.MethodGet, utilRoutePrefix+"/build", Wrapper(controller.ExportJSON))

}
