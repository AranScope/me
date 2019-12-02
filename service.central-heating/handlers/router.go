package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Router() http.Handler {
	router := httprouter.New()
	router.PUT("/temperature", loggingMiddleware(handleSetTargetTemperature))
	router.GET("/temperature", loggingMiddleware(handleGetTargetTemperature))
	return router
}
