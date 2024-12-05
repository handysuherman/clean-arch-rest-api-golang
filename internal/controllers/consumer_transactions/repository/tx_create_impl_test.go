package repository

import (
	"context"
	"testing"
	"time"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestCreateTx(t *testing.T) {
	createRandomTx(t)
}

func createRandomTx(t *testing.T) *ConsumerTransaction {
	consumerArg := createRandomConsumer(t)
	affiliatedDealerArg := createRandomAffiliatedDealer(t)
	currentTime := time.Now().Format(time.RFC3339Nano)

	arg := CreateParams{
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

	res, err := testStore.CreateTx(context.TODO(), &CreateTxParams{Create: arg})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, arg.ConsumerID, res.ConsumerTransaction.ConsumerID)
	require.Equal(t, arg.ContractNumber, res.ConsumerTransaction.ContractNumber)
	require.Equal(t, arg.AdminFeeAmount.Decimal.String(), res.ConsumerTransaction.AdminFeeAmount.Decimal.String())
	require.Equal(t, arg.InstallmentAmount.Decimal.String(), res.ConsumerTransaction.InstallmentAmount.Decimal.String())
	require.Equal(t, arg.OtrAmount.Decimal.String(), res.ConsumerTransaction.OtrAmount.Decimal.String())
	require.Equal(t, arg.InterestRate.Decimal.String(), res.ConsumerTransaction.InterestRate.Decimal.String())
	require.Equal(t, arg.TransactionDate, res.ConsumerTransaction.TransactionDate)
	require.Equal(t, arg.AffiliatedDealerID, res.ConsumerTransaction.AffiliatedDealerID)
	require.Equal(t, arg.CreatedAt, res.ConsumerTransaction.CreatedAt)

	return res.ConsumerTransaction
}
