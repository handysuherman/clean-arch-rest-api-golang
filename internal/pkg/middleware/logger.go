package middleware

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

func (mw *middlewareManager) RequestLoggerMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		start := time.Now()
		s := time.Since(start)

		if !mw.checkIgnoredURI(ctx.FullPath(), []string{"/metrics"}) {
			mw.log.HttpMiddlewareAccessLogger(string(ctx.Request.Header.Method()), string(ctx.Path()), ctx.Response.StatusCode(), int64(len(ctx.GetResponse().BodyBytes())), s)
		}

		ctx.Next(c)
	}
}

func (mw *middlewareManager) CORSMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x-api-key, _needsrefresh")
		ctx.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
		ctx.Header("Vary", "Origin")

		if string(ctx.Request.Method()) == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next(c)
	}
}
