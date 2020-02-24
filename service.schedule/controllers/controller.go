package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/AranScope/me/service.device-discovery/controller"
	"github.com/AranScope/me/service.tplink-smart-plug/types"
	"github.com/jasonlvhit/gocron"
	"io/ioutil"
	"net/http"
)

func getPlugIpAddress(mac string) (string, error) {
	rsp, err := http.Get("http://service.device-discovery/device/" + mac)
	if err != nil {
		return "", err
	}
	if rsp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status returned not 200: actual: %d", rsp.StatusCode)
	}

	rspBytes, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}

	device := controller.Device{}
	err = json.Unmarshal(rspBytes, rsp.Body)
	if err != nil {
		return "", err
	}

	return device.IpAddr, nil
}

func switchOffPlug(ip string) error {
	js, err := json.Marshal(types.PatchPlugBody{State: "off"})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, "http://service.tplink-smart-plug:8082/plug/"+ip, bytes.NewReader(js))
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status returned not 200: actual: %d", resp.StatusCode)
	}

	return nil
}

func Init() {
	plugs := []string{"CC:32:E5:D7:7D:21", "CC:32:E5:D7:7D:86", "B0:BE:76:D6:D5:74"}
	go (func() {
		gocron.Every(1).Day().At("22:30").Do(func() {
			for _, mac := range plugs {
				ip, err := getPlugIpAddress(mac)
				if err != nil {
					continue
				}
				_ = switchOffPlug(ip)
			}
		})
		<-gocron.Start()
	})()
}
