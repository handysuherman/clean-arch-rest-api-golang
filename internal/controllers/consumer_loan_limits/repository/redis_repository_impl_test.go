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
	require.Equal(t, arg.ConsumerID, res.ConsumerID)
	require.Equal(t, arg.Tenor, res.Tenor)
	require.Equal(t, arg.Amount.String(), res.Amount.String())
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
	require.Equal(t, arg.Tenor, res.Tenor)
	require.Equal(t, arg.Amount.String(), res.Amount.String())
	require.Equal(t, arg.CreatedAt, res.CreatedAt)
	require.Equal(t, arg.UpdatedAt, res.UpdatedAt)
	require.Equal(t, arg.UpdatedBy, res.UpdatedBy)

	testStore.Del(context.TODO(), strconv.FormatInt(arg.ID, 10))

	res2, err := testStore.Get(context.TODO(), strconv.FormatInt(arg.ID, 10))
	require.Error(t, err)
	require.Empty(t, res2)

}

func createRandomCache(t *testing.T) *ConsumerLoanLimit {
	arg := createRandom(t)

	testStore.Put(context.TODO(), strconv.FormatInt(arg.ID, 10), arg)

	return arg
}
