package handlers

import (
	"encoding/json"
	"github.com/AranScope/me/service.config/domain"
	"github.com/AranScope/me/service.config/types"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func handleGetConfig(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	body := &types.GetConfigRequest{}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(err, w)
		return
	}

	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		handleError(err, w)
		return
	}

	cfg, err := domain.ReadConfig(body.Path)
	if err != nil {
		handleError(err, w)
		return
	}

	cfgBytes, err := json.Marshal(cfg)
	if err != nil {
		handleError(err, w)
		return
	}

	rsp := &types.GetConfigResponse{
		Path: body.Path,
		Body: string(cfgBytes),
	}

	rspBytes, err := json.Marshal(rsp)
	if err != nil {
		handleError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, _ = w.Write(rspBytes)
}
