package main

import (
	"github.com/AranScope/me/service.api.monzo/handlers"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", handlers.Router())
	if err != nil {
		panic(err)
	}
}
