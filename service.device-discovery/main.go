package main

import (
	"github.com/AranScope/me/common/service"
	"github.com/AranScope/me/service.device-discovery/controller"
	"github.com/AranScope/me/service.device-discovery/handlers"
)

func main() {
	controller.Init()
	service.
		Init("service.device-discovery").
		WithRouter(8087, handlers.Router()).
		Start()
}
