package v1handler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/mapper"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/constants"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/http"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/tracing"
)

// Create Affiliated Dealers
// @Summary Create a new Affiliated Dealers
// @Description Create a new Affiliated Dealers with the provided request parameters
// @Tags Affiliated-Dealers
// @Accept json
// @Produce json
// @Param body body domain.CreateAffiliatedDealerDTORequestParams{} true "body"
// @Success 200 {object} http.SuccessResponseDto{data=http.IDResponse} "Successfully created consumer"
// @Failure 400 {object} httpError.RestError{}
// @Failure 403 {object} httpError.RestError{}
// @Failure 404 {object} httpError.RestError{}
// @Failure 500 {object} httpError.RestError{}
// @Security ApiKeyAuth
// @Router /affiliated-dealers [post]
func (h *Handler) Create(ctx context.Context, appCtx *app.RequestContext) {
	ctx, span := tracing.StartHttpServerTracerSpan(ctx, appCtx, "AffiliatedDealers.Create")
	defer span.Finish()
	h.metrics.CreateAffiliatedDealerHTTPRequest.Inc()

	dto := &domain.CreateAffiliatedDealerDTORequestParams{}

	if err := appCtx.Bind(dto); err != nil {
		err := &http.ErrResp{
			Debug:             true,
			LogDetails:        "appCtx.Bind.err",
			Log:               h.log,
			Err:               err,
			ErrMetricsCounter: h.metrics.ErrorHTTPRequest,
		}
		tracing.TraceWithError(span, fmt.Errorf("%s:%s", err.LogDetails, err.Err))
		appCtx.AbortWithStatusJSON(http.ErrorResponse(err))
		return
	}

	if err := h.validator.StructCtx(ctx, dto); err != nil {
		err := &http.ErrResp{
			Debug:             true,
			LogDetails:        "h.validator.StructCtx.err",
			Log:               h.log,
			Err:               err,
			ErrMetricsCounter: h.metrics.ErrorHTTPRequest,
		}
		tracing.TraceWithError(span, fmt.Errorf("%s:%s", err.LogDetails, err.Err))
		appCtx.AbortWithStatusJSON(http.ErrorResponse(err))
		return
	}

	res, err := h.usecase.Create(ctx, mapper.NewCreateRequestParams(dto))
	if err != nil {
		err := &http.ErrResp{
			Debug:             true,
			LogDetails:        "h.validator.StructCtx.err",
			Log:               h.log,
			Err:               err,
			ErrMetricsCounter: h.metrics.ErrorHTTPRequest,
		}
		appCtx.AbortWithStatusJSON(http.ErrorResponse(err))
		return
	}

	http.SuccessResponse(appCtx, &http.SuccessResp{
		Status: consts.StatusOK,
		Data: http.IDResponse{
			Id: strconv.FormatInt(res, 10),
		},
		SuccessMetricsCounter: h.metrics.SuccessHTTPRequest,
		Message:               "OK",
	})
}

// Find Affiliated Dealers
// @Summary find list of Affiliated Dealers
// @Description find list of Affiliated Dealers with the provided request parameters
// @Tags Affiliated-Dealers
// @Accept json
// @Produce json
// @Param q query string true "search q, determine either full_name or legal_name of the user"
// @Param page_size query string true "search page_size, determine the size of page / limit"
// @Param page_id query string true "search page_id, determine the number page / offset"
// @Success 200 {object} http.SuccessResponseDto{data=domain.AffiliatedDealerList}
// @Failure 400 {object} httpError.RestError{}
// @Failure 403 {object} httpError.RestError{}
// @Failure 404 {object} httpError.RestError{}
// @Failure 500 {object} httpError.RestError{}
// @Security ApiKeyAuth
// @Router /affiliated-dealers [get]
func (h *Handler) Find(ctx context.Context, appCtx *app.RequestContext) {
	ctx, span := tracing.StartHttpServerTracerSpan(ctx, appCtx, "AffiliatedDealers.Find")
	defer span.Finish()
	h.metrics.FindConsumerHTTPRequest.Inc()

	pq := helper.NewPaginationFromQueryParams(appCtx.Query(constants.PageSize), appCtx.Query(constants.PageID))
	query := appCtx.Query(constants.Q)

	dto := &domain.FetchDTORequestParams{
		Query:      &query,
		Pagination: pq,
	}

	res, err := h.usecase.List(ctx, mapper.NewFetchParams(dto))
	if err != nil {
		err := &http.ErrResp{
			Debug:      true,
			LogDetails: "h.usecase.List.err",
			Log:        h.log,
			Err:        err,
		}
		appCtx.AbortWithStatusJSON(http.ErrorResponse(err))
		return
	}

	http.SuccessResponse(appCtx, &http.SuccessResp{
		Status:                consts.StatusOK,
		Data:                  res,
		SuccessMetricsCounter: h.metrics.SuccessHTTPRequest,
		Message:               "OK",
	})
}

