package prometheus

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Run(path, port string) {
	fmt.Println("prometheus server running")
	http.Handle(path, promhttp.Handler())
	http.ListenAndServe(":"+port, nil)
}
