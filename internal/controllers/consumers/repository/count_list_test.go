package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountList(t *testing.T) {
	arg := createRandom(t)

	countListArg := CountListParams{
		FullName:  arg.FullName,
		LegalName: arg.LegalName,
	}

	res, err := testStore.CountList(context.TODO(), &countListArg)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.True(t, res == 1)
}
