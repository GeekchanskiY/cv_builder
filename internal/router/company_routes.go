package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/solndev/cv_builder/internal/controllers"
)

const companyRoutePrefix = "/companies"

func CreateCompanyRoutes(router *httprouter.Router, controller *controllers.CompanyController) {
	router.Handle(http.MethodGet, companyRoutePrefix, Wrapper(controller.GetAll))
	router.Handle(http.MethodPost, companyRoutePrefix, Wrapper(controller.Create))
	router.Handle(http.MethodPut, companyRoutePrefix, Wrapper(controller.Update))
	router.Handle(http.MethodDelete, companyRoutePrefix, Wrapper(controller.Delete))
	router.Handle(http.MethodGet, companyRoutePrefix+"/:id", Wrapper(controller.Get))

}
