package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Router() http.Handler {
	router := httprouter.New()
	router.GET("/devices", handleGetDevices)
	router.GET("/device/:id", handleGetDevice)
	return router
}
