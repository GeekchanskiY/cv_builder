package router

import (
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

const utilRoutePrefix = "/utils"

func CreateUtilRoutes(router *httprouter.Router, controller *controllers.UtilsController) {
	router.Handle(http.MethodGet, utilRoutePrefix+"/export", Wrapper(controller.ExportJSON))

}
