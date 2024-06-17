package router

func CreateRoutes() *Router {
	router := &Router{}
	router.Use(LoggingMiddleware)
	return router
}
