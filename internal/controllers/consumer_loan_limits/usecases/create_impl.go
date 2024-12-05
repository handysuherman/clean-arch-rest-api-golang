package usecases

import (
	"context"
	"strconv"
	"time"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_loan_limits/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_loan_limits/repository"
	"github.com/opentracing/opentracing-go"
	"github.com/shopspring/decimal"
)

func (u *usecaseImpl) Create(ctx context.Context, arg *domain.CreateRequestParams) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UsecaseImpl.Create")
	defer span.Finish()

	ctx, cancel := context.WithTimeout(ctx, u.cfg.Services.Internal.OperationTimeout)
	defer cancel()

	amount, err := decimal.NewFromString(arg.Amount)
	if err != nil {
		return 0, u.errorResponse(span, "amount.decimal.NewFromString.err", err)
	}

	args := repository.CreateParams{
		ConsumerID: arg.ConsumerID,
		Tenor:      arg.Tenor,
		Amount:     amount,
		CreatedAt:  time.Now().Format(time.RFC3339Nano),
	}

	res, err := u.repository.Create(ctx, &args)
	if err != nil {
		return 0, u.errorResponse(span, "u.repository.Create.err", err)
	}

	resultId, err := res.LastInsertId()
	if err != nil {
		return 0, u.errorResponse(span, "response.LastInsertId.err", err)
	}

	updated_res, err := u.repository.FindByID(ctx, resultId)
	if err != nil {
		return 0, u.errorResponse(span, "u.repository.FindByID", err)
	}

	u.repository.Put(ctx, strconv.FormatInt(updated_res.ID, 10), updated_res)

	return resultId, nil
}
