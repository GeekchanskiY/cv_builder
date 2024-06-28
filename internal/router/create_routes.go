package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	panic("ASDASD")
	w.Write([]byte("Hello, World!"))
}

func handler3(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("Hello, World!"))
}

func CreateRoutes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", Wrapper(handler))
	router.GET("/3", Wrapper(handler3))
	return router
}
