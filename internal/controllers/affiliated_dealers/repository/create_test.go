package repository

import (
	"context"
	"testing"
	"time"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	createRandom(t)
}

func createRandom(t *testing.T) *AffiliatedDealer {
	arg := &CreateParams{
		AffiliatedDealerName: helper.RandomString(100),
		CreatedAt:            time.Now().Format(time.RFC3339),
	}

	resultID, err := testStore.Create(context.TODO(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, resultID)

	id, err := resultID.LastInsertId()
	require.NoError(t, err)
	require.NotEmpty(t, id)

	res, err := testStore.FindByID(context.TODO(), id)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	return res
}
