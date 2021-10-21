package bootstrap

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/infrastructure/httpapi"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
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

type Handler interface {
	Register(*httpapi.Server, string)
}

func NewHTTPAPIServer(logger applogger.Logger, config *Config) *HTTPAPIServer {
	server := &HTTPAPIServer{
		logger: logger,

		listeningAddress: config.HTTP.ListeningAddress,
		routePrefix:      config.HTTP.RoutePrefix,

		pprof: config.Debug,
	}

	server.httpServer = httpapi.NewServer(
		server.listeningAddress,
	).WithLogger(
		server.logger,
	)

	if len(config.HTTP.CorsAllowedOrigins) != 0 {
		server.httpServer = server.httpServer.WithCors(cors.Options{
			AllowedOrigins: config.HTTP.CorsAllowedOrigins,
			AllowedMethods: config.HTTP.CorsAllowedMethods,
			AllowedHeaders: config.HTTP.CorsAllowedHeaders,
			Debug:          true,
		})
	}

	server.httpServer.GET(fmt.Sprintf("%s/api/v1/health", server.routePrefix), func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody([]byte("Ok"))
	})

	return server
}

func (server *HTTPAPIServer) RegisterHandlers(handlers []Handler) {
	for _, handler := range handlers {
		handler.Register(server.httpServer, server.routePrefix)
	}
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
