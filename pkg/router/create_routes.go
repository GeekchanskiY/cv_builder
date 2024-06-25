package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("Hello, World!"))
}

func CreateRoutes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", Wrapper(handler))
	return router
}
