package app

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/config"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/metrics"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/logger"
	"github.com/labstack/echo/v4"
)

type app struct {
	log           logger.Logger
	cfg           *config.Config
	server        *server.Hertz
	metrics       *metrics.Metrics
	metricsServer *echo.Echo

	jaegerCloser io.Closer

	doneCh chan struct{}
}

func New(log logger.Logger, cfg *config.Config) *app {
	return &app{
		log:    log.WithPrefix("APP"),
		cfg:    cfg,
		doneCh: make(chan struct{}),
		server: server.New(server.WithHostPorts(
			fmt.Sprintf("%s:%d", "0.0.0.0", cfg.Services.Internal.Port),
		)),
	}
}

func (a *app) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	a.metrics = metrics.New(a.cfg.Services.Internal)

	go func() {
		if err := a.runHTTPServer(ctx, cancel); err != nil {
			a.log.Debugf("a.runHealthCheck.err: %v", err)
			cancel()
		}
	}()

	a.shutdownProcess(ctx)

	return nil
}
