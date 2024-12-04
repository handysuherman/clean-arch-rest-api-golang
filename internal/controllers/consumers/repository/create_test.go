package repository

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	createRandom(t)
}

func createRandom(t *testing.T) *Consumer {
	birthDateStr := fmt.Sprintf("%d-%02d-%02d", helper.RandomInt(2000, 2099), helper.RandomInt(1, 12), helper.RandomInt(1, 28))
	birthDateLayoutStr := "2006-01-02"

	parsedBirthDate, err := time.Parse(birthDateLayoutStr, birthDateStr)
	require.NoError(t, err)
	require.NotEmpty(t, parsedBirthDate)

	arg := CreateParams{
		Nik:      helper.RandomString(12),
		FullName: helper.RandomString(32),
		LegalName: sql.NullString{
			String: helper.RandomString(32),
			Valid:  true,
		},
		BirthPlace: sql.NullString{
			String: helper.RandomString(21),
			Valid:  true,
		},
		BirthDate: sql.NullTime{
			Time:  parsedBirthDate,
			Valid: true,
		},
		Salary: decimal.NewNullDecimal(decimal.NewFromInt(helper.RandomInt(1000000, 10000000))),
		KtpPhoto: sql.NullString{
			String: fmt.Sprintf("%s/%s.jpeg", helper.RandomUrl(), helper.RandomString(12)),
			Valid:  true,
		},
		SelfiePhoto: sql.NullString{
			String: fmt.Sprintf("%s/%s.jpeg", helper.RandomUrl(), helper.RandomString(12)),
			Valid:  true,
		},
		CreatedAt: time.Now().String(),
	}

	res, err := testStore.Create(context.TODO(), &arg)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	id, err := res.LastInsertId()
	require.NoError(t, err)
	require.NotEmpty(t, id)

	data, err := testStore.FindByID(context.TODO(), id)
	require.NoError(t, err)
	require.NotEmpty(t, data)

	return data
}
