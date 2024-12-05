package usecases

import (
	"context"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/mapper"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/repository"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

func (u *usecaseImpl) List(ctx context.Context, arg *domain.FetchParams) (*domain.AffiliatedDealerList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UsecaseImpl.FindByID")
	defer span.Finish()
	span.LogFields(log.String("search_text", arg.SearchText))

	ctx, cancel := context.WithTimeout(ctx, u.cfg.Services.Internal.OperationTimeout)
	defer cancel()

	searchText := "%" + arg.SearchText + "%"

	total, err := u.repository.CountList(ctx, searchText)
	if err != nil {
		return nil, u.errorResponse(span, "list.u.repository.countlist.err", err)
	}

	args := &repository.ListParams{
		AffiliatedDealerName: searchText,
		Limit:                int32(arg.Pagination.GetLimit()),
		Offset:               int32(arg.Pagination.GetOffset()),
	}

	res, err := u.repository.List(ctx, args)
	if err != nil {
		return nil, u.errorResponse(span, "list.u.repository.list.err", err)
	}

	return mapper.NewAffiliatedDealerList(
		res,
		total,
		arg.Pagination,
	), nil
}
