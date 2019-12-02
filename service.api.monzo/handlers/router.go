package handlers

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Router() http.Handler {
	router := httprouter.New()
	router.GET("/accounts", loggingMiddleware(handleGetAccounts))
	router.GET("/balance/:account_id", loggingMiddleware(handleGetBalance))
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)
		_, _ = w.Write([]byte("Couldn't find it"))
	})

	return router
}
