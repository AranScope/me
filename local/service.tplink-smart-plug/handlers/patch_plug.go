package handlers

import (
	"encoding/json"
	"errors"
	"github.com/AranScope/hs1xxplug"
	"github.com/AranScope/me/local/service.tplink-smart-plug/types"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func handlePatchPlug(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	ip := params.ByName("ip")

	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		handleError(err, w)
		return
	}

	body := types.PatchPlugBody{}
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		handleError(err, w)
		return
	}

	plug := hs1xxplug.Hs1xxPlug{IPAddress: ip}

	switch body.State {
	case "on":
		{
			err = plug.TurnOn()
			break
		}
	case "off":
		{
			err = plug.TurnOff()
			break
		}
	default:
		{
			err = errors.New("invalid plug state must be one of: on, off")
		}
	}

	js, err := json.Marshal(body)

	if err != nil {
		handleError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, _ = w.Write(js)
}
