package main

import (
	"github.com/AranScope/me/common/service"
	"github.com/AranScope/me/service.api.marcus/handlers"
)

func main() {
	service.
		Init("service.api.marcus").
		WithRouter(8080, handlers.Router()).
		Start()
}
