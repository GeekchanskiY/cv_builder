package router

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func loggingMiddleware(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()

		n(w, r, ps)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	}
}

func Wrapper(h httprouter.Handle) httprouter.Handle {
	return loggingMiddleware(h)
}
