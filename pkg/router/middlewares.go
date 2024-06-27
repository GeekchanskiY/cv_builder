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

func recoverer(next httprouter.Handle) httprouter.Handle {
	fn := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		defer func() {
			if rvr := recover(); rvr != nil {
				if rvr == http.ErrAbortHandler {
					// we don't recover http.ErrAbortHandler so the response
					// to the client is aborted, this should not be logged
					panic(rvr)
				}

				log.Println("Recovered from panic")

				if r.Header.Get("Connection") != "Upgrade" {
					w.WriteHeader(http.StatusInternalServerError)
				}
			}
		}()

		next(w, r, ps)
	}

	return httprouter.Handle(fn)
}

func Wrapper(h httprouter.Handle) httprouter.Handle {
	return recoverer(loggingMiddleware(h))
}
