package usecases

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/repository"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/repository/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_MOCK_CREATE(t *testing.T) {
	mockArgs := createRandom(t)

	testCases := []struct {
		tname          string
		body           *domain.CreateRequestParams
		idempotencyKey *string
		stubs          func(store *mock.MockRepository)
		checkResponse  func(t *testing.T, res int64, err error)
	}{
		{
			tname:          "OK_IDEMPOTENCY_HIT",
			body:           mockArgs.createParams,
			idempotencyKey: &mockArgs.idempotencyKey,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().GetIdempotencyCreate(gomock.Any(), gomock.Eq(mockArgs.idempotencyKey)).Times(1).Return(int64(1), nil)

				store.EXPECT().CreateTx(gomock.Any(), EqCreateTxParamsMatcher(&repository.CreateTxParams{Create: *mockArgs.createRepoParams})).Times(0)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(mockArgs.repoResponse.ID), gomock.Eq(mockArgs.repoResponse)).Times(0)

				store.EXPECT().PutIdempotencyCreate(gomock.Any(), gomock.Eq(mockArgs.idempotencyKey), gomock.Eq(mockArgs.repoResponse.ID)).Times(0)
			},
			checkResponse: func(t *testing.T, res int64, err error) {
				require.NoError(t, err)
			},
		},
		{
			tname:          "OK_IDEMPOTENCY_MISS",
			body:           mockArgs.createParams,
			idempotencyKey: &mockArgs.idempotencyKey,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().GetIdempotencyCreate(gomock.Any(), gomock.Eq(mockArgs.idempotencyKey)).Times(1).Return(int64(0), errors.New("not found"))

				store.EXPECT().CreateTx(gomock.Any(), EqCreateTxParamsMatcher(&repository.CreateTxParams{Create: *mockArgs.createRepoParams})).Times(1).Return(repository.CreateTxResult{ConsumerTransaction: mockArgs.repoResponse}, nil)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(mockArgs.repoResponse.ID), gomock.Eq(mockArgs.repoResponse)).Times(1)

				store.EXPECT().PutIdempotencyCreate(gomock.Any(), gomock.Eq(mockArgs.idempotencyKey), gomock.Eq(mockArgs.repoResponse.ID)).Times(1)
			},
			checkResponse: func(t *testing.T, res int64, err error) {
				require.NoError(t, err)
			},
		},
		{
			tname:          "ERR_CREATE_TX_ROLLBACK",
			body:           mockArgs.createParams,
			idempotencyKey: &mockArgs.idempotencyKey,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().GetIdempotencyCreate(gomock.Any(), gomock.Eq(mockArgs.idempotencyKey)).Times(1).Return(int64(0), errors.New("not found"))

				store.EXPECT().CreateTx(gomock.Any(), EqCreateTxParamsMatcher(&repository.CreateTxParams{Create: *mockArgs.createRepoParams})).Times(1).Return(repository.CreateTxResult{}, sql.ErrTxDone)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(mockArgs.repoResponse.ID), gomock.Eq(mockArgs.repoResponse)).Times(0)

				store.EXPECT().PutIdempotencyCreate(gomock.Any(), gomock.Eq(mockArgs.idempotencyKey), gomock.Eq(mockArgs.repoResponse.ID)).Times(0)
			},
			checkResponse: func(t *testing.T, res int64, err error) {
				require.Error(t, err)
				require.Empty(t, res)
			},
		},
		{
			tname:          "ERR_CREATE_TX_INTERNAL_SERVER_ERROR",
			body:           mockArgs.createParams,
			idempotencyKey: &mockArgs.idempotencyKey,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().GetIdempotencyCreate(gomock.Any(), gomock.Eq(mockArgs.idempotencyKey)).Times(1).Return(int64(0), errors.New("not found"))

				store.EXPECT().CreateTx(gomock.Any(), EqCreateTxParamsMatcher(&repository.CreateTxParams{Create: *mockArgs.createRepoParams})).Times(1).Return(repository.CreateTxResult{}, sql.ErrConnDone)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(mockArgs.repoResponse.ID), gomock.Eq(mockArgs.repoResponse)).Times(0)

				store.EXPECT().PutIdempotencyCreate(gomock.Any(), gomock.Eq(mockArgs.idempotencyKey), gomock.Eq(mockArgs.repoResponse.ID)).Times(0)
			},
			checkResponse: func(t *testing.T, res int64, err error) {
				require.Error(t, err)
				require.Empty(t, res)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.tname, func(t *testing.T) {
			storeCtrl := gomock.NewController(t)
			defer storeCtrl.Finish()

			store := mock.NewMockRepository(storeCtrl)

			u := New(tlog, cfg, store)
			tc.stubs(store)

			actualBody, actualError := u.Create(context.TODO(), tc.body, tc.idempotencyKey)
			tc.checkResponse(t, actualBody, actualError)
		})
	}
}

type eqCreateTxParamsMatcher struct {
	arg *repository.CreateTxParams
}

func EqCreateTxParamsMatcher(arg *repository.CreateTxParams) gomock.Matcher {
	return &eqCreateTxParamsMatcher{arg: arg}
}

func (ex *eqCreateTxParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(*repository.CreateTxParams)
	if !ok {
		return false
	}

	ex.arg.Create.ConsumerID = arg.Create.ConsumerID
	ex.arg.Create.ContractNumber = arg.Create.ContractNumber
	ex.arg.Create.AdminFeeAmount = arg.Create.AdminFeeAmount
	ex.arg.Create.InstallmentAmount = arg.Create.InstallmentAmount
	ex.arg.Create.OtrAmount = arg.Create.OtrAmount
	ex.arg.Create.InterestRate = arg.Create.InterestRate
	ex.arg.Create.TransactionDate = arg.Create.TransactionDate
	ex.arg.Create.CreatedAt = arg.Create.CreatedAt
	ex.arg.Create.AffiliatedDealerID = arg.Create.AffiliatedDealerID

	if ex.arg.Create.ConsumerID == int64(0) {
		return false
	}

	if ex.arg.Create.ContractNumber == "" {
		return false
	}

	if ex.arg.Create.AffiliatedDealerID == int64(0) {
		return false
	}

	_, err := time.Parse(time.RFC3339Nano, ex.arg.Create.TransactionDate)
	if err != nil {
		return false
	}

	_, err = time.Parse(time.RFC3339Nano, ex.arg.Create.CreatedAt)
	if err != nil {
		return false
	}

	return reflect.DeepEqual(ex.arg, arg)
}

func (ex *eqCreateTxParamsMatcher) String() string {
	var errMsg string

	if ex.arg.Create.ConsumerID == int64(0) {
		errMsg += fmt.Sprint("consumer id should not be empty or zero")
	}

	if ex.arg.Create.ContractNumber == "" {
		errMsg += fmt.Sprint("contract number should not be empty")
	}

	if ex.arg.Create.AffiliatedDealerID == int64(0) {
		errMsg += fmt.Sprint("affiliate dealer id should not be empty")
	}

	_, err := time.Parse(time.RFC3339Nano, ex.arg.Create.TransactionDate)
	if err != nil {
		errMsg += fmt.Sprintf("transaction date doesnt reflect the time.RFC3339Nano Layout")
	}

	_, err = time.Parse(time.RFC3339Nano, ex.arg.Create.CreatedAt)
	if err != nil {
		errMsg += fmt.Sprint("created at doesnt reflect the time.RFC3339Nano layout")
	}

	return errMsg + fmt.Sprintf("matches arg: %v", ex.arg)
}
