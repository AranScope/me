package main

import (
	"github.com/AranScope/me/common/service"
	"github.com/AranScope/me/service.tplink-smart-plug/handlers"
)

func main() {
	service.
		Init("service.api.vanguard-investor").
		WithRouter(8080, handlers.Router()).
		Start()
}
