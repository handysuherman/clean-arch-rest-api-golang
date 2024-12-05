package repository

import (
	"context"
	"testing"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestUpdateTx(t *testing.T) {
	arg := createRandom(t)

	updateArg := UpdateParams{
		AdminFeeAmount: decimal.NewNullDecimal(decimal.NewFromInt(helper.RandomInt(1000000, 10000000))),
		ID:             arg.ID,
	}

	res, err := testStore.UpdateTx(context.TODO(), &UpdateTxParams{Update: updateArg})
	require.NoError(t, err)

	require.NotEqual(t, arg.AdminFeeAmount.Decimal.String(), res.ConsumerTransaction.AdminFeeAmount.Decimal.String())
	require.Equal(t, updateArg.AdminFeeAmount.Decimal.String(), res.ConsumerTransaction.AdminFeeAmount.Decimal.String())

	require.Equal(t, arg.ID, res.ConsumerTransaction.ID)
	require.Equal(t, arg.ConsumerID, res.ConsumerTransaction.ConsumerID)
	require.Equal(t, arg.ContractNumber, res.ConsumerTransaction.ContractNumber)
	require.Equal(t, arg.InstallmentAmount.Decimal.String(), res.ConsumerTransaction.InstallmentAmount.Decimal.String())
	require.Equal(t, arg.OtrAmount.Decimal.String(), res.ConsumerTransaction.OtrAmount.Decimal.String())
	require.Equal(t, arg.InterestRate.Decimal.String(), res.ConsumerTransaction.InterestRate.Decimal.String())
	require.Equal(t, arg.TransactionDate, res.ConsumerTransaction.TransactionDate)
	require.Equal(t, arg.AffiliatedDealerID, res.ConsumerTransaction.AffiliatedDealerID)
	require.Equal(t, arg.CreatedAt, res.ConsumerTransaction.CreatedAt)
	require.Equal(t, arg.UpdatedAt, res.ConsumerTransaction.UpdatedAt)
	require.Equal(t, arg.UpdatedBy, res.ConsumerTransaction.UpdatedBy)
}
