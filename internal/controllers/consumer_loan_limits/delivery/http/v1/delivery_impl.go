package v1handler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_loan_limits/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_loan_limits/mapper"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/constants"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/http"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/tracing"
)

// Create Consumer Loan Limits
// @Summary Create a new Consumer Loan Limits
// @Description Create a new consumer Loan Limits with the provided request parameters
// @Tags Consumer-Loan-Limits
// @Accept json
// @Produce json
// @Param body body domain.CreateConsumerLoanLimitDTORequestParams{} true "body"
// @Success 200 {object} http.SuccessResponseDto{data=http.IDResponse} "Successfully created consumer"
// @Failure 400 {object} httpError.RestError{}
// @Failure 403 {object} httpError.RestError{}
// @Failure 404 {object} httpError.RestError{}
// @Failure 500 {object} httpError.RestError{}
// @Security ApiKeyAuth
// @Router /consumer-loan-limits [post]
func (h *Handler) Create(ctx context.Context, appCtx *app.RequestContext) {
	ctx, span := tracing.StartHttpServerTracerSpan(ctx, appCtx, "ConsumerLoanLimits.Create")
	defer span.Finish()
	h.metrics.CreateConsumerHTTPRequest.Inc()

	dto := &domain.CreateConsumerLoanLimitDTORequestParams{}

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

// Find Consumer Loan Limits
// @Summary find list of Consumer Loan Limits
// @Description find list of Consumer Loan Limits with the provided request parameters
// @Tags Consumer-Loan-Limits
// @Accept json
// @Produce json
// @Param customer_id query string true "customer_id, determine the customer id"
// @Param page_size query string true "search page_size, determine the size of page / limit"
// @Param page_id query string true "search page_id, determine the number page / offset"
// @Success 200 {object} http.SuccessResponseDto{data=domain.ConsumerLoanLimitList}
// @Failure 400 {object} httpError.RestError{}
// @Failure 403 {object} httpError.RestError{}
// @Failure 404 {object} httpError.RestError{}
// @Failure 500 {object} httpError.RestError{}
// @Security ApiKeyAuth
// @Router /consumer-loan-limits [get]
func (h *Handler) Find(ctx context.Context, appCtx *app.RequestContext) {
	ctx, span := tracing.StartHttpServerTracerSpan(ctx, appCtx, "ConsumerLoanLimits.Find")
	defer span.Finish()
	h.metrics.FindConsumerHTTPRequest.Inc()

	pq := helper.NewPaginationFromQueryParams(appCtx.Query(constants.PageSize), appCtx.Query(constants.PageID))
	consumerId, err := strconv.ParseInt(appCtx.Query("customer_id"), 10, 64)
	if err != nil {
		err := &http.ErrResp{
			Debug:      true,
			LogDetails: "consumerId.strconv.ParseInt.err",
			Log:        h.log,
			Err:        err,
		}
		tracing.TraceWithError(span, fmt.Errorf("%s:%s", err.LogDetails, err.Err))
		appCtx.AbortWithStatusJSON(http.ErrorResponse(err))
		return
	}

	dto := &domain.FetchDTORequestParams{
		ConsumerID: consumerId,
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

// Update Consumer Loan Limit
// @Summary Update an existing Consumer Loan Limit by ID
// @Description Update an existing Consumer Loan Limit by the provided ID and request parameters
// @Tags Consumer-Loan-Limits
// @Accept json
// @Produce json
// @Param id path string true "ID of the consumer transaction to update"
// @Param body body domain.UpdateConsumerLoanLimitDTORequestParams true "Request body for updating the consumer"
// @Success 200 {object} http.SuccessResponseDto{data=http.IDResponse} "Successfully updated consumer transaction"
// @Failure 400 {object} httpError.RestError{}
// @Failure 403 {object} httpError.RestError{}
// @Failure 404 {object} httpError.RestError{}
// @Failure 500 {object} httpError.RestError{}
// @Security ApiKeyAuth
// @Router /consumer-loan-limits/{id} [put]
func (h *Handler) Update(ctx context.Context, appCtx *app.RequestContext) {
	ctx, span := tracing.StartHttpServerTracerSpan(ctx, appCtx, "ConsumerLoanLimits.Update")
	defer span.Finish()
	h.metrics.UpdateConsumerHTTPRequest.Inc()

	dto := &domain.UpdateConsumerLoanLimitDTORequestParams{}

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

// FindByID Consumer Loan Limit
// @Summary Find Consumer Loan Limit
// @Description Find Consumer Loan Limit by associated id
// @Tags Consumer-Loan-Limits
// @Accept json
// @Produce json
// @Param id path string true "Any Associated id From your Source, this parameter is required"
// @Success 200 {object} http.SuccessResponseDto{data=domain.ConsumerLoanLimit}
// @Failure 400 {object} httpError.RestError{}
// @Failure 403 {object} httpError.RestError{}
// @Failure 404 {object} httpError.RestError{}
// @Failure 500 {object} httpError.RestError{}
// @Security ApiKeyAuth
// @Router /consumer-loan-limits/{id} [get]
func (h *Handler) FindByID(ctx context.Context, appCtx *app.RequestContext) {
	ctx, span := tracing.StartHttpServerTracerSpan(ctx, appCtx, "ConsumerLoanLimits.FindByID")
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
