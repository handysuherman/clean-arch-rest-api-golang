package app

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/go-playground/validator/v10"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/config"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/metrics"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type app struct {
	log               logger.Logger
	cfg               *config.Config
	server            *server.Hertz
	metrics           *metrics.Metrics
	metricsServer     *echo.Echo
	healthCheckServer *http.Server
	validator         *validator.Validate

	jaegerCloser io.Closer

	mysqlConnection *sql.DB
	redisConnection redis.UniversalClient

	doneCh chan struct{}
}

func New(log logger.Logger, cfg *config.Config) *app {
	return &app{
		log:       log.WithPrefix("APP"),
		cfg:       cfg,
		doneCh:    make(chan struct{}),
		validator: validator.New(),
		server: server.New(server.WithHostPorts(
			fmt.Sprintf("%s:%d", "0.0.0.0", cfg.Services.Internal.Port),
		)),
	}
}

func (a *app) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	a.metrics = metrics.New(a.cfg.Services.Internal)

	if err := a.mysql(ctx); err != nil {
		return err
	}
	a.log.Infof("connection to database successfully established...")
	defer a.mysqlConnection.Close()

	if err := a.runDBMigration(); err != nil {
		return err
	}
	a.log.Info("Database Migration successfully migrated ...")

	if err := a.redis(ctx); err != nil {
		return err
	}
	a.log.Infof("connection to redis successfully established...")
	defer a.redisConnection.Close()

	if err := a.jaeger(); err != nil {
		return err
	}
	a.log.Infof("connection to jaeger successfully established...")
	defer a.jaegerCloser.Close()

	go func() {
		if err := a.runHTTPServer(ctx, cancel); err != nil {
			a.log.Debugf("a.runHealthCheck.err: %v", err)
			cancel()
		}
	}()
	a.log.Infof(
		"APP server running on: :%v, w/ SwaggerDocs Enabled: %v, w/ TLS Enabled: %v",
		a.cfg.Services.Internal.Port,
		a.cfg.Services.Internal.Environment != "production",
		a.cfg.Services.Internal.EnableTLS,
	)

	go func() {
		if err := a.runHealthCheckServer(ctx); err != nil {
			a.log.Errorf("a.runHealthCheck: %v", err)
			cancel()
		}
	}()
	a.log.Infof("health check server is running on: %v...", a.cfg.Monitoring.Probes.Port)

	a.runMetricsServer(cancel)

	a.log.Infof("metrics server is running on port: %v ...", a.cfg.Monitoring.Probes.Prometheus.Port)

	a.shutdownProcess(ctx)

	return nil
}
