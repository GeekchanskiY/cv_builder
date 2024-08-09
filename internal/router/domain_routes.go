package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/solndev/cv_builder/internal/controllers"
)

const domainRoutePrefix = "/domains"

func CreateDomainRoutes(router *httprouter.Router, domainController *controllers.DomainController) {
	router.Handle(http.MethodGet, domainRoutePrefix, Wrapper(domainController.GetAll))
	router.Handle(http.MethodPost, domainRoutePrefix, Wrapper(domainController.Create))
	router.Handle(http.MethodPut, domainRoutePrefix, Wrapper(domainController.Update))
	router.Handle(http.MethodDelete, domainRoutePrefix, Wrapper(domainController.Delete))
	router.Handle(http.MethodGet, domainRoutePrefix+"/:id", Wrapper(domainController.Get))

}
