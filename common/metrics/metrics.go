package metrics

import (
	"fmt"
	"github.com/AranScope/me/common/service"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var gauges = map[string]prometheus.Gauge{
}

func Float(name string, value float64) {
	var gauge prometheus.Gauge
	
	nameWithServicePrefix := fmt.Sprintf("%s_%s", service.Name(), name)

	if _, exists := gauges[nameWithServicePrefix]; !exists {
		gauge = promauto.NewGauge(prometheus.GaugeOpts{
			Name: nameWithServicePrefix,
		})
		gauges[nameWithServicePrefix] = gauge
	}

	if gauge != nil {
		gauge.Set(value)
	}
}
