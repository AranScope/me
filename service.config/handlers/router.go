package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Router() http.Handler {
	router := httprouter.New()
	router.PUT("/config", loggingMiddleware(handleGetConfig))
	return router
}
