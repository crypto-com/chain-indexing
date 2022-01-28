package bootstrap

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/bootstrap/config"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	"github.com/lab259/cors"
	"github.com/valyala/fasthttp"
)

type HTTPAPIServer struct {
	logger applogger.Logger

	listeningAddress string
	routePrefix      string

	pprof config.Debug

	httpServer *httpapi.Server
}

type RouteRegistry interface {
	Register(*httpapi.Server, string)
}

func NewHTTPAPIServer(logger applogger.Logger, config *config.Config) *HTTPAPIServer {
	httpListeningAddress := config.HTTPService.ListeningAddress
	httpRoutePrefix := config.HTTPService.RoutePrefix

	httpServer := httpapi.NewServer(
		httpListeningAddress,
	).WithLogger(
		logger,
	)

	if len(config.HTTPService.CorsAllowedOrigins) != 0 {
		httpServer = httpServer.WithCors(cors.Options{
			AllowedOrigins: config.HTTPService.CorsAllowedOrigins,
			AllowedMethods: config.HTTPService.CorsAllowedMethods,
			AllowedHeaders: config.HTTPService.CorsAllowedHeaders,
			Debug:          true,
		})
	}

	return &HTTPAPIServer{
		logger: logger,

		listeningAddress: httpListeningAddress,
		routePrefix:      httpRoutePrefix,

		pprof: config.Debug,

		httpServer: httpServer,
	}
}

func (server *HTTPAPIServer) RegisterRoutes(registry RouteRegistry) {
	server.httpServer.GET(fmt.Sprintf("%s/api/v1/health", server.routePrefix), func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody([]byte("Ok"))
	})

	registry.Register(server.httpServer, server.routePrefix)
}

func (server *HTTPAPIServer) Run() error {
	if server.pprof.PprofEnable {
		pprofServer := httpapi.NewServer(
			server.pprof.PprofListeningAddress,
		).WithLogger(
			server.logger,
		)
		fixPath := "/debug/pprof"
		pprofServer = pprofServer.WithPprof(fixPath)
		go func() {
			server.logger.Infof("pprof server start listening on: %s%s", server.pprof.PprofListeningAddress, fixPath)
			if err := pprofServer.ListenAndServe(); err != nil {
				panic(fmt.Errorf("error listening and serving HTTP pprof server: %w", err))
			}
		}()
	}

	server.logger.Infof("server start listening on: %s", server.listeningAddress)
	if err := server.httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("error listening and serving HTTP API server: %v", err)
	}

	return nil
}
