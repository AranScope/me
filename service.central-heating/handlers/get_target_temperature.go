package handlers

import (
	"encoding/json"
	"github.com/AranScope/me/service.central-heating/controllers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type GetTargetTemperatureResponse struct {
	TargetTemperature  float64 `json:"target_temperature"`
	CurrentTemperature float64 `json:"current_temperature"`
}

func handleGetTargetTemperature(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	body := GetTargetTemperatureResponse{
		TargetTemperature:  controllers.TargetTemp,
		CurrentTemperature: controllers.CurrentTemp,
	}

	js, err := json.Marshal(body)
	if err != nil {
		handleError(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	_, _ = w.Write(js)
}
