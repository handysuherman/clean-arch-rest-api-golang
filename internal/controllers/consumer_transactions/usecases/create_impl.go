package usecases

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/repository"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/shopspring/decimal"
)

func (u *usecaseImpl) Create(ctx context.Context, arg *domain.CreateRequestParams, idempotencyKey *string) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UsecaseImpl.Create")
	defer span.Finish()

	ctx, cancel := context.WithTimeout(ctx, u.cfg.Services.Internal.OperationTimeout)
	defer cancel()

	if idempotencyKey != nil {
		if idempotency_payload, err := u.repository.GetIdempotencyCreate(ctx, *idempotencyKey); err == nil && idempotency_payload != 0 {
			return idempotency_payload, nil
		}
	}

	contractNumber, err := helper.GenerateULID()
	if err != nil {
		return 0, tracing.TraceWithError(span, fmt.Errorf("create.helper.GenerateUlid.err: %v", err))
	}

	currentTime := time.Now().Format(time.RFC3339Nano)

	args := repository.CreateTxParams{
		Create: repository.CreateParams{
			ConsumerID:         arg.ConsumerID,
			ContractNumber:     contractNumber.String(),
			AffiliatedDealerID: arg.AffiliatedDealerID,
			TransactionDate:    currentTime,
			CreatedAt:          currentTime,
		},
	}

	if arg.AdminFeeAmount != nil {
		adminFeeAmount, err := decimal.NewFromString(*arg.AdminFeeAmount)
		if err != nil {
			return 0, u.errorResponse(span, "create.adminFeeAmount.decimal.NewFromString.err", err)
		}

		args.Create.AdminFeeAmount = decimal.NewNullDecimal(adminFeeAmount)
	}

	if arg.InstallmentAmount != nil {
		installmentAmount, err := decimal.NewFromString(*arg.InstallmentAmount)
		if err != nil {
			return 0, u.errorResponse(span, "create.installmentAmount.decimal.NewFromString.err", err)
		}

		args.Create.InstallmentAmount = decimal.NewNullDecimal(installmentAmount)
	}

	if arg.OtrAmount != nil {
		otrAmount, err := decimal.NewFromString(*arg.OtrAmount)
		if err != nil {
			return 0, u.errorResponse(span, "create.otrAmount.decimal.NewFromString.err", err)
		}

		args.Create.OtrAmount = decimal.NewNullDecimal(otrAmount)
	}

	if arg.InterestRate != nil {
		interestRate, err := decimal.NewFromString(*arg.InterestRate)
		if err != nil {
			return 0, u.errorResponse(span, "create.interestRate.decimal.NewFromString.err", err)
		}

		args.Create.InterestRate = decimal.NewNullDecimal(interestRate)
	}

	res, err := u.repository.CreateTx(ctx, &args)
	if err != nil {
		return 0, u.errorResponse(span, "create.u.repository.CreateTx.err: %v", err)
	}

	u.repository.Put(ctx, strconv.FormatInt(res.ConsumerTransaction.ID, 10), res.ConsumerTransaction)

	if idempotencyKey != nil {
		u.repository.PutIdempotencyCreate(ctx, *idempotencyKey, res.ConsumerTransaction.ID)
	}

	return res.ConsumerTransaction.ID, nil
}
