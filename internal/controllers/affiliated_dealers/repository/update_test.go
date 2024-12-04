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
		AffiliatedDealerName: sql.NullString{
			String: helper.RandomString(12),
			Valid:  true,
		},
		ID: arg.ID,
	}

	err := testStore.Update(context.TODO(), &updateArg)
	require.NoError(t, err)

	res, err := testStore.FindByID(context.TODO(), arg.ID)
	require.NoError(t, err)

	require.NotEqual(t, arg.AffiliatedDealerName, res.AffiliatedDealerName)
	require.Equal(t, updateArg.AffiliatedDealerName.String, res.AffiliatedDealerName)

	require.Equal(t, arg.ID, res.ID)
	require.Equal(t, arg.CreatedAt, res.CreatedAt)
	require.Equal(t, arg.UpdatedAt, res.UpdatedAt)
	require.Equal(t, arg.UpdatedBy, res.UpdatedBy)
	require.Equal(t, arg.IsActivated, res.IsActivated)
	require.Equal(t, arg.IsActivatedAt, res.IsActivatedAt)
	require.Equal(t, arg.IsActivatedUpdatedAt, res.IsActivatedUpdatedAt)
}
