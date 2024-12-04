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
	require.Equal(t, arg.ConsumerID, res.ConsumerID)
	require.Equal(t, arg.Tenor, res.Tenor)
	require.Equal(t, arg.Amount.String(), res.Amount.String())
	require.Equal(t, arg.CreatedAt, res.CreatedAt)
	require.Equal(t, arg.UpdatedAt, res.UpdatedAt)
	require.Equal(t, arg.UpdatedBy, res.UpdatedBy)
}
