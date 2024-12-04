package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountList(t *testing.T) {
	arg := createRandom(t)

	res, err := testStore.CountList(context.TODO(), arg.AffiliatedDealerName)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.True(t, res == 1)
}
