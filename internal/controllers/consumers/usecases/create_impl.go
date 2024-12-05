package usecases

import (
	"context"
	"database/sql"
	"time"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/repository"
	"github.com/opentracing/opentracing-go"
	"github.com/shopspring/decimal"
)

func (u *usecaseImpl) Create(ctx context.Context, arg *domain.CreateRequestParams) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UsecaseImpl.Create")
	defer span.Finish()

	ctx, cancel := context.WithTimeout(ctx, u.cfg.Services.Internal.OperationTimeout)
	defer cancel()

	args := repository.CreateParams{
		Nik:       arg.Nik,
		FullName:  arg.FullName,
		CreatedAt: time.Now().Format(time.RFC3339Nano),
	}

	if arg.LegalName != nil {
		args.LegalName = sql.NullString{
			String: *arg.LegalName,
			Valid:  arg.LegalName != nil,
		}
	}

	if arg.BirthPlace != nil {
		args.BirthPlace = sql.NullString{
			String: *arg.BirthPlace,
			Valid:  arg.BirthPlace != nil,
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

	if arg.Salary != nil {
		salaryFloat := decimal.NewFromFloat(*arg.Salary)

		args.Salary = decimal.NewNullDecimal(salaryFloat)
	}

	if arg.KTPPhoto != nil {
		args.KtpPhoto = sql.NullString{
			String: *arg.KTPPhoto,
			Valid:  arg.KTPPhoto != nil,
		}
	}

	if arg.SelfiePhoto != nil {
		args.SelfiePhoto = sql.NullString{
			String: *arg.SelfiePhoto,
			Valid:  arg.SelfiePhoto != nil,
		}
	}

	res, err := u.repository.Create(ctx, &args)
	if err != nil {
		return 0, u.errorResponse(span, "u.repository.Create.err", err)
	}

	resultId, err := res.LastInsertId()
	if err != nil {
		return 0, u.errorResponse(span, "response.LastInsertId.err", err)
	}

	return resultId, nil
}
