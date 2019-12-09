package main

import (
	"github.com/AranScope/me/common/service"
	"github.com/AranScope/me/local/service.tplink-smart-plug/handlers"
)

func main() {
	service.
		Init("service.tplink-smart-plug").
		WithRouter(8082, handlers.Router()).
		Start()
}
