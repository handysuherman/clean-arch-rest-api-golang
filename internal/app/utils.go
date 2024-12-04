package app

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"time"
)

const (
	waitShutdownDur = 3 * time.Second
)

func (a *app) loadTLsCerts(caPath string, certPath string, keyPath string) (*tls.Config, error) {
	ca_cert, err := os.ReadFile(caPath)
	if err != nil {
		a.log.Fatalf("ca_cert.os.ReadFile.err: %v", err)
		return nil, err
	}

	cert_pool := x509.NewCertPool()
	if ok := cert_pool.AppendCertsFromPEM(ca_cert); !ok {
		err := fmt.Errorf("failed to append PEM")
		a.log.Fatalf("cert_pool.AppendCertsFromPEM.%v", err)
	}

	client_cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		a.log.Fatalf("client_cert.tls.LoadX509KeyPair.err: %v", err)
		return nil, err
	}

	return &tls.Config{
		RootCAs:      cert_pool,
		Certificates: []tls.Certificate{client_cert},
	}, nil
}

func (a *app) shutdownProcess(ctx context.Context) {
	<-ctx.Done()
	a.waitGracefulShutdown(waitShutdownDur)

	a.log.Info("shutdowning server ...")

	if err := a.server.Shutdown(ctx); err != nil {
		a.log.Warnf("a.server.shutdown.err: %v", err)
	}

	if err := a.metricsServer.Shutdown(ctx); err != nil {
		a.log.Infof("a.metricsServer.Shutdown.err: %v", err)
	}

	if err := a.gracefulShutDownHealthCheckServer(ctx); err != nil {
		a.log.Infof("a.gracefulShutDownHealthCheckServer.err: %v", err)
	}

	if err := a.mysqlConnection.Close(); err != nil {
		a.log.Infof("a.mysqlConnection.Close.err: %v", err)
	}

	if err := a.redisConnection.Close(); err != nil {
		a.log.Infof("a.redisConnection.Close.err: %v", err)
	}

	if err := a.jaegerCloser.Close(); err != nil {
		a.log.Infof("a.jaegerCloser.Close.err: %v", err)
	}

	<-a.doneCh
	a.log.Info("server shutdown gracefully")
}

func (a *app) waitGracefulShutdown(duration time.Duration) {
	go func() {
		time.Sleep(duration)
		a.doneCh <- struct{}{}
	}()
}
