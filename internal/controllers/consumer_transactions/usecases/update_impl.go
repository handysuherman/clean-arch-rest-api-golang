package usecases

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/repository"
	"github.com/opentracing/opentracing-go"
	"github.com/shopspring/decimal"
)

func (u *usecaseImpl) Update(ctx context.Context, id int64, arg *domain.UpdateRequestParams, idempotencyKey *string) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UsecaseImpl.Update")
	defer span.Finish()

	ctx, cancel := context.WithTimeout(ctx, u.cfg.Services.Internal.OperationTimeout)
	defer cancel()

	if idempotencyKey != nil {
		if idempotency_payload, err := u.repository.GetIdempotencyUpdate(ctx, *idempotencyKey); err == nil && idempotency_payload != 0 {
			return idempotency_payload, nil
		}
	}

	args := repository.UpdateTxParams{
		Update: repository.UpdateParams{
			ID: id,
		},
	}

	if arg.AdminFeeAmount != nil {
		adminFeeAmount, err := decimal.NewFromString(*arg.AdminFeeAmount)
		if err != nil {
			return 0, u.errorResponse(span, "update.adminFeeAmount.decimal.NewFromString.err", err)
		}

		args.Update.AdminFeeAmount = decimal.NewNullDecimal(adminFeeAmount)
	}

	if arg.InstallmentAmount != nil {
		installmentAmount, err := decimal.NewFromString(*arg.InstallmentAmount)
		if err != nil {
			return 0, u.errorResponse(span, "update.installmentAmount.decimal.NewFromString.err", err)
		}

		args.Update.InstallmentAmount = decimal.NewNullDecimal(installmentAmount)
	}

	if arg.OtrAmount != nil {
		otrAmount, err := decimal.NewFromString(*arg.OtrAmount)
		if err != nil {
			return 0, u.errorResponse(span, "update.otrAmount.decimal.NewFromString.err", err)
		}

		args.Update.OtrAmount = decimal.NewNullDecimal(otrAmount)
	}

	if arg.InterestRate != nil {
		interestRate, err := decimal.NewFromString(*arg.InterestRate)
		if err != nil {
			return 0, u.errorResponse(span, "update.interestRate.decimal.NewFromString.err", err)
		}

		args.Update.InterestRate = decimal.NewNullDecimal(interestRate)
	}

	if args.Update.AdminFeeAmount.Valid || args.Update.InstallmentAmount.Valid || args.Update.OtrAmount.Valid || args.Update.InterestRate.Valid {
		args.Update.UpdatedAt = sql.NullString{
			String: time.Now().Format(time.RFC3339Nano),
			Valid:  true,
		}

		args.Update.UpdatedBy = sql.NullString{
			String: "system",
			Valid:  true,
		}
	}

	res, err := u.repository.UpdateTx(ctx, &args)
	if err != nil {
		return 0, u.errorResponse(span, "update.u.repository.UpdateTx.err: %v", err)
	}

	u.repository.Put(ctx, strconv.FormatInt(res.ConsumerTransaction.ID, 10), res.ConsumerTransaction)

	if idempotencyKey != nil {
		u.repository.PutIdempotencyUpdate(ctx, *idempotencyKey, res.ConsumerTransaction.ID)
	}

	return res.ConsumerTransaction.ID, nil
}
