package main

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
	BalancePence int
}

// todo: pass login credentials from calling service
func GetBalance(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	creds, err := RetrieveVanguardCredentials()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	isaRequestBody := strings.NewReader(fmt.Sprintf(IsaRequestJson, creds.Username, creds.password))
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

	w.Header().Set("Content-Type", "application/json")

	cleanMoneyString(jsonBody["balance"])
	balance := Balance{BalancePence: int(100 * cleanMoneyString(jsonBody["balance"]))}

	js, err := json.Marshal(balance)
	if err != nil {
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

func main() {
	router := httprouter.New()
	router.GET("/balance", GetBalance)
	_ = http.ListenAndServe(":3002", router)
}
