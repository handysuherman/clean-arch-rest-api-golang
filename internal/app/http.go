package app

import (
	"context"
	"fmt"

	appCtx "github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/middleware"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
)

func (a *app) runHTTPServer(ctx context.Context, cancel context.CancelFunc) error {
	addr := fmt.Sprintf("%s:%d", "0.0.0.0", a.cfg.Services.Internal.Port)

	a.server = server.New(server.WithHostPorts(addr))

	mw := middleware.NewMiddlewareManager(a.log)
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	a.server.Use(mw.RequestLoggerMiddleware())
	a.server.Use(mw.CORSMiddleware())

	if a.cfg.Services.Internal.Environment != "production" {
		url := swagger.URL(fmt.Sprintf("%s:%d/swagger/doc.json", "0.0.0.0", a.cfg.Services.Internal.Port))

		a.server.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url, swagger.DefaultModelsExpandDepth(-1)))
		a.server.GET("/docs", func(c context.Context, ctx *appCtx.RequestContext) {
			ctx.Redirect(consts.StatusMovedPermanently, []byte("/swagger/index.html"))
			ctx.Abort()
		})
	}

	a.server.Use(mw.AuthPermission(a.cfg.Services.Internal.XApiKey, true))
	a.bootstrapHandlers(ctx, cancel)

	err := a.server.Run()
	if err != nil {
		return err
	}

	return nil
}

func (a *app) bootstrapHandlers(ctx context.Context, cancel context.CancelFunc) {
	a.affiliatedDealersHandlers().MapRoutes()
	a.consumersHandlers().MapRoutes()
	a.consumerLoanLimitsHandlers().MapRoutes()
	a.consumerTransactionsHandlers().MapRoutes()
}
