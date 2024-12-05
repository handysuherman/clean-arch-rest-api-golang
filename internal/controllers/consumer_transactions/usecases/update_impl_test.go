package usecases

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/repository"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/repository/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_MOCK_UPDATE(t *testing.T) {
	mockArgs := createRandom(t)

	testCases := []struct {
		tname          string
		id             int64
		body           *domain.UpdateRequestParams
		idempotencyKey *string
		stubs          func(store *mock.MockRepository)
		checkResponse  func(t *testing.T, res int64, err error)
	}{
		{
			tname:          "OK_IDEMPOTENCY_HIT",
			id:             mockArgs.updateRepoParams.ID,
			body:           mockArgs.updateParams,
			idempotencyKey: &mockArgs.idempotencyKey,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().GetIdempotencyUpdate(gomock.Any(), gomock.Eq(mockArgs.idempotencyKey)).Times(1).Return(int64(1), nil)

				store.EXPECT().UpdateTx(gomock.Any(), EqUpdateTxParamsMatcher(&repository.UpdateTxParams{Update: *mockArgs.updateRepoParams})).Times(0)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.repoResponse.ID, 10)), gomock.Eq(mockArgs.repoResponse)).Times(0)

				store.EXPECT().PutIdempotencyUpdate(gomock.Any(), gomock.Eq(mockArgs.idempotencyKey), gomock.Eq(mockArgs.repoResponse.ID)).Times(0)
			},
			checkResponse: func(t *testing.T, res int64, err error) {
				require.NoError(t, err)
			},
		},
		{
			tname:          "OK_IDEMPOTENCY_MISS",
			id:             mockArgs.updateRepoParams.ID,
			body:           mockArgs.updateParams,
			idempotencyKey: &mockArgs.idempotencyKey,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().GetIdempotencyUpdate(gomock.Any(), gomock.Eq(mockArgs.idempotencyKey)).Times(1).Return(int64(0), errors.New("not found"))

				store.EXPECT().UpdateTx(gomock.Any(), EqUpdateTxParamsMatcher(&repository.UpdateTxParams{Update: *mockArgs.updateRepoParams})).Times(1).Return(repository.UpdateTxResult{ConsumerTransaction: mockArgs.repoResponse}, nil)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.repoResponse.ID, 10)), gomock.Eq(mockArgs.repoResponse)).Times(1)

				store.EXPECT().PutIdempotencyUpdate(gomock.Any(), gomock.Eq(mockArgs.idempotencyKey), gomock.Eq(mockArgs.repoResponse.ID)).Times(1)
			},
			checkResponse: func(t *testing.T, res int64, err error) {
				require.NoError(t, err)
			},
		},
		{
			tname:          "OK_IDEMPOTENCY_NIL",
			id:             mockArgs.updateRepoParams.ID,
			body:           mockArgs.updateParams,
			idempotencyKey: nil,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().GetIdempotencyUpdate(gomock.Any(), gomock.Eq(mockArgs.idempotencyKey)).Times(0)

				store.EXPECT().UpdateTx(gomock.Any(), EqUpdateTxParamsMatcher(&repository.UpdateTxParams{Update: *mockArgs.updateRepoParams})).Times(1).Return(repository.UpdateTxResult{ConsumerTransaction: mockArgs.repoResponse}, nil)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.repoResponse.ID, 10)), gomock.Eq(mockArgs.repoResponse)).Times(1)

				store.EXPECT().PutIdempotencyUpdate(gomock.Any(), gomock.Eq(mockArgs.idempotencyKey), gomock.Eq(mockArgs.repoResponse.ID)).Times(0)
			},
			checkResponse: func(t *testing.T, res int64, err error) {
				require.NoError(t, err)
			},
		},
		{
			tname:          "ERR_CREATE_TX_ROLLBACK",
			id:             mockArgs.updateRepoParams.ID,
			body:           mockArgs.updateParams,
			idempotencyKey: &mockArgs.idempotencyKey,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().GetIdempotencyUpdate(gomock.Any(), gomock.Eq(mockArgs.idempotencyKey)).Times(1).Return(int64(0), errors.New("not found"))

				store.EXPECT().UpdateTx(gomock.Any(), EqUpdateTxParamsMatcher(&repository.UpdateTxParams{Update: *mockArgs.updateRepoParams})).Times(1).Return(repository.UpdateTxResult{}, sql.ErrTxDone)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.repoResponse.ID, 10)), gomock.Eq(mockArgs.repoResponse)).Times(0)

				store.EXPECT().PutIdempotencyUpdate(gomock.Any(), gomock.Eq(mockArgs.idempotencyKey), gomock.Eq(mockArgs.repoResponse.ID)).Times(0)
			},
			checkResponse: func(t *testing.T, res int64, err error) {
				require.Error(t, err)
				require.Empty(t, res)
			},
		},
		{
			tname:          "ERR_CREATE_TX_INTERNAL_SERVER_ERROR",
			id:             mockArgs.updateRepoParams.ID,
			body:           mockArgs.updateParams,
			idempotencyKey: &mockArgs.idempotencyKey,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().GetIdempotencyUpdate(gomock.Any(), gomock.Eq(mockArgs.idempotencyKey)).Times(1).Return(int64(0), errors.New("not found"))

				store.EXPECT().UpdateTx(gomock.Any(), EqUpdateTxParamsMatcher(&repository.UpdateTxParams{Update: *mockArgs.updateRepoParams})).Times(1).Return(repository.UpdateTxResult{}, sql.ErrConnDone)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.repoResponse.ID, 10)), gomock.Eq(mockArgs.repoResponse)).Times(0)

				store.EXPECT().PutIdempotencyUpdate(gomock.Any(), gomock.Eq(mockArgs.idempotencyKey), gomock.Eq(mockArgs.repoResponse.ID)).Times(0)
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

			actualBody, actualError := u.Update(context.TODO(), tc.id, tc.body, tc.idempotencyKey)
			tc.checkResponse(t, actualBody, actualError)
		})
	}
}

