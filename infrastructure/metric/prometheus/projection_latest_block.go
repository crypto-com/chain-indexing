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

func RecordProjectionLatestHeight(projectionName string, value float64) {
	projectionLatestHeight.With(
		map[string]string{
			projectionLatestHeightProjectionLabel: projectionName,
		},
	).Set(value)
}
