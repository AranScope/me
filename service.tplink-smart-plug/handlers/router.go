package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Router() http.Handler {
	router := httprouter.New()
	router.PATCH("/plug/:ip", loggingMiddleware(handlePatchPlug))
	router.GET("/plug/:ip", loggingMiddleware(handleGetPlug))
	return router
}
