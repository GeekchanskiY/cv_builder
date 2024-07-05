package router

import (
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	"github.com/julienschmidt/httprouter"
	"github.com/zc2638/swag"
	"github.com/zc2638/swag/endpoint"
	"github.com/zc2638/swag/types"
)

const employeeRoutePrefix = "/employees"

func CreateEmployeeRoutes(router *httprouter.Router, employeeController *controllers.EmployeeController, api *swag.API) {
	routes := []Route{
		{
			Method:  http.MethodGet,
			Handler: employeeController.GetEmployees,
			Route:   employeeRoutePrefix,
		},
		{
			Method:  http.MethodPost,
			Handler: employeeController.CreateEmployee,
			Route:   employeeRoutePrefix,
		},
		{
			Method:  http.MethodGet,
			Handler: employeeController.GetEmployee,
			Route:   employeeRoutePrefix + "/:id",
			Attributes: map[string]string{
				"id": "integer",
			},
		},
		{
			Method:  http.MethodDelete,
			Handler: employeeController.DeleteEmployee,
			Route:   employeeRoutePrefix,
		},
		{
			Method:  http.MethodPut,
			Handler: employeeController.UpdateEmployee,
			Route:   employeeRoutePrefix,
		},
	}
	for _, route := range routes {
		router.Handle(route.Method, route.Route, Wrapper(route.Handler))
		options := []endpoint.Option{
			endpoint.Handler(route.Handler),
			endpoint.Response(http.StatusOK, "successful operation", endpoint.SchemaResponseOption(struct{}{})),
		}
		for k, v := range route.Attributes {
			options = append(options, endpoint.Path(v, types.ParameterType(k), "", true))
		}
		api.AddEndpoint(
			endpoint.New(
				route.Method,
				route.Route,
				options...),
		)
	}
}
