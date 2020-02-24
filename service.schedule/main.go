package main

import (
	"github.com/AranScope/me/common/service"
	"github.com/AranScope/me/service.schedule/controllers"
	"github.com/AranScope/me/service.schedule/handlers"
)

func main() {
	controllers.Init()
	service.
		Init("service.schedule").
		WithRouter(80, handlers.Router()).
		Start()
}
