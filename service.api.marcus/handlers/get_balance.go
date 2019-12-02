package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Balance struct {
	Balance int `json:"balance"`
}

// todo: pass login credentials from calling service
func handleGetBalance(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	creds, err := RetrieveMarcusCredentials()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	isaRequestBody := strings.NewReader(fmt.Sprintf(MarcusBalanceRequestJson, creds.Username, creds.Password))
	resp, err := http.Post("http://service.browser:3000/scrape", "application/json", isaRequestBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	jsonBody := make(map[string]string)

	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	cleanMoneyString(jsonBody["balance"])
	// return balance in minor units to avoid floating point issues
	balance := Balance{Balance: int(100 * cleanMoneyString(jsonBody["balance"]))}

	js, err := json.Marshal(balance)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, _ = w.Write(js)
}

func cleanMoneyString(moneyStr string) float64 {
	val, _ := strconv.ParseFloat(strings.ReplaceAll(strings.TrimPrefix(moneyStr, "£"), ",", ""), 64)
	return val
}
