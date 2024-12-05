package usecases

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_loan_limits/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_loan_limits/repository"
	"github.com/opentracing/opentracing-go"
	"github.com/shopspring/decimal"
)

func (u *usecaseImpl) Update(ctx context.Context, id int64, arg *domain.UpdateRequestParams) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UsecaseImpl.Update")
	defer span.Finish()

	ctx, cancel := context.WithTimeout(ctx, u.cfg.Services.Internal.OperationTimeout)
	defer cancel()

	currentTime := time.Now().Format(time.RFC3339Nano)

	args := repository.UpdateParams{
		ID: id,
	}

	if arg.Tenor != nil {
		args.Tenor = sql.NullInt16{
			Int16: *arg.Tenor,
			Valid: true,
		}
	}

	if arg.Amount != nil {
		amount, err := decimal.NewFromString(*arg.Amount)
		if err != nil {
			return 0, u.errorResponse(span, "amount.decimal.NewFromString.err", err)
		}

		args.Amount = decimal.NewNullDecimal(amount)
	}

	if arg.Tenor != nil || arg.Amount != nil {
		args.UpdatedAt = sql.NullString{
			String: currentTime,
			Valid:  true,
		}

		args.UpdatedBy = sql.NullString{
			String: "system",
			Valid:  true,
		}
	}

	err := u.repository.Update(ctx, &args)
	if err != nil {
		return 0, u.errorResponse(span, "u.repository.Update", err)
	}

	updated_res, err := u.repository.FindByID(ctx, args.ID)
	if err != nil {
		return 0, u.errorResponse(span, "u.repository.FindByID", err)
	}

	u.repository.Put(ctx, strconv.FormatInt(updated_res.ID, 10), updated_res)

	return updated_res.ID, nil
}
