package router

import (
	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	"github.com/julienschmidt/httprouter"
)

func CreateEmployeeRoutes(router *httprouter.Router, employeeController *controllers.EmployeeController) {
	router.GET("/employees/", Wrapper(employeeController.GetEmployees))
}
