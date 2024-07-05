package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/zc2638/swag"
	"github.com/zc2638/swag/endpoint"
)

func handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("Hello, World! 2"))
}

func handler3(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("Hello, World!"))
}

type Route struct {
	Method     string
	Handler    func(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Route      string
	Attributes map[string]string
}

func CreateRoutes(api *swag.API) *httprouter.Router {
	router := httprouter.New()

	routes := []Route{
		{
			Method:  http.MethodGet,
			Handler: handler,
			Route:   "/",
		},
		{
			Method:  http.MethodGet,
			Handler: handler3,
			Route:   "/3",
		},
	}

	for _, route := range routes {
		router.Handle(route.Method, route.Route, Wrapper(route.Handler))
		api.AddEndpoint(
			endpoint.New(route.Method, route.Route, endpoint.Handler(route.Handler), endpoint.Response(http.StatusOK, "successful operation", endpoint.SchemaResponseOption(struct{}{}))),
		)
	}
	// router.GET("/", Wrapper(handler))
	// router.GET("/3", Wrapper(handler3))
	router.Handler(http.MethodGet, "/swagger/json", api.Handler())
	router.Handler(http.MethodGet, "/swagger/ui/*any", swag.UIHandler("/swagger/ui", "/swagger/json", true))
	// api.AddEndpoint(
	// 	endpoint.New(http.MethodGet, "/", endpoint.Handler(handler), endpoint.Response(http.StatusOK, "successful operation", endpoint.SchemaResponseOption(struct{}{}))),
	// )
	return router
}
