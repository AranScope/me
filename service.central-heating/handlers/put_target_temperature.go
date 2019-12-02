package handlers

import (
	"encoding/json"
	"github.com/AranScope/me/service.central-heating/controllers"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

type SetTargetTemperatureRequest struct {
	Temperature float64 `json:"temperature"`
}

func handleSetTargetTemperature(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		handleError(err, w)
		return
	}

	body := SetTargetTemperatureRequest{}
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		handleError(err, w)
		return
	}

	controllers.TargetTemp = body.Temperature

	w.WriteHeader(200)
	_, _ = w.Write(bodyBytes)
}
