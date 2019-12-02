package main

import (
	"github.com/AranScope/me/common/service"
	"github.com/AranScope/me/service.api.monzo/handlers"
)

func main() {
	service.
		Init("service.api.monzo").
		WithRouter(8080, handlers.Router()).
		Start()
}
