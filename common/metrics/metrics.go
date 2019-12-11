package metrics

import (
	"fmt"
	"github.com/AranScope/me/common/service"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strings"
)

var gauges = map[string]prometheus.Gauge{
}

func serviceNameToMetricName(name string) string {
	name = strings.ReplaceAll(name, "-", "_")
	name = strings.ReplaceAll(name, ".", "_")
	return name
}

func Float(name string, value float64) {
	var gauge prometheus.Gauge

	nameWithServicePrefix := fmt.Sprintf("%s_%s", serviceNameToMetricName(service.Name()), name)

	if _, exists := gauges[nameWithServicePrefix]; !exists {
		gauge = promauto.NewGauge(prometheus.GaugeOpts{
			Name: nameWithServicePrefix,
		})
		gauges[nameWithServicePrefix] = gauge
	}

	if gauge != nil {
		// todo: is this correct?
		gauge.SetToCurrentTime()
		gauge.Set(value)
	}
}
