package router

import (
	"context"
	"log"
	"net/http"
	"regexp"
)

type Router struct {
	routes      []RouteEntry
	middlewares []func(http.Handler) http.Handler
}

type RouteEntry struct {
	Path    *regexp.Regexp
	Method  string
	Handler http.HandlerFunc
	// middlewares []func(http.Handler) http.Handler
}

func (rtr *Router) Route(method, path string, handlerFunc http.HandlerFunc, middlewares ...http.HandlerFunc) {
	exactPath := regexp.MustCompile("^" + path + "$")

	e := RouteEntry{
		Method:  method,
		Path:    exactPath,
		Handler: handlerFunc,
	}
	rtr.routes = append(rtr.routes, e)
}

func (ent *RouteEntry) Match(r *http.Request) map[string]string {
	match := ent.Path.FindStringSubmatch(r.URL.Path)
	if match == nil {
		return nil
	}

	// Create a map to store URL parameters in
	params := make(map[string]string)
	groupNames := ent.Path.SubexpNames()
	for i, group := range match {
		params[groupNames[i]] = group
	}

	return params
}

func (r *Router) ApplyMiddlewares(handler http.Handler) http.Handler {
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		handler = r.middlewares[i](handler)
	}
	return handler
}

func (r *Router) Use(mw func(http.Handler) http.Handler) {
	r.middlewares = append(r.middlewares, mw)
}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Error: \n %+v", r) // Log the error
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}()

	for _, e := range rtr.routes {
		params := e.Match(r)
		if params == nil {
			continue
		}
		if e.Method != r.Method {
			continue
		}

		ctx := context.WithValue(r.Context(), "params", params)
		handler := rtr.ApplyMiddlewares(e.Handler)
		handler.ServeHTTP(w, r.WithContext(ctx))
		return
	}

	http.NotFound(w, r)
}
