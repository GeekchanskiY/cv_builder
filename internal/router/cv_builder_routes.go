package router

import (
	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const cvBuilderRoutePrefix = "/cv_builder"

func CreateCvBuilderRoutes(router *httprouter.Router, controller *controllers.CVBuilderController) {

	router.Handle(http.MethodPost, cvBuilderRoutePrefix+"/build", Wrapper(controller.Build))

}
