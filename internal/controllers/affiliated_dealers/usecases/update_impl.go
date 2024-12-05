package usecases

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/repository"
	"github.com/opentracing/opentracing-go"
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

	if arg.AffiliatedDealerName != nil {
		args.AffiliatedDealerName = sql.NullString{
			String: *arg.AffiliatedDealerName,
			Valid:  true,
		}
	}

	if arg.IsActivated != nil {
		args.IsActivated = sql.NullBool{
			Bool:  *arg.IsActivated,
			Valid: true,
		}

		args.IsActivatedUpdatedAt = sql.NullString{
			String: currentTime,
			Valid:  true,
		}
	}

	if arg.AffiliatedDealerName != nil {
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
