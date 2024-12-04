package v1handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/http"
)

func (h *Handler) Create(ctx context.Context, appCtx *app.RequestContext) {
	h.metrics.CreateAffiliatedDealerHTTPRequest.Inc()

	// TODO: implementation
	h.metrics.SuccessHTTPRequest.Inc()

	http.SuccessResponse(appCtx, &http.SuccessResp{
		Status:                consts.StatusOK,
		SuccessMetricsCounter: nil,
		Message:               "ok",
	})
}

func (h *Handler) Find(ctx context.Context, appCtx *app.RequestContext) {
	h.metrics.FindAffiliatedDealerHTTPRequest.Inc()

	// TODO: implementation
	h.metrics.SuccessHTTPRequest.Inc()

	http.SuccessResponse(appCtx, &http.SuccessResp{
		Status:                consts.StatusOK,
		SuccessMetricsCounter: nil,
		Message:               "ok",
	})
}

func (h *Handler) Update(ctx context.Context, appCtx *app.RequestContext) {
	h.metrics.UpdateAffiliatedDealerHTTPRequest.Inc()

	// TODO: implementation
	h.metrics.SuccessHTTPRequest.Inc()

	http.SuccessResponse(appCtx, &http.SuccessResp{
		Status:                consts.StatusOK,
		SuccessMetricsCounter: nil,
		Message:               "ok",
	})
}

func (h *Handler) FindByID(ctx context.Context, appCtx *app.RequestContext) {
	h.metrics.FindByIDAffiliatedDealerHTTPRequest.Inc()

	// TODO: implementation
	h.metrics.SuccessHTTPRequest.Inc()

	http.SuccessResponse(appCtx, &http.SuccessResp{
		Status:                consts.StatusOK,
		SuccessMetricsCounter: nil,
		Message:               "ok",
	})
}
