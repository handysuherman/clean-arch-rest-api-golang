package app

import (
	"context"
	"fmt"

	appCtx "github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/middleware"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
)

func (a *app) runHTTPServer(ctx context.Context, cancel context.CancelFunc) error {
	addr := fmt.Sprintf("%s:%d", "0.0.0.0", a.cfg.App.Port)

	a.server = server.New(server.WithHostPorts(addr))

	mw := middleware.NewMiddlewareManager(a.log)

	a.server.Use(mw.RequestLoggerMiddleware())
	a.server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:50050"},
		AllowHeaders:     []string{"Origin, Content-Type, Accept"},
		AllowCredentials: true,
		AllowMethods: []string{
			consts.MethodGet,
			consts.MethodPatch,
			consts.MethodPost,
			consts.MethodDelete,
			consts.MethodHead,
			consts.MethodOptions,
		},
	}))

	if a.cfg.App.Environment != "production" {
		url := swagger.URL(fmt.Sprintf("%s:%d/swagger/doc.json", "0.0.0.0", a.cfg.App.Port))

		a.server.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url, swagger.DefaultModelsExpandDepth(-1)))
		a.server.GET("/docs", func(c context.Context, ctx *appCtx.RequestContext) {
			ctx.Redirect(consts.StatusMovedPermanently, []byte("/swagger/index.html"))
			ctx.Abort()
		})
	}

	a.bootstrapHandlers(ctx, cancel)

	err := a.server.Run()
	if err != nil {
		return err
	}

	a.log.Infof(
		"APP server running on: %v, SwaggerDocs Enabled: %v",
		addr,
		a.cfg.App.Environment != "production",
	)

	return nil
}

func (a *app) bootstrapHandlers(ctx context.Context, cancel context.CancelFunc) {
	a.affiliatedDealersHandlers().MapRoutes()
	a.consumersHandlers().MapRoutes()
	a.consumerLoanLimitsHandlers().MapRoutes()
}
