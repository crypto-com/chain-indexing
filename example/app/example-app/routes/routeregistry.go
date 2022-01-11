package routes

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	"github.com/valyala/fasthttp"
)

type RouteRegistry struct {
	routes []Route
}

type Route struct {
	Method  string
	path    string
	handler fasthttp.RequestHandler
}

func (registry *RouteRegistry) Register(server *httpapi.Server, routePrefix string) {
	if routePrefix == "/" {
		routePrefix = ""
	}

	for _, route := range registry.routes {
		registerRoute(server, routePrefix, route)
	}
}

func registerRoute(server *httpapi.Server, routePrefix string, route Route) {
	switch route.Method {
	case GET:
		server.GET(fmt.Sprintf("%s/%s", routePrefix, route.path), route.handler)
	}
}

const (
	GET = "GET"
)
