package prometheus

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	defaultRegistry                       = prometheus.NewRegistry()
	Registerer      prometheus.Registerer = defaultRegistry
	Gatherer        prometheus.Gatherer   = defaultRegistry
)

const (
	READ_HEADER_TIMEOUT = 5 * time.Second
)

func Run(path, port string) error {
	register()
	http.Handle(path, promhttp.InstrumentMetricHandler(
		Registerer, promhttp.HandlerFor(Gatherer, promhttp.HandlerOpts{}),
	))
	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		ReadHeaderTimeout: READ_HEADER_TIMEOUT,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func register() {
	Registerer.MustRegister(projectionExecTime)
	Registerer.MustRegister(projectionLatestHeight)
}
