package usecases

import (
	"context"
	"database/sql"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/mapper"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/repository"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

func (u *usecaseImpl) List(ctx context.Context, arg *domain.FetchParams) (*domain.ConsumerList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UsecaseImpl.FindByID")
	defer span.Finish()
	span.LogFields(log.String("search_text", arg.SearchText))

	ctx, cancel := context.WithTimeout(ctx, u.cfg.Services.Internal.OperationTimeout)
	defer cancel()

	searchText := "%" + arg.SearchText + "%"

	countArgs := &repository.CountListParams{
		FullName: searchText,
		LegalName: sql.NullString{
			String: searchText,
			Valid:  true,
		},
	}

	total, err := u.repository.CountList(ctx, countArgs)
	if err != nil {
		return nil, u.errorResponse(span, "list.u.repository.countlist.err", err)
	}

	args := &repository.ListParams{
		FullName:  countArgs.FullName,
		LegalName: countArgs.LegalName,
		Limit:     int32(arg.Pagination.GetLimit()),
		Offset:    int32(arg.Pagination.GetOffset()),
	}

	res, err := u.repository.List(ctx, args)
	if err != nil {
		return nil, u.errorResponse(span, "list.u.repository.list.err", err)
	}

	return mapper.NewConsumerList(
		res,
		total,
		arg.Pagination,
	), nil
}
