package router

import (
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

const employeeRoutePrefix = "/employees"

func CreateEmployeeRoutes(router *httprouter.Router, employeeController *controllers.EmployeeController) {
	router.Handle(http.MethodGet, employeeRoutePrefix, Wrapper(employeeController.GetAll))
	router.Handle(http.MethodPost, employeeRoutePrefix, Wrapper(employeeController.Create))
	router.Handle(http.MethodPut, employeeRoutePrefix, Wrapper(employeeController.Update))
	router.Handle(http.MethodDelete, employeeRoutePrefix, Wrapper(employeeController.Delete))
	router.Handle(http.MethodGet, employeeRoutePrefix+"/:id", Wrapper(employeeController.Get))

}
