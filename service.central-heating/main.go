package main

import (
	"github.com/AranScope/me/common/service"
	"github.com/AranScope/me/service.central-heating/controllers"
	"github.com/AranScope/me/service.central-heating/handlers"
	"log"
)

func main() {
	log.Println("starting service")

	controllers.Start()
	service.
		Init("service.central-heating").
		WithRouter(8081, handlers.Router()).
		Start()
}
