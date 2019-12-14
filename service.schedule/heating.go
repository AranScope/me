package main

import (
	"encoding/json"
	log "github.com/mgutz/logxi/v1"
	"net/http"
	"strings"
	"time"
)

type ScheduledTemperature struct {
	time        time.Duration
	temperature float64
}

type TemperatureRequest struct {
	temperature float64
}

var weekday = []ScheduledTemperature{
	{
		time:        0,
		temperature: 16,
	},
	{
		time:        time.Hour * 18,
		temperature: 20,
	},
	{
		time:        time.Hour * 22,
		temperature: 16,
	},
}

var weekend = []ScheduledTemperature{
	{
		time:        0,
		temperature: 16,
	},
	{
		time:        time.Hour * 9,
		temperature: 20,
	},
	{
		time:        time.Hour * 22,
		temperature: 16,
	},
}

var schedule = map[time.Weekday][]ScheduledTemperature{
	time.Monday:    weekday,
	time.Tuesday:   weekday,
	time.Wednesday: weekday,
	time.Thursday:  weekday,
	time.Friday:    weekday,
	time.Saturday:  weekend,
	time.Sunday:    weekend,
}

func every(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}

func Start() {
	go func() {
		every(time.Minute*5, Tick)
	}()
}

func Tick() {
	currTime := time.Now()
	todaysSchedule := schedule[currTime.Weekday()]

	for i := 0; i < len(todaysSchedule)-1; i++ {
		if todaysSchedule[i+1].time.Minutes() > float64(currTime.Minute()) {
			// use the current time
			temp := todaysSchedule[i].temperature
			err := changeTemperature(temp)
			if err != nil {
				log.Error("error changing temperature: %v", err)
			}
		}
	}
}

func changeTemperature(temperature float64) error {

	t := TemperatureRequest{temperature: temperature}
	bytes, err := json.Marshal(t)
	if err != nil {
		return err
	}

	reader := strings.NewReader(string(bytes))
	_, err = http.Post("http://service.central-heating:8081/temperature", "application/json", reader)
	if err != nil {
		return err
	}

	return nil
}
