package main

import (
	"github.com/AranScope/me/common/service"
	"github.com/AranScope/me/service.tplink-smart-plug/handlers"
)

func main() {
	service.
		Init("service.tplink-smart-plug").
		WithRouter(8080, handlers.Router()).
		Start()
}
