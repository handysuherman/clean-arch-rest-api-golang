package usecases

import (
	"context"
	"strconv"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_loan_limits/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_loan_limits/mapper"
	"github.com/opentracing/opentracing-go"
)

func (u *usecaseImpl) FindByID(ctx context.Context, id int64) (*domain.ConsumerLoanLimit, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UsecaseImpl.FindByID")
	defer span.Finish()

	ctx, cancel := context.WithTimeout(ctx, u.cfg.Services.Internal.OperationTimeout)
	defer cancel()

	if payload, err := u.repository.Get(ctx, strconv.FormatInt(id, 10)); err == nil && payload != nil {
		return mapper.ToDTO(payload), nil
	}

	res, err := u.repository.FindByID(ctx, id)
	if err != nil {
		return nil, u.errorResponse(span, "find_by_id.u.repository.FindByID.err", err)
	}

	u.repository.Put(ctx, strconv.FormatInt(id, 10), res)
	return mapper.ToDTO(res), nil
}
