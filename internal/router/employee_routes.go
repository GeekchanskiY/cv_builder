package router

import (
	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

const employeeRoutePrefix = "/employees"

func CreateEmployeeRoutes(router *httprouter.Router, employeeController *controllers.EmployeeController) {
	router.GET(employeeRoutePrefix, Wrapper(employeeController.GetEmployees))
	router.PUT(employeeRoutePrefix, Wrapper(employeeController.UpdateEmployee))
	router.POST(employeeRoutePrefix, Wrapper(employeeController.CreateEmployee))
	router.GET(employeeRoutePrefix+"/:id", Wrapper(employeeController.GetEmployee))
	router.DELETE(employeeRoutePrefix, Wrapper(employeeController.DeleteEmployee))
}
