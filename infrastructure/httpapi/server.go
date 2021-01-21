package httpapi

import (
	"fmt"
	"strings"

	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/fasthttp/router"
	"github.com/lab259/cors"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/pprofhandler"
)

type Server struct {
	router           *router.Router
	listeningAddress string

	middlewares      []Middleware
	corsMiddleware   Middleware
	loggerMiddleware Middleware
}

func NewServer(listeningAddress string) *Server {
	r := router.New()
	middlewares := make([]Middleware, 0)
	return &Server{
		r,
		listeningAddress,

		middlewares,
		nil,
		nil,
	}
}

func (server *Server) GET(path string, handler fasthttp.RequestHandler) *Server {
	server.router.GET(path, handler)
	return server
}

func (server *Server) Use(middleware Middleware) *Server {
	server.middlewares = append(server.middlewares, middleware)
	return server
}

func (server *Server) WithCors(options cors.Options) *Server {
	server.corsMiddleware = cors.New(options).Handler
	return server
}

func (server *Server) WithLogger(logger applogger.Logger) *Server {
	logger = logger.WithFields(applogger.LogFields{
		"module": "httpapi",
	})
	server.loggerMiddleware = func(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			logger.Infof("%s %s", ctx.Method(), ctx.Path())

			handler(ctx)
		}
	}
	server.WithPanicHandler(func(ctx *fasthttp.RequestCtx, err interface{}) {
		logger.Errorf("%s %s: PANIC %v", ctx.Method(), ctx.Path(), err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.WriteString("Internal server error") // nolint:errcheck
	})
	return server
}

func (server *Server) WithPanicHandler(handler func(ctx *fasthttp.RequestCtx, err interface{})) *Server {
	server.router.PanicHandler = handler
	return server
}

func (server *Server) WithPprof(path string) *Server {
	server.router.ANY(fmt.Sprintf("%s/{path:*}", strings.TrimRight(path, "/")), pprofhandler.PprofHandler)
	return server
}

func (server *Server) ListenAndServe() error {
	handler := server.router.Handler
	if server.corsMiddleware != nil {
		handler = server.corsMiddleware(handler)
	}
	if server.loggerMiddleware != nil {
		handler = server.loggerMiddleware(handler)
	}
	for _, middleware := range server.middlewares {
		handler = middleware(handler)
	}
	return fasthttp.ListenAndServe(server.listeningAddress, handler)
}

type Middleware = func(fasthttp.RequestHandler) fasthttp.RequestHandler
