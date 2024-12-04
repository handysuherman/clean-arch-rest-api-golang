package repository

import (
	"context"
	"testing"
	"time"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
	"github.com/stretchr/testify/require"
)

func TestCreateAffiliatedDealer(t *testing.T) {
	createRandomAffiliatedDealer(t)
}

func createRandomAffiliatedDealer(t *testing.T) int64 {
	arg := &CreateAffiliatedDealerParams{
		AffiliatedDealerName: helper.RandomString(100),
		CreatedAt:            time.Now().Format(time.RFC3339),
	}

	resultID, err := testStore.CreateAffiliatedDealer(context.TODO(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, resultID)

	id, err := resultID.LastInsertId()
	require.NoError(t, err)
	require.NotEmpty(t, id)

	return id
}
