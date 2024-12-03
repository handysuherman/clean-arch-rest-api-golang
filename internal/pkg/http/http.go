package http

import (
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	httpError "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/http_error"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
)

type SuccessResponseDto struct {
	Status    int         `json:"status"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}

type SuccessResp struct {
	Status                int         `json:"status"`
	Data                  interface{} `json:"data"`
	Message               string      `json:"message,omitempty"`
	SuccessMetricsCounter prometheus.Counter
}

type ErrResp struct {
	Err               error
	LogDetails        string
	Log               logger.Logger
	Debug             bool
	ErrMetricsCounter prometheus.Counter
}

func SuccessResponse(ctx *app.RequestContext, arg *SuccessResp) {
	if arg.SuccessMetricsCounter != nil {
		arg.SuccessMetricsCounter.Inc()
	}

	ctx.AbortWithStatusJSON(arg.Status, &SuccessResponseDto{
		Status:    arg.Status,
		Timestamp: time.Now().Unix(),
		Data:      arg.Data,
	})
}

func ErrorResponse(arg *ErrResp) (int, interface{}) {
	if arg.ErrMetricsCounter != nil {
		arg.ErrMetricsCounter.Inc()
	}

	arg.Log.Warnf("%s: %v", arg.LogDetails, arg.Err)

	return httpError.ErrorCtxResponse(arg.Err, arg.Debug)
}

func TraceError(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
}
