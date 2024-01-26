package http

import (
	"context"
	"fmt"
	"gopaseto/internal/core/config"
	"io"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const defaultHost = "0.0.0.0"

type HttpServer interface {
	Start()
	io.Closer
}

type httpServer struct {
	Port   uint
	server *http.Server
}

// Close implements HttpServer.
func (h *httpServer) Close() error {
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Duration(3)*time.Second,
	)

	defer cancel()

	return h.server.Shutdown(ctx)
}

// Start implements HttpServer.
func (h *httpServer) Start() {
	go func() {
		if err := h.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf(
				"failed to stater HttpServer listen port %d, err=%s\n",
				h.Port, err.Error(),
			)
		}
	}()
	slog.Info("Start Service with port", "port", h.Port)
}

func NewHttpServer(router *gin.Engine, config config.HttpServerConfig) HttpServer {
	return &httpServer{
		Port: config.Port,
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", defaultHost, config.Port),
			Handler: router,
		},
	}
}
