package usecases

import (
	"context"
	"strconv"
	"time"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/repository"
	"github.com/opentracing/opentracing-go"
)

func (u *usecaseImpl) Create(ctx context.Context, arg *domain.CreateRequestParams) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UsecaseImpl.Create")
	defer span.Finish()

	ctx, cancel := context.WithTimeout(ctx, u.cfg.Services.Internal.OperationTimeout)
	defer cancel()

	args := repository.CreateParams{
		AffiliatedDealerName: arg.AffiliatedDealerName,
		CreatedAt:            time.Now().Format(time.RFC3339Nano),
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
