package usecases

import (
	"context"
	"database/sql"
	"testing"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/repository"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/repository/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_MOCK_LIST(t *testing.T) {
	mockArgs := createRandom(t)

	mockList := []*repository.ConsumerTransaction{mockArgs.repoResponse}

	testCases := []struct {
		tname         string
		body          *domain.FetchParams
		stubs         func(store *mock.MockRepository)
		checkResponse func(t *testing.T, res *domain.ConsumerTransactionList, err error)
	}{
		{
			tname: "OK",
			body:  mockArgs.listParams,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().CountList(gomock.Any(), gomock.Eq(mockArgs.repoResponse.ConsumerID)).Times(1).Return(int64(1), nil)
				store.EXPECT().List(gomock.Any(), gomock.Eq(mockArgs.listRepoParams)).Times(1).Return(mockList, nil)
			},
			checkResponse: func(t *testing.T, res *domain.ConsumerTransactionList, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, res)
			},
		},
		{
			tname: "ERR_COUNT_LIST_INTERNAL_SERVER_ERROR",
			body:  mockArgs.listParams,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().CountList(gomock.Any(), gomock.Eq(mockArgs.repoResponse.ConsumerID)).Times(1).Return(int64(0), sql.ErrConnDone)
				store.EXPECT().List(gomock.Any(), gomock.Eq(mockArgs.listRepoParams)).Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.ConsumerTransactionList, err error) {
				require.Error(t, err)
				require.Empty(t, res)
			},
		},
		{
			tname: "ERR_LIST_INTERNAL_SERVER_ERROR",
			body:  mockArgs.listParams,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().CountList(gomock.Any(), gomock.Eq(mockArgs.repoResponse.ConsumerID)).Times(1).Return(int64(1), nil)
				store.EXPECT().List(gomock.Any(), gomock.Eq(mockArgs.listRepoParams)).Times(1).Return(nil, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, res *domain.ConsumerTransactionList, err error) {
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

			actualBody, actualError := u.List(context.TODO(), tc.body)
			tc.checkResponse(t, actualBody, actualError)
		})
	}
}
