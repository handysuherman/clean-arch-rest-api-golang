package middleware

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/constants"
	httpError "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/http_error"
)

func (mw *middlewareManager) AuthPermission(expectedKey string, debug bool) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		x_api_key := ctx.GetHeader(constants.XAPIKey)

		if len(string(x_api_key)) == 0 {
			ctx.AbortWithStatusJSON(httpError.ErrorCtxResponse(errors.New("required header api-key not provided"), debug))
			return
		}

		if string(x_api_key) != expectedKey {
			ctx.AbortWithStatusJSON(httpError.ErrorCtxResponse(errors.New("api-key not match"), debug))
			return
		}

		ctx.Next(c)
	}
}
