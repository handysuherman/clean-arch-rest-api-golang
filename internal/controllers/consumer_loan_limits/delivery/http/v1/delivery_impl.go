package v1handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/http"
)

func (h *Handler) Create(ctx context.Context, appCtx *app.RequestContext) {
	h.metrics.CreateConsumerLoanLimitHTTPRequest.Inc()

	// TODO: implementation
	h.metrics.SuccessHTTPRequest.Inc()

	http.SuccessResponse(appCtx, &http.SuccessResp{
		Status:                consts.StatusOK,
		SuccessMetricsCounter: nil,
		Message:               "ok",
	})
}

func (h *Handler) Find(ctx context.Context, appCtx *app.RequestContext) {
	h.metrics.FindConsumerLoanLimitHTTPRequest.Inc()

	// TODO: implementation
	h.metrics.SuccessHTTPRequest.Inc()

	http.SuccessResponse(appCtx, &http.SuccessResp{
		Status:                consts.StatusOK,
		SuccessMetricsCounter: nil,
		Message:               "ok",
	})
}

func (h *Handler) Update(ctx context.Context, appCtx *app.RequestContext) {
	h.metrics.UpdateConsumerLoanLimitHTTPRequest.Inc()

	// TODO: implementation
	h.metrics.SuccessHTTPRequest.Inc()

	http.SuccessResponse(appCtx, &http.SuccessResp{
		Status:                consts.StatusOK,
		SuccessMetricsCounter: nil,
		Message:               "ok",
	})
}

func (h *Handler) FindByID(ctx context.Context, appCtx *app.RequestContext) {
	h.metrics.FindByIDConsumerLoanLimitHTTPRequest.Inc()

	// TODO: implementation
	h.metrics.SuccessHTTPRequest.Inc()

	http.SuccessResponse(appCtx, &http.SuccessResp{
		Status:                consts.StatusOK,
		SuccessMetricsCounter: nil,
		Message:               "ok",
	})
}
