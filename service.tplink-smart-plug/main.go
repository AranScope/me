package main

import (
	handlers2 "github.com/AranScope/personal-finance-center/service.tplink-smart-plug/handlers"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":80", handlers2.Router())
	if err != nil {
		panic(err)
	}
}