type eqUpdateTxParamsMatcher struct {
	arg *repository.UpdateTxParams
}

func EqUpdateTxParamsMatcher(arg *repository.UpdateTxParams) gomock.Matcher {
	return &eqUpdateTxParamsMatcher{arg: arg}
}

func (ex *eqUpdateTxParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(*repository.UpdateTxParams)
	if !ok {
		return false
	}

	ex.arg.Update.AdminFeeAmount = arg.Update.AdminFeeAmount
	ex.arg.Update.InstallmentAmount = arg.Update.InstallmentAmount
	ex.arg.Update.OtrAmount = arg.Update.OtrAmount
	ex.arg.Update.InterestRate = arg.Update.InterestRate

	ex.arg.Update.UpdatedAt = arg.Update.UpdatedAt
	ex.arg.Update.UpdatedBy = arg.Update.UpdatedBy

	if ex.arg.Update.ID == int64(0) {
		return false
	}

	if ex.arg.Update.AdminFeeAmount.Valid || ex.arg.Update.InstallmentAmount.Valid || ex.arg.Update.OtrAmount.Valid || ex.arg.Update.InterestRate.Valid {
		if ex.arg.Update.UpdatedAt.String == "" {
			return false
		}

		_, err := time.Parse(time.RFC3339Nano, ex.arg.Update.UpdatedAt.String)
		if err != nil {
			return false
		}

		if ex.arg.Update.UpdatedBy.String == "" {
			return false
		}
	}

	return reflect.DeepEqual(ex.arg, arg)
}

func (ex *eqUpdateTxParamsMatcher) String() string {
	var errMsg string

	if ex.arg.Update.ID == int64(0) {
		errMsg += "id should not be empty or zero\n"
	}

	if ex.arg.Update.AdminFeeAmount.Valid || ex.arg.Update.InstallmentAmount.Valid || ex.arg.Update.OtrAmount.Valid || ex.arg.Update.InterestRate.Valid {
		if ex.arg.Update.UpdatedAt.String == "" {
			errMsg += "updated_at should not be empty if any of updated field related is set\n"
		}

		_, err := time.Parse(time.RFC3339Nano, ex.arg.Update.UpdatedAt.String)
		if err != nil {
			errMsg += "updated at doesnt reflect the time.RFC3339Nano layout\n"
		}

		if ex.arg.Update.UpdatedBy.String == "" {
			errMsg += "updated by should not be empty\n"
		}
	}

	return errMsg + fmt.Sprintf("matches arg: %v\n", ex.arg)
}
