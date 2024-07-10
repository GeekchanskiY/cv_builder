package router

import (
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

const domainRoutePrefix = "/domains"

func CreateDomainRoutes(router *httprouter.Router, domainController *controllers.DomainController) {
	router.Handle(http.MethodGet, domainRoutePrefix, Wrapper(domainController.GetDomains))
	router.Handle(http.MethodPost, domainRoutePrefix, Wrapper(domainController.CreateDomain))
	router.Handle(http.MethodPut, domainRoutePrefix, Wrapper(domainController.UpdateDomain))
	router.Handle(http.MethodDelete, domainRoutePrefix, Wrapper(domainController.DeleteDomain))
	router.Handle(http.MethodGet, domainRoutePrefix+"/:id", Wrapper(domainController.GetDomain))

}
