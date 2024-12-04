package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindByID(t *testing.T) {
	arg := createRandom(t)

	res, err := testStore.FindByID(context.TODO(), arg.ID)
	require.NoError(t, err)

	require.Equal(t, arg.ID, res.ID)
	require.Equal(t, arg.Nik, res.Nik)
	require.Equal(t, arg.FullName, res.FullName)
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
