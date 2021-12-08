package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	projectionExecTimeName            = "projection_execution_time_per_block"
	projectionExecTimeProjectionLabel = "projection"
)

var (
	projectionExecTime = promauto.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: projectionExecTimeName,
		},
		[]string{
			projectionExecTimeProjectionLabel,
		},
	)
)

func RecordProjectionExecTime(projectionName string, timeInMilliseconds int64) {
	projectionExecTime.With(
		prometheus.Labels{
			projectionExecTimeProjectionLabel: projectionName,
		},
	).Observe(float64(timeInMilliseconds))
}
