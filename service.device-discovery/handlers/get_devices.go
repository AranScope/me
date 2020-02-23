package handlers

import (
	"encoding/json"
	"github.com/AranScope/me/service.device-discovery/controller"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func handleGetDevices(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var devices []controller.Device
	for _, device := range controller.DeviceRegistry {
		devices = append(devices, device)
	}

	bytes, err := json.Marshal(devices)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(bytes)
}
