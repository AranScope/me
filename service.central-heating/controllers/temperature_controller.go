package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/AranScope/me/local/service.tplink-smart-plug/types"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	CurrentTemp   float64 = 20
	TargetTemp            = 18.8 // if temp hits target, switch off
	Threshold             = 0.1  // if delta goes below threshold, switch on
	RadiatorState         = "on"
)

func every(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}

func Start() {
	go func() {
		every(time.Second, Tick)
	}()
}

type TemperatureResponse struct {
	Temperature float64 `json:"temperature"`
	HeatIndex   float64 `json:"heat_index"`
	Humidity    float64 `json:"humidity"`
}

func getCurrentTemperature() (float64, error) {
	rsp, err := http.Get(fmt.Sprintf("http://%s:%s/data", os.Getenv("SENSOR_NODE_IP"), os.Getenv("SENSOR_NODE_PORT")))
	if err != nil {
		return 0, err
	}

	tempRspBytes, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return 0, err
	}

	tempRsp := &TemperatureResponse{}
	err = json.Unmarshal(tempRspBytes, tempRsp)
	if err != nil {
		return 0, err
	}

	return tempRsp.Temperature, nil
}

func setRadiatorState(state string) error {
	js, err := json.Marshal(types.PatchPlugBody{State: state})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, "http://localhost:8082/plug/192.168.1.119", bytes.NewReader(js))
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

	RadiatorState = state
	return nil
}

func Tick() {
	t, err := getCurrentTemperature()
	if err != nil {
		log.Println(err.Error())
		return
	}
	CurrentTemp = t
	deltaTemp := TargetTemp - CurrentTemp
	log.Printf("current temp: %.1fc, target temp: %.1fc", CurrentTemp, TargetTemp)

	if CurrentTemp > TargetTemp {
		if RadiatorState == "on" {
			log.Printf("switching radiator off")
			err := setRadiatorState("off")
			if err != nil {
				log.Println(err.Error())
				return
			}
		}

	} else if deltaTemp > Threshold {
		if RadiatorState == "off" {
			log.Printf("switching radiator on")
			err := setRadiatorState("on")
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}