package prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Run(path, port string) {
	register()
	http.Handle(path, promhttp.Handler())
	http.ListenAndServe(":"+port, nil)
}

func register() {
	prometheus.MustRegister(projectionExecTime)
	prometheus.MustRegister(projectionLatestHeight)
}
