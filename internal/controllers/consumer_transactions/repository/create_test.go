package repository

import (
	"context"
	"testing"
	"time"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	createRandom(t)
}

func createRandom(t *testing.T) *ConsumerTransaction {
	consumerArg := createRandomConsumer(t)
	affiliatedDealerArg := createRandomAffiliatedDealer(t)
	currentTime := time.Now().Format(time.RFC3339Nano)

	arg := &CreateParams{
		ConsumerID:         consumerArg,
		AffiliatedDealerID: affiliatedDealerArg,
		ContractNumber:     helper.RandomString(16),
		AdminFeeAmount:     decimal.NewNullDecimal(decimal.NewFromInt(helper.RandomInt(1000000, 10000000))),
		InstallmentAmount:  decimal.NewNullDecimal(decimal.NewFromInt(helper.RandomInt(1000000, 10000000))),
		OtrAmount:          decimal.NewNullDecimal(decimal.NewFromInt(helper.RandomInt(1000000, 10000000))),
		InterestRate:       decimal.NewNullDecimal(decimal.NewFromInt(helper.RandomInt(1, 30))),
		TransactionDate:    currentTime,
		CreatedAt:          currentTime,
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
