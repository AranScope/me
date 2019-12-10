package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var gauges = map[string]prometheus.Gauge{
}

func Float(name string, value float64) {
	var gauge prometheus.Gauge
	if _, exists := gauges[name]; !exists {
		gauge = promauto.NewGauge(prometheus.GaugeOpts{
			Name: name,
		})
		gauges[name] = gauge
	}

	if gauge != nil {
		gauge.Set(value)
	}
}
