package main

import (
	"github.com/AranScope/me/common/service"
	"github.com/AranScope/me/service.config/handlers"
)

func main() {
	service.
		Init("service.config").
		WithRouter(8082, handlers.Router()).
		Start()
}
