package metrics

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/AranScope/me/common/service"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
	"strings"
)

var gauges = map[string]prometheus.Gauge{
}

type MetricsRoute struct {
	Service string `json:"service"`
	Port    int    `json:"port"`
	Path    string `json:"path"`
}

func Init() error {
	js, err := json.Marshal(MetricsRoute{
		Service: service.Name(),
		Port:    2112,
		Path:    "/metrics",
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, "http://service.prometheus-metrics-aggregator/register", bytes.NewReader(js))
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

func serviceNameToMetricName(name string) string {
	name = strings.ReplaceAll(name, "-", "_")
	name = strings.ReplaceAll(name, ".", "_")
	return name
}

func Float(name string, value float64) {
	nameWithServicePrefix := fmt.Sprintf("%s_%s", serviceNameToMetricName(service.Name()), name)
	gauge, exists := gauges[nameWithServicePrefix]

	if !exists {
		gauge = promauto.NewGauge(prometheus.GaugeOpts{
			Name: nameWithServicePrefix,
		})
		gauges[nameWithServicePrefix] = gauge
	}

	gauge.Set(value)
}
