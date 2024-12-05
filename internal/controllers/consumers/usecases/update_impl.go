package usecases

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/repository"
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

	if arg.FullName != nil {
		args.FullName = sql.NullString{
			String: *arg.FullName,
			Valid:  true,
		}
	}

	if arg.BirthPlace != nil {
		args.BirthPlace = sql.NullString{
			String: *arg.BirthPlace,
			Valid:  true,
		}
	}

	if arg.BirthDate != nil {
		birthDateLayoutStr := "2006-01-02"

		parsedBirthDate, err := time.Parse(birthDateLayoutStr, *arg.BirthDate)
		if err != nil {
			return 0, u.errorResponse(span, "parsedBirthDate.time.Parse.err", err)
		}

		args.BirthDate = sql.NullTime{
			Time:  parsedBirthDate,
			Valid: true,
		}
	}

	if arg.KTPPhoto != nil {
		args.KtpPhoto = sql.NullString{
			String: *arg.KTPPhoto,
			Valid:  true,
		}
	}

	if arg.SelfiePhoto != nil {
		args.SelfiePhoto = sql.NullString{
			String: *arg.SelfiePhoto,
			Valid:  true,
		}
	}

	if arg.Salary != nil {
		salaryFloat := decimal.NewFromFloat(*arg.Salary)

		args.Salary = decimal.NewNullDecimal(salaryFloat)
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

	if arg.FullName != nil || arg.BirthPlace != nil || arg.BirthDate != nil || arg.Salary != nil || arg.KTPPhoto != nil || arg.SelfiePhoto != nil {
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
