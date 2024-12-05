package repository

import (
	"context"
	"strconv"
	"testing"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
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
	require.Equal(t, arg.ConsumerID, res.ConsumerID)
	require.Equal(t, arg.ContractNumber, res.ContractNumber)
	require.Equal(t, arg.AdminFeeAmount.Decimal.String(), res.AdminFeeAmount.Decimal.String())
	require.Equal(t, arg.InstallmentAmount.Decimal.String(), res.InstallmentAmount.Decimal.String())
	require.Equal(t, arg.OtrAmount.Decimal.String(), res.OtrAmount.Decimal.String())
	require.Equal(t, arg.InterestRate.Decimal.String(), res.InterestRate.Decimal.String())
	require.Equal(t, arg.TransactionDate, res.TransactionDate)
	require.Equal(t, arg.AffiliatedDealerID, res.AffiliatedDealerID)
	require.Equal(t, arg.CreatedAt, res.CreatedAt)
	require.Equal(t, arg.UpdatedAt, res.UpdatedAt)
	require.Equal(t, arg.UpdatedBy, res.UpdatedBy)
}

func TestDel(t *testing.T) {
	arg := createRandomCache(t)

	res, err := testStore.Get(context.TODO(), strconv.FormatInt(arg.ID, 10))
	require.NoError(t, err)

	require.Equal(t, arg.ID, res.ID)
	require.Equal(t, arg.ConsumerID, res.ConsumerID)
	require.Equal(t, arg.ContractNumber, res.ContractNumber)
	require.Equal(t, arg.AdminFeeAmount.Decimal.String(), res.AdminFeeAmount.Decimal.String())
	require.Equal(t, arg.InstallmentAmount.Decimal.String(), res.InstallmentAmount.Decimal.String())
	require.Equal(t, arg.OtrAmount.Decimal.String(), res.OtrAmount.Decimal.String())
	require.Equal(t, arg.InterestRate.Decimal.String(), res.InterestRate.Decimal.String())
	require.Equal(t, arg.TransactionDate, res.TransactionDate)
	require.Equal(t, arg.AffiliatedDealerID, res.AffiliatedDealerID)
	require.Equal(t, arg.CreatedAt, res.CreatedAt)
	require.Equal(t, arg.UpdatedAt, res.UpdatedAt)
	require.Equal(t, arg.UpdatedBy, res.UpdatedBy)

	testStore.Del(context.TODO(), strconv.FormatInt(arg.ID, 10))

	res2, err := testStore.Get(context.TODO(), strconv.FormatInt(arg.ID, 10))
	require.Error(t, err)
	require.Empty(t, res2)

}

func TestPutIdempotencyCreate(t *testing.T) {
	testStore.PutIdempotencyCreate(context.TODO(), helper.RandomString(12), 1)
}

func TestGetIdempotencyCreate(t *testing.T) {
	key := helper.RandomString(12)
	testStore.PutIdempotencyCreate(context.TODO(), key, 1)

	res, err := testStore.GetIdempotencyCreate(context.TODO(), key)
	require.NoError(t, err)

	require.Equal(t, int64(1), res)
}

func TestDelIdempotencyCreate(t *testing.T) {
	key := helper.RandomString(12)
	testStore.PutIdempotencyCreate(context.TODO(), key, 1)

	res, err := testStore.GetIdempotencyCreate(context.TODO(), key)
	require.NoError(t, err)
	require.Equal(t, int64(1), res)

	testStore.DelIdempotencyCreate(context.TODO(), key)

	res2, err := testStore.GetIdempotencyCreate(context.TODO(), key)
	require.Error(t, err)
	require.Empty(t, res2)

}

func TestPutIdempotencyUpdate(t *testing.T) {
	testStore.PutIdempotencyUpdate(context.TODO(), helper.RandomString(12), 1)
}

func TestGetIdempotencyUpdate(t *testing.T) {
	key := helper.RandomString(12)
	testStore.PutIdempotencyUpdate(context.TODO(), key, 1)

	res, err := testStore.GetIdempotencyUpdate(context.TODO(), key)
	require.NoError(t, err)

	require.Equal(t, int64(1), res)
}

func TestDelIdempotencyUpdate(t *testing.T) {
	key := helper.RandomString(12)
	testStore.PutIdempotencyUpdate(context.TODO(), key, 1)

	res, err := testStore.GetIdempotencyUpdate(context.TODO(), key)
	require.NoError(t, err)
	require.Equal(t, int64(1), res)

	testStore.DelIdempotencyUpdate(context.TODO(), key)

	res2, err := testStore.GetIdempotencyUpdate(context.TODO(), key)
	require.Error(t, err)
	require.Empty(t, res2)

}

func createRandomCache(t *testing.T) *ConsumerTransaction {
	arg := createRandom(t)

	testStore.Put(context.TODO(), strconv.FormatInt(arg.ID, 10), arg)

	return arg
}
