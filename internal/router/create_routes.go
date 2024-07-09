package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("Hello, World!"))
}

type Route struct {
	Method     string
	Handler    func(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Route      string
	Attributes map[string]string
}

func CreateRoutes() *httprouter.Router {
	router := httprouter.New()

	router.Handle(http.MethodGet, "/", Wrapper(handler))

	return router
}
