package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/AranScope/me/service.central-heating/handlers"
	"github.com/AranScope/me/service.device-discovery/controller"
	"github.com/AranScope/me/service.tplink-smart-plug/types"
	"github.com/jasonlvhit/gocron"
	"io/ioutil"
	"net/http"
)

type GetTargetTemperatureResponse struct {
	TargetTemperature  float64 `json:"target_temperature"`
	CurrentTemperature float64 `json:"current_temperature"`
}

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

func setTargetTemperature(temp float64) error {
	js, err := json.Marshal(&handlers.SetTargetTemperatureRequest{
		Temperature: temp,
		Threshold:   0.5,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, "http://service.central-heating:8081/temperature", bytes.NewReader(js))
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
	//plugs := []string{"CC:32:E5:D7:7D:21", "CC:32:E5:D7:7D:86", "B0:BE:76:D6:D5:74"}
	//go (func() {
	//	gocron.Every(1).Day().At("22:30").Do(func() {
	//		for _, mac := range plugs {
	//			ip, err := getPlugIpAddress(mac)
	//			if err != nil {
	//				continue
	//			}
	//			_ = switchOffPlug(ip)
	//		}
	//	})
	//	<-gocron.Start()
	//})()
	go (func() {
		gocron.Every(1).Day().At("22:30").Do(func() {
			temp := 10.0
			fmt.Printf("🌡 Setting temperature to %.1f\n", temp)
			err := setTargetTemperature(temp)
			if err != nil {
				fmt.Printf("❌ Failed to set temperature: %v", err)
			}
		})
		<-gocron.Start()
	})()
	go (func() {
		gocron.Every(1).Day().At("19:00").Do(func() {
			temp := 16.5
			fmt.Printf("🌡 Setting temperature to %.1f\n", temp)
			err := setTargetTemperature(temp)
			if err != nil {
				fmt.Printf("❌ Failed to set temperature: %v", err)
			}
		})
		<-gocron.Start()
	})()
	go (func() {
		absentTemp := 10.0
		prevTemp := 0.0

		gocron.Every(5).Minutes().Do(func() {
			targetTemp, err := getCurrentTargetTemperature()
			if err != nil {
				fmt.Printf("❌ Failed to get device presence: %v", err)
				return
			}
			if targetTemp != absentTemp {
				rsp, err := http.Get("http://service.device-discovery/device/OnePlus3.lan")
				if err != nil {
					fmt.Printf("❌ Failed to get device presence: %v", err)
					return
				}
				if rsp.StatusCode != 200 {
					prevTemp = targetTemp
					fmt.Printf("🌡 User absent, setting temperature to %.1f\n", absentTemp)
					err := setTargetTemperature(absentTemp)
					if err != nil {
						fmt.Printf("❌ Failed to set temperature: %v", err)
					}
				} else {
					fmt.Printf("🌡 User present, Setting temperature back to %.1f\n", targetTemp)
					err := setTargetTemperature(prevTemp)
					if err != nil {
						fmt.Printf("❌ Failed to set temperature: %v", err)
					}
				}
			}
		})
		<-gocron.Start()
	})()

}

func getCurrentTargetTemperature() (float64, error) {
	rsp, err := http.Get("http://service.central-heating:8081/temperature")
	if err != nil {
		return 0, err
	}

	if rsp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("status returned not 200: actual: %d", rsp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return 0, err
	}

	body := GetTargetTemperatureResponse{
	}

	err = json.Unmarshal(bytes, &body)
	if err != nil {
		return 0, err
	}

	return body.TargetTemperature, nil
}
