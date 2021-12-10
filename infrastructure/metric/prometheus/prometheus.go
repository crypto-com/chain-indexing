package prometheus

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Run(path, port string) error {
	register()
	http.Handle(path, promhttp.Handler())
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		return err
	}

	return nil
}

func register() {
	prometheus.MustRegister(projectionExecTime)
	prometheus.MustRegister(projectionLatestHeight)
}
