package v1handler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/mapper"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/constants"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/http"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/tracing"
)

// Create Consumer Transaction
// @Summary Create a new Consumer-Transaction
// @Description Create a new consumer transaction with the provided request parameters
// @Tags Consumer-Transactions
// @Accept json
// @Produce json
// @Param body body domain.CreateConsumerTransactionDTORequestParams{} true "body"
// @Param x-idempotency-key header string false "Optional idempotency key to ensure the request is only processed once"
// @Success 200 {object} http.SuccessResponseDto{data=http.IDResponse} "Successfully created consumer transaction"
// @Failure 400 {object} httpError.RestError{}
// @Failure 403 {object} httpError.RestError{}
// @Failure 404 {object} httpError.RestError{}
// @Failure 500 {object} httpError.RestError{}
// @Security ApiKeyAuth
// @Router /consumer-transactions [post]
func (h *Handler) Create(ctx context.Context, appCtx *app.RequestContext) {
	ctx, span := tracing.StartHttpServerTracerSpan(ctx, appCtx, "ConsumerTransactions.Create")
	defer span.Finish()
	h.metrics.CreateConsumerTransactionHTTPRequest.Inc()

	dto := &domain.CreateConsumerTransactionDTORequestParams{}

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

	idempotencyKey := string(appCtx.GetHeader(constants.XIdempotencyKey))

	res, err := h.usecase.Create(ctx, mapper.NewCreateRequestParams(dto), &idempotencyKey)
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
		Status:                consts.StatusOK,
		SuccessMetricsCounter: h.metrics.SuccessHTTPRequest,
		Data: http.IDResponse{
			Id: strconv.FormatInt(res, 10),
		},
		Message: "OK",
	})
}

// Find Consumer Transactions
// @Summary find list of Consumer-Transactions
// @Description find list of consumer transactions with the provided request parameters
// @Tags Consumer-Transactions
// @Accept json
// @Produce json
// @Param customer_id query string true "customer_id, determine the customer id"
// @Param page_size query string true "search page_size, determine the size of page / limit"
// @Param page_id query string true "search page_id, determine the number page / offset"
// @Success 200 {object} http.SuccessResponseDto{data=domain.ConsumerTransactionList}
// @Failure 400 {object} httpError.RestError{}
// @Failure 403 {object} httpError.RestError{}
// @Failure 404 {object} httpError.RestError{}
// @Failure 500 {object} httpError.RestError{}
// @Security ApiKeyAuth
// @Router /consumer-transactions [get]
func (h *Handler) Find(ctx context.Context, appCtx *app.RequestContext) {
	ctx, span := tracing.StartHttpServerTracerSpan(ctx, appCtx, "ConsumerTransactions.Find")
	defer span.Finish()
	h.metrics.FindConsumerTransactionHTTPRequest.Inc()

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
		SuccessMetricsCounter: h.metrics.SuccessHTTPRequest,
		Data:                  res,
		Message:               "OK",
	})
}

// Update Consumer Transaction
// @Summary Update an existing Consumer-Transaction by ID
// @Description Update an existing consumer transaction by the provided ID and request parameters
// @Tags Consumer-Transactions
// @Accept json
// @Produce json
// @Param id path string true "ID of the consumer transaction to update"
// @Param body body domain.UpdateConsumerTransactionDTORequestParams true "Request body for updating the consumer transaction"
// @Param x-idempotency-key header string false "Optional idempotency key to ensure the request is only processed once"
// @Success 200 {object} http.SuccessResponseDto{data=http.IDResponse} "Successfully updated consumer transaction"
// @Failure 400 {object} httpError.RestError{}
// @Failure 403 {object} httpError.RestError{}
// @Failure 404 {object} httpError.RestError{}
// @Failure 500 {object} httpError.RestError{}
// @Security ApiKeyAuth
// @Router /consumer-transactions/{id} [put]
func (h *Handler) Update(ctx context.Context, appCtx *app.RequestContext) {
	ctx, span := tracing.StartHttpServerTracerSpan(ctx, appCtx, "ConsumerTransactions.Update")
	defer span.Finish()
	h.metrics.UpdateConsumerTransactionHTTPRequest.Inc()

	dto := &domain.UpdateConsumerTransactionDTORequestParams{}

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

	idempotencyKey := string(appCtx.GetHeader(constants.XIdempotencyKey))

	res, err := h.usecase.Update(ctx, id, mapper.NewUpdateRequestParams(dto), &idempotencyKey)
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
		Status:                consts.StatusOK,
		Data:                  res,
		SuccessMetricsCounter: h.metrics.SuccessHTTPRequest,
		Message:               "OK",
	})
}

// FindByID Consumer Transaction
// @Summary Find Consumer-Transaction
// @Description Find consumer transaction by associated id
// @Tags Consumer-Transactions
// @Accept json
// @Produce json
// @Param id path string true "Any Associated id From your Source, this parameter is required"
// @Success 200 {object} http.SuccessResponseDto{data=domain.ConsumerTransaction}
// @Failure 400 {object} httpError.RestError{}
// @Failure 403 {object} httpError.RestError{}
// @Failure 404 {object} httpError.RestError{}
// @Failure 500 {object} httpError.RestError{}
// @Security ApiKeyAuth
// @Router /consumer-transactions/{id} [get]
func (h *Handler) FindByID(ctx context.Context, appCtx *app.RequestContext) {
	ctx, span := tracing.StartHttpServerTracerSpan(ctx, appCtx, "ConsumerTransactions.FindByID")
	defer span.Finish()
	h.metrics.FindByIDConsumerTransactionHTTPRequest.Inc()

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
