package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func handleGetBalance(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	creds, err := RetrieveMonzoCredentials()
	if err != nil {
		handleError(err, w)
		return
	}

	accountId := params.ByName("accountid")

	fmt.Println("https://api.monzo.com/balance?account_id=" + accountId)
	req, err := http.NewRequest(http.MethodGet, "https://api.monzo.com/balance?account_id="+accountId, nil)
	if err != nil {
		handleError(err, w)
		return
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", creds.AccessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		handleError(err, w)
		return
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		handleError(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(respBody)
}
