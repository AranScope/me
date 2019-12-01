package service_device_registry

import (
	"github.com/AranScope/personal-finance-center/service.device-registry/handlers"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":80", handlers.Router())
	if err != nil {
		panic(err)
	}
}