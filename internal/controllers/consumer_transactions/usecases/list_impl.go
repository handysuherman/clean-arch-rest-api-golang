package usecases

import (
	"context"
	"strconv"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/mapper"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/repository"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

func (u *usecaseImpl) List(ctx context.Context, arg *domain.FetchParams) (*domain.ConsumerTransactionList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UsecaseImpl.FindByID")
	defer span.Finish()
	span.LogFields(log.String("ConsumerID", strconv.FormatInt(arg.ConsumerID, 10)))

	ctx, cancel := context.WithTimeout(ctx, u.cfg.Services.Internal.OperationTimeout)
	defer cancel()

	total, err := u.repository.CountList(ctx, arg.ConsumerID)
	if err != nil {
		return nil, u.errorResponse(span, "list.u.repository.countlist.err", err)
	}

	args := &repository.ListParams{
		ConsumerID: arg.ConsumerID,
		Limit:      int32(arg.Pagination.GetLimit()),
		Offset:     int32(arg.Pagination.GetOffset()),
	}

	res, err := u.repository.List(ctx, args)
	if err != nil {
		return nil, u.errorResponse(span, "list.u.repository.list.err", err)
	}

	return mapper.NewConsumerTransactionList(
		res,
		total,
		arg.Pagination,
	), nil
}
