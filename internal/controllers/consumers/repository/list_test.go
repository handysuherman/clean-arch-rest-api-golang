package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	arg := createRandom(t)

	listArg := ListParams{
		FullName:  arg.FullName,
		LegalName: arg.LegalName,
		Limit:     10,
		Offset:    0,
	}

	res, err := testStore.List(context.TODO(), &listArg)
	require.NoError(t, err)
	require.NotEmpty(t, res)
}
