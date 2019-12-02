package service_device_registry

import (
	"github.com/AranScope/me/common/service"
	"github.com/AranScope/personal-finance-center/service.device-registry/handlers"
)

func main() {
	service.
		Init("service.device-registry").
		WithRouter(8080, handlers.Router()).
		Start()
}