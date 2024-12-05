package v1handler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/mapper"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/constants"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/http"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/tracing"
)

// Create Consumer
// @Summary Create a new Consumer
// @Description Create a new consumer with the provided request parameters
// @Tags Consumers
// @Accept json
// @Produce json
// @Param body body domain.CreateConsumerDTORequestParams{} true "body"
// @Success 200 {object} http.SuccessResponseDto{data=http.IDResponse} "Successfully created consumer"
// @Failure 400 {object} httpError.RestError{}
// @Failure 403 {object} httpError.RestError{}
// @Failure 404 {object} httpError.RestError{}
// @Failure 500 {object} httpError.RestError{}
// @Security ApiKeyAuth
// @Router /consumers [post]
func (h *Handler) Create(ctx context.Context, appCtx *app.RequestContext) {
	ctx, span := tracing.StartHttpServerTracerSpan(ctx, appCtx, "Consumers.Create")
	defer span.Finish()
	h.metrics.CreateConsumerHTTPRequest.Inc()

	dto := &domain.CreateConsumerDTORequestParams{}

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

// Find Consumers
// @Summary find list of Consumers
// @Description find list of consumers with the provided request parameters
// @Tags Consumers
// @Accept json
// @Produce json
// @Param q query string true "search q, determine either full_name or legal_name of the user"
// @Param page_size query string true "search page_size, determine the size of page / limit"
// @Param page_id query string true "search page_id, determine the number page / offset"
// @Success 200 {object} http.SuccessResponseDto{data=domain.ConsumerList}
// @Failure 400 {object} httpError.RestError{}
// @Failure 403 {object} httpError.RestError{}
// @Failure 404 {object} httpError.RestError{}
// @Failure 500 {object} httpError.RestError{}
// @Security ApiKeyAuth
// @Router /consumers [get]
func (h *Handler) Find(ctx context.Context, appCtx *app.RequestContext) {
	ctx, span := tracing.StartHttpServerTracerSpan(ctx, appCtx, "Consumers.Find")
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

// Update Consumer
// @Summary Update an existing Consumer by ID
// @Description Update an existing consumer by the provided ID and request parameters
// @Tags Consumers
// @Accept json
// @Produce json
// @Param id path string true "ID of the consumer transaction to update"
// @Param body body domain.UpdateConsumerDTORequestParams true "Request body for updating the consumer"
// @Success 200 {object} http.SuccessResponseDto{data=http.IDResponse} "Successfully updated consumer transaction"
// @Failure 400 {object} httpError.RestError{}
// @Failure 403 {object} httpError.RestError{}
// @Failure 404 {object} httpError.RestError{}
// @Failure 500 {object} httpError.RestError{}
// @Security ApiKeyAuth
// @Router /consumers/{id} [put]
func (h *Handler) Update(ctx context.Context, appCtx *app.RequestContext) {
	ctx, span := tracing.StartHttpServerTracerSpan(ctx, appCtx, "Consumers.Update")
	defer span.Finish()
	h.metrics.UpdateConsumerHTTPRequest.Inc()

	dto := &domain.UpdateConsumerDTORequestParams{}

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

// FindByID Consumer
// @Summary Find Consumer
// @Description Find consumer by associated id
// @Tags Consumers
// @Accept json
// @Produce json
// @Param id path string true "Any Associated id From your Source, this parameter is required"
// @Success 200 {object} http.SuccessResponseDto{data=domain.Consumer}
// @Failure 400 {object} httpError.RestError{}
// @Failure 403 {object} httpError.RestError{}
// @Failure 404 {object} httpError.RestError{}
// @Failure 500 {object} httpError.RestError{}
// @Security ApiKeyAuth
// @Router /consumers/{id} [get]
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
