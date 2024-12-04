package tracing

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func StartHttpServerTracerSpan(c context.Context, ctx *app.RequestContext, operationName string) (context.Context, opentracing.Span) {
	headers := make(map[string][]string)
	ctx.Request.Header.VisitAll(func(k, v []byte) {
		key := ctx.GetString(string(k))
		headers[key] = append(headers[key], ctx.GetString(string(v)))
	})

	spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(headers))
	if err != nil {
		serverSpan := opentracing.GlobalTracer().StartSpan(operationName)
		ctx := opentracing.ContextWithSpan(c, serverSpan)
		return ctx, serverSpan
	}

	serverSpan := opentracing.GlobalTracer().StartSpan(operationName, ext.RPCServerOption(spanCtx))
	cc := opentracing.ContextWithSpan(c, serverSpan)
	return cc, serverSpan
}

func TraceError(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
}

func TraceWithError(span opentracing.Span, err error) error {
	if err != nil {
		span.SetTag("error", true)
		span.LogKV("error_code", err.Error())
	}
	return err
}
