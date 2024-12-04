package logger

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/constants"
)

type Logger interface {
	WithPrefix(prefix string) *appLogger
	Print(level zerolog.Level, args ...interface{})
	Printf(ctx context.Context, format string, v ...interface{})
	Fatalf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	HttpMiddlewareAccessLogger(method, uri string, status int, size int64, time time.Duration)
}

type appLogger struct {
	prefix *string
}

func NewLogger() *appLogger {
	return &appLogger{}
}

func (logger *appLogger) WithPrefix(prefix string) *appLogger {
	return &appLogger{prefix: &prefix}
}

func (logger *appLogger) Print(level zerolog.Level, args ...interface{}) {
	msg := fmt.Sprint(args...)

	if logger.prefix != nil {
		msg = fmt.Sprintf("[%s]: %v", *logger.prefix, msg)
	}

	log.WithLevel(level).Msg(msg)
}

func (logger *appLogger) Printf(ctx context.Context, format string, v ...interface{}) {
	log.WithLevel(zerolog.DebugLevel).Msgf(format, v...)
}

func (logger *appLogger) Errorf(format string, v ...interface{}) {
	if logger.prefix != nil {
		format = fmt.Sprintf("[%s]: %s", *logger.prefix, format)
	}
	log.WithLevel(zerolog.ErrorLevel).Msgf(format, v...)
}

func (logger *appLogger) Fatalf(format string, v ...interface{}) {
	if logger.prefix != nil {
		format = fmt.Sprintf("[%s]: %s", *logger.prefix, format)
	}
	log.WithLevel(zerolog.FatalLevel).Msgf(format, v...)
}

func (logger *appLogger) Warnf(format string, v ...interface{}) {
	if logger.prefix != nil {
		format = fmt.Sprintf("[%s]: %s", *logger.prefix, format)
	}
	log.WithLevel(zerolog.WarnLevel).Msgf(format, v...)
}

func (logger *appLogger) Infof(format string, v ...interface{}) {
	if logger.prefix != nil {
		format = fmt.Sprintf("[%s]: %s", *logger.prefix, format)
	}
	log.WithLevel(zerolog.InfoLevel).Msgf(format, v...)
}

func (logger *appLogger) Debugf(format string, v ...interface{}) {
	if logger.prefix != nil {
		format = fmt.Sprintf("[%s]: %s", *logger.prefix, format)
	}
	log.WithLevel(zerolog.DebugLevel).Msgf(format, v...)
}

func (logger *appLogger) Debug(args ...interface{}) {
	logger.Print(zerolog.DebugLevel, args...)
}

func (logger *appLogger) Info(args ...interface{}) {
	logger.Print(zerolog.InfoLevel, args...)
}

func (logger *appLogger) Warn(args ...interface{}) {
	logger.Print(zerolog.WarnLevel, args...)
}

func (logger *appLogger) Error(args ...interface{}) {
	logger.Print(zerolog.ErrorLevel, args...)
}

func (logger *appLogger) Fatal(args ...interface{}) {
	logger.Print(zerolog.FatalLevel, args...)
}

func (logger *appLogger) HttpMiddlewareAccessLogger(
	method, uri string,
	status int,
	size int64,
	time time.Duration,
) {
	logs := log.WithLevel(zerolog.InfoLevel)

	if status != http.StatusOK {
		logs = log.WithLevel(zerolog.WarnLevel)
	}

	msg := fmt.Sprintf("Received %v Request", constants.HTTP)

	if logger.prefix != nil {
		msg = fmt.Sprintf("[%s]: %s", *logger.prefix, msg)
	}

	logs.
		Str(constants.Protocol, constants.HTTP).
		Str(constants.URI, uri).
		Str(constants.Method, method).
		Int(constants.Status, status).
		Int64(constants.Size, size).
		Dur(constants.Duration, time).
		Msgf(msg)
}
