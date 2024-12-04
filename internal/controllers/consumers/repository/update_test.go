package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	arg := createRandom(t)

	updateArg := UpdateParams{
		FullName: sql.NullString{
			String: helper.RandomString(12),
			Valid:  true,
		},
		ID: arg.ID,
	}

	err := testStore.Update(context.TODO(), &updateArg)
	require.NoError(t, err)

	res, err := testStore.FindByID(context.TODO(), arg.ID)
	require.NoError(t, err)

	require.NotEqual(t, arg.FullName, res.FullName)
	require.Equal(t, updateArg.FullName.String, res.FullName)

	require.Equal(t, arg.ID, res.ID)
	require.Equal(t, arg.Nik, res.Nik)
	require.Equal(t, arg.LegalName, res.LegalName)
	require.Equal(t, arg.BirthPlace, res.BirthPlace)
	require.Equal(t, arg.BirthDate, res.BirthDate)
	require.Equal(t, arg.Salary, res.Salary)
	require.Equal(t, arg.KtpPhoto, res.KtpPhoto)
	require.Equal(t, arg.SelfiePhoto, res.SelfiePhoto)
	require.Equal(t, arg.CreatedAt, res.CreatedAt)
	require.Equal(t, arg.UpdatedAt, res.UpdatedAt)
	require.Equal(t, arg.UpdatedBy, res.UpdatedBy)
	require.Equal(t, arg.IsActivated, res.IsActivated)
	require.Equal(t, arg.IsActivatedAt, res.IsActivatedAt)
	require.Equal(t, arg.IsActivatedUpdatedAt, res.IsActivatedUpdatedAt)
}
