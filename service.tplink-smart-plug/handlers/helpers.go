package handlers

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func loggingMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		log.Printf("%s %s", r.Method, r.URL)
		next(w, r, params)
	}
}

func handleError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(err.Error()))
}
