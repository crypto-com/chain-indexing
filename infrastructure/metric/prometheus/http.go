package prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Run(path, port string) {
	http.Handle(path, promhttp.Handler())
	http.ListenAndServe(":"+port, nil)
}
