package handlers

import (
	"encoding/json"
	"github.com/AranScope/hs1xxplug"
	"github.com/AranScope/me/local/service.tplink-smart-plug/types"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func handleGetPlug(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	ip := params.ByName("ip")

	plug := hs1xxplug.Hs1xxPlug{IPAddress: ip}
	res, err := plug.SystemInfo()
	if err != nil {
		handleError(err, w)
		return
	}

	resJson := make(map[string]interface{})
	err = json.Unmarshal([]byte(res), &resJson)
	if err != nil {
		handleError(err, w)
		return
	}

	rsp := tplinkResponseToResponse(resJson)

	js, err := json.Marshal(rsp)

	if err != nil {
		handleError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, _ = w.Write(js)
}

func tplinkResponseToResponse(tplinkRes map[string]interface{}) *types.GetPlugResponse {
	systemInfo := tplinkRes["system"].(map[string]interface{})["get_sysinfo"].(map[string]interface{})
	relayState := int(systemInfo["relay_state"].(float64))
	model := systemInfo["model"].(string)

	return &types.GetPlugResponse{
		State: relayStateToState(relayState),
		Model: model,
	}
}

func relayStateToState(relayState int) string {
	if relayState == 0 {
		return "off"
	}
	return "on"
}
