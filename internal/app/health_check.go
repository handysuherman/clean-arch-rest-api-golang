package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/constants"
	"github.com/heptiolabs/healthcheck"
)

func (a *app) runHealthCheckServer(ctx context.Context) error {
	health := healthcheck.NewHandler()

	mux := http.NewServeMux()
	mux.HandleFunc(a.cfg.Monitoring.Probes.LivenessPath, health.LiveEndpoint)
	mux.HandleFunc(a.cfg.Monitoring.Probes.ReadinessPath, health.ReadyEndpoint)

	a.healthCheckServer = &http.Server{
		Handler:      mux,
		Addr:         a.cfg.Monitoring.Probes.Port,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}

	a.configureHealthCheckEndpoints(ctx, health)
	return a.healthCheckServer.ListenAndServe()
}

func (a *app) configureHealthCheckEndpoints(ctx context.Context, health healthcheck.Handler) {
	health.AddReadinessCheck(constants.MySQL, healthcheck.AsyncWithContext(ctx, func() error {
		return a.mysqlConnection.PingContext(ctx)
	}, time.Duration(a.cfg.Monitoring.Probes.CheckInterval)*time.Second))
}

func (a *app) gracefulShutDownHealthCheckServer(ctx context.Context) error {
	if err := a.healthCheckServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("a.healthCheckServer.Shutdown.err: %v", err)
	}

	return nil
}
