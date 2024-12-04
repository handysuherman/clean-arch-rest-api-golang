package repository

import (
	"context"
	"testing"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	arg := createRandom(t)

	updateArg := UpdateParams{
		AdminFeeAmount: decimal.NewNullDecimal(decimal.NewFromInt(helper.RandomInt(1000000, 10000000))),
		ID:             arg.ID,
	}

	err := testStore.Update(context.TODO(), &updateArg)
	require.NoError(t, err)

	res, err := testStore.FindByID(context.TODO(), arg.ID)
	require.NoError(t, err)

	require.NotEqual(t, arg.AdminFeeAmount.Decimal.String(), res.AdminFeeAmount.Decimal.String())
	require.Equal(t, updateArg.AdminFeeAmount.Decimal.String(), res.AdminFeeAmount.Decimal.String())

	require.Equal(t, arg.ID, res.ID)
	require.Equal(t, arg.ConsumerID, res.ConsumerID)
	require.Equal(t, arg.ContractNumber, res.ContractNumber)
	require.Equal(t, arg.InstallmentAmount.Decimal.String(), res.InstallmentAmount.Decimal.String())
	require.Equal(t, arg.OtrAmount.Decimal.String(), res.OtrAmount.Decimal.String())
	require.Equal(t, arg.InterestRate.Decimal.String(), res.InterestRate.Decimal.String())
	require.Equal(t, arg.TransactionDate, res.TransactionDate)
	require.Equal(t, arg.AffiliatedDealerID, res.AffiliatedDealerID)
	require.Equal(t, arg.CreatedAt, res.CreatedAt)
	require.Equal(t, arg.UpdatedAt, res.UpdatedAt)
	require.Equal(t, arg.UpdatedBy, res.UpdatedBy)
}
