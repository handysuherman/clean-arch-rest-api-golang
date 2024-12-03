package app

import (
	"context"
	"time"
)

const (
	waitShutdownDur = 3 * time.Second
)

func (a *app) shutdownProcess(ctx context.Context) error {
	<-ctx.Done()
	a.waitGracefulShutdown(waitShutdownDur)

	if err := a.server.Shutdown(ctx); err != nil {
		a.log.Warnf("a.server.shutdown.err: %v", err)
		return err
	}

	<-a.doneCh
	a.log.Info("server shutdown gracefully")
	return nil
}

func (a *app) waitGracefulShutdown(duration time.Duration) {
	go func() {
		time.Sleep(duration)
		a.doneCh <- struct{}{}
	}()
}
