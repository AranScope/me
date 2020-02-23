package handlers

import (
	"encoding/json"
	"github.com/AranScope/me/service.device-discovery/controller"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func handleGetDevice(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	if device, ok := controller.DeviceRegistry[id]; ok {
		bytes, err := json.Marshal(device)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(bytes)
	}

	w.WriteHeader(http.StatusNotFound)
}
