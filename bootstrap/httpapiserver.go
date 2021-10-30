package bootstrap

import (
	"fmt"

	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	"github.com/lab259/cors"
	"github.com/valyala/fasthttp"
)

type HTTPAPIServer struct {
	logger applogger.Logger

	listeningAddress string
	routePrefix      string

	pprof DebugConfig

	httpServer *httpapi.Server
}

type RouteRegistry interface {
	Register(*httpapi.Server, string)
}

func NewHTTPAPIServer(logger applogger.Logger, config *Config) *HTTPAPIServer {
	httpListeningAddress := config.HTTP.ListeningAddress
	httpRoutePrefix := config.HTTP.RoutePrefix

	httpServer := httpapi.NewServer(
		httpListeningAddress,
	).WithLogger(
		logger,
	)

	if len(config.HTTP.CorsAllowedOrigins) != 0 {
		httpServer = httpServer.WithCors(cors.Options{
			AllowedOrigins: config.HTTP.CorsAllowedOrigins,
			AllowedMethods: config.HTTP.CorsAllowedMethods,
			AllowedHeaders: config.HTTP.CorsAllowedHeaders,
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
