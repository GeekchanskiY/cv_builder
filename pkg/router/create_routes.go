package router

import "net/http"

func CreateRoutes() *Router {
	router := &Router{}
	router.Use(LoggingMiddleware)
	router.Route("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Obsidian-go API"))
	})
	return router
}
