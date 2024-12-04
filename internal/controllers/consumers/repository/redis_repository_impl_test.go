package repository

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPut(t *testing.T) {
	createRandomCache(t)
}

func TestGet(t *testing.T) {
	arg := createRandomCache(t)

	res, err := testStore.Get(context.TODO(), strconv.FormatInt(arg.ID, 10))
	require.NoError(t, err)

	require.Equal(t, arg.ID, res.ID)
	require.Equal(t, arg.Nik, res.Nik)
	require.Equal(t, arg.FullName, res.FullName)
	require.Equal(t, arg.LegalName, res.LegalName)
	require.Equal(t, arg.BirthPlace, res.BirthPlace)
	require.Equal(t, arg.BirthDate, res.BirthDate)
	require.Equal(t, arg.Salary.Decimal.String(), res.Salary.Decimal.String())
	require.Equal(t, arg.KtpPhoto, res.KtpPhoto)
	require.Equal(t, arg.SelfiePhoto, res.SelfiePhoto)
	require.Equal(t, arg.CreatedAt, res.CreatedAt)
	require.Equal(t, arg.UpdatedAt, res.UpdatedAt)
	require.Equal(t, arg.UpdatedBy, res.UpdatedBy)
	require.Equal(t, arg.IsActivated, res.IsActivated)
	require.Equal(t, arg.IsActivatedAt, res.IsActivatedAt)
	require.Equal(t, arg.IsActivatedUpdatedAt, res.IsActivatedUpdatedAt)
}

func TestDel(t *testing.T) {
	arg := createRandomCache(t)

	res, err := testStore.Get(context.TODO(), strconv.FormatInt(arg.ID, 10))
	require.NoError(t, err)

	require.Equal(t, arg.ID, res.ID)
	require.Equal(t, arg.Nik, res.Nik)
	require.Equal(t, arg.FullName, res.FullName)
	require.Equal(t, arg.LegalName, res.LegalName)
	require.Equal(t, arg.BirthPlace, res.BirthPlace)
	require.Equal(t, arg.BirthDate, res.BirthDate)
	require.Equal(t, arg.Salary.Decimal.String(), res.Salary.Decimal.String())
	require.Equal(t, arg.KtpPhoto, res.KtpPhoto)
	require.Equal(t, arg.SelfiePhoto, res.SelfiePhoto)
	require.Equal(t, arg.CreatedAt, res.CreatedAt)
	require.Equal(t, arg.UpdatedAt, res.UpdatedAt)
	require.Equal(t, arg.UpdatedBy, res.UpdatedBy)
	require.Equal(t, arg.IsActivated, res.IsActivated)
	require.Equal(t, arg.IsActivatedAt, res.IsActivatedAt)
	require.Equal(t, arg.IsActivatedUpdatedAt, res.IsActivatedUpdatedAt)

	testStore.Del(context.TODO(), strconv.FormatInt(arg.ID, 10))

	res2, err := testStore.Get(context.TODO(), strconv.FormatInt(arg.ID, 10))
	require.Error(t, err)
	require.Empty(t, res2)

}

func createRandomCache(t *testing.T) *Consumer {
	arg := createRandom(t)

	testStore.Put(context.TODO(), strconv.FormatInt(arg.ID, 10), arg)

	return arg
}
