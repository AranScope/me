package handlers

import (
	"encoding/json"
	"github.com/AranScope/me/service.device-discovery/controller"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func handleGetDevices(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var devices []controller.Device
	visited := map[string]interface{}{}
	for _, device := range controller.DeviceRegistry {
		if _, ok := visited[device.MacAddr]; !ok {
			devices = append(devices, device)
			visited[device.MacAddr] = struct{}{}
		}
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
