package router

import (
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

const employeeRoutePrefix = "/employees"

func CreateEmployeeRoutes(router *httprouter.Router, employeeController *controllers.EmployeeController) {
	router.Handle(http.MethodGet, employeeRoutePrefix, Wrapper(employeeController.GetEmployees))
	router.Handle(http.MethodPost, employeeRoutePrefix, Wrapper(employeeController.CreateEmployee))
	router.Handle(http.MethodPut, employeeRoutePrefix, Wrapper(employeeController.UpdateEmployee))
	router.Handle(http.MethodDelete, employeeRoutePrefix, Wrapper(employeeController.DeleteEmployee))
	router.Handle(http.MethodGet, employeeRoutePrefix+"/:id", Wrapper(employeeController.GetEmployee))

}
