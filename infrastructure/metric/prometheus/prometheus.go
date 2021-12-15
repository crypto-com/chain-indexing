package prometheus

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	defaultRegistry                       = prometheus.NewRegistry()
	Registerer      prometheus.Registerer = defaultRegistry
	Gatherer        prometheus.Gatherer   = defaultRegistry
)

func Run(path, port string) error {
	register()
	http.Handle(path, promhttp.InstrumentMetricHandler(
		Registerer, promhttp.HandlerFor(Gatherer, promhttp.HandlerOpts{}),
	))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		return err
	}

	return nil
}

func register() {
	Registerer.MustRegister(projectionExecTime)
	Registerer.MustRegister(projectionLatestHeight)
}
