package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/solndev/cv_builder/internal/controllers"
)

const utilRoutePrefix = "/utils"

func CreateUtilRoutes(router *httprouter.Router, controller *controllers.UtilsController) {
	router.Handle(http.MethodGet, utilRoutePrefix+"/export", Wrapper(controller.ExportJSON))
	router.Handle(http.MethodPost, utilRoutePrefix+"/import", Wrapper(controller.ImportJSON))

}