// Update Affiliated Dealer
// @Summary Update an existing Affiliated Dealer by ID
// @Description Update an existing Affiliated Dealer by the provided ID and request parameters
// @Tags Affiliated-Dealers
// @Accept json
// @Produce json
// @Param id path string true "ID of the affiliated dealer transaction to update"
// @Param body body domain.UpdateAffiliatedDealerDTORequestParams true "Request body for updating the affiliated dealer"
// @Success 200 {object} http.SuccessResponseDto{data=http.IDResponse} "Successfully updated affiliated dealer transaction"
// @Failure 400 {object} httpError.RestError{}
// @Failure 403 {object} httpError.RestError{}
// @Failure 404 {object} httpError.RestError{}
// @Failure 500 {object} httpError.RestError{}
// @Security ApiKeyAuth
// @Router /affiliated-dealers/{id} [put]
func (h *Handler) Update(ctx context.Context, appCtx *app.RequestContext) {
	ctx, span := tracing.StartHttpServerTracerSpan(ctx, appCtx, "AffiliatedDealers.Update")
	defer span.Finish()
	h.metrics.UpdateConsumerHTTPRequest.Inc()

	dto := &domain.UpdateAffiliatedDealerDTORequestParams{}

	if err := appCtx.Bind(dto); err != nil {
		err := &http.ErrResp{
			Debug:             true,
			LogDetails:        "appCtx.Bind.err",
			Log:               h.log,
			Err:               err,
			ErrMetricsCounter: h.metrics.ErrorHTTPRequest,
		}
		tracing.TraceWithError(span, fmt.Errorf("%s:%s", err.LogDetails, err.Err))
		appCtx.AbortWithStatusJSON(http.ErrorResponse(err))
		return
	}

	id, err := strconv.ParseInt(appCtx.Param(constants.ID), 10, 64)
	if err != nil {
		err := &http.ErrResp{
			Debug:             true,
			LogDetails:        "strconv.ParseInt.err",
			Log:               h.log,
			Err:               err,
			ErrMetricsCounter: h.metrics.ErrorHTTPRequest,
		}
		tracing.TraceWithError(span, fmt.Errorf("%s:%s", err.LogDetails, err.Err))
		appCtx.AbortWithStatusJSON(http.ErrorResponse(err))
		return
	}

	res, err := h.usecase.Update(ctx, id, mapper.NewUpdateRequestParams(dto))
	if err != nil {
		err := &http.ErrResp{
			Debug:             true,
			LogDetails:        "h.validator.StructCtx.err",
			Log:               h.log,
			Err:               err,
			ErrMetricsCounter: h.metrics.ErrorHTTPRequest,
		}
		appCtx.AbortWithStatusJSON(http.ErrorResponse(err))
		return
	}

	http.SuccessResponse(appCtx, &http.SuccessResp{
		Status: consts.StatusOK,
		Data: http.IDResponse{
			Id: strconv.FormatInt(res, 10),
		},
		SuccessMetricsCounter: h.metrics.SuccessHTTPRequest,
		Message:               "OK",
	})
}

// FindByID Affiliated Dealer
// @Summary Find Affiliated Dealer
// @Description Find Affiliated Dealer by associated id
// @Tags Affiliated-Dealers
// @Accept json
// @Produce json
// @Param id path string true "Any Associated id From your Source, this parameter is required"
// @Success 200 {object} http.SuccessResponseDto{data=domain.AffiliatedDealer}
// @Failure 400 {object} httpError.RestError{}
// @Failure 403 {object} httpError.RestError{}
// @Failure 404 {object} httpError.RestError{}
// @Failure 500 {object} httpError.RestError{}
// @Security ApiKeyAuth
// @Router /affiliated-dealers/{id} [get]
func (h *Handler) FindByID(ctx context.Context, appCtx *app.RequestContext) {
	ctx, span := tracing.StartHttpServerTracerSpan(ctx, appCtx, "Consumers.FindByID")
	defer span.Finish()
	h.metrics.FindByIDConsumerHTTPRequest.Inc()

	id, err := strconv.ParseInt(appCtx.Param(constants.ID), 10, 64)
	if err != nil {
		err := &http.ErrResp{
			Debug:             true,
			LogDetails:        "strconv.ParseInt.err",
			Log:               h.log,
			Err:               err,
			ErrMetricsCounter: h.metrics.ErrorHTTPRequest,
		}
		tracing.TraceWithError(span, fmt.Errorf("%s:%s", err.LogDetails, err.Err))
		appCtx.AbortWithStatusJSON(http.ErrorResponse(err))
		return
	}

	res, err := h.usecase.FindByID(ctx, id)
	if err != nil {
		err := &http.ErrResp{
			Debug:             true,
			LogDetails:        "h.usecase.FindByID.err",
			Log:               h.log,
			Err:               err,
			ErrMetricsCounter: h.metrics.ErrorHTTPRequest,
		}
		tracing.TraceWithError(span, fmt.Errorf("%s:%s", err.LogDetails, err.Err))
		appCtx.AbortWithStatusJSON(http.ErrorResponse(err))
		return
	}

	http.SuccessResponse(appCtx, &http.SuccessResp{
		Status:                consts.StatusOK,
		SuccessMetricsCounter: h.metrics.SuccessHTTPRequest,
		Data:                  res,
		Message:               "OK",
	})
}
