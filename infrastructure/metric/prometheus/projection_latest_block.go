package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	projectionLatestHeightName            = "projection_latest_block_height"
	projectionLatestHeightProjectionLabel = "projection"
)

var (
	projectionLatestHeight = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: projectionLatestHeightName,
		},
		[]string{
			projectionLatestHeightProjectionLabel,
		},
	)
)

func RecordProjectionLatestHeight(projectionName string, height int64) {
	projectionLatestHeight.With(
		prometheus.Labels{
			projectionLatestHeightProjectionLabel: projectionName,
		},
	).Set(float64(height))
}
