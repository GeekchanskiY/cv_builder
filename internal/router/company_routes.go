package router

import (
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

const companyRoutePrefix = "/companies"

func CreateCompanyRoutes(router *httprouter.Router, controller *controllers.CompanyController) {
	router.Handle(http.MethodGet, companyRoutePrefix, Wrapper(controller.GetAll))
	router.Handle(http.MethodPost, companyRoutePrefix, Wrapper(controller.Create))
	router.Handle(http.MethodPut, companyRoutePrefix, Wrapper(controller.Update))
	router.Handle(http.MethodDelete, companyRoutePrefix, Wrapper(controller.Delete))
	router.Handle(http.MethodGet, companyRoutePrefix+"/:id", Wrapper(controller.Get))

}
