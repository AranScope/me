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

var counters = map[string]prometheus.Counter{}

func serviceNameToMetricName(name string) string {
	name = strings.ReplaceAll(name, "-", "_")
	name = strings.ReplaceAll(name, ".", "_")
	return name
}

func newMetricName(name string) string {
	return fmt.Sprintf("%s_%s", serviceNameToMetricName(service.Name()), name)
}

func Count(name string, add float64) {
	metricName := newMetricName(name)

	counter, exists := counters[metricName]
	if !exists {
		counter = promauto.NewCounter(prometheus.CounterOpts{
			Name: metricName,
		})
		counters[metricName] = counter
	}
	counter.Add(add)

}

func Float(name string, value float64) {
	metricName := newMetricName(name)
	gauge, exists := gauges[metricName]

	if !exists {
		gauge = promauto.NewGauge(prometheus.GaugeOpts{
			Name: metricName,
		})
		gauges[metricName] = gauge
	}

	gauge.Set(value)
}
