package usecases

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/repository"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/repository/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_MOCK_CREATE(t *testing.T) {
	mockArgs := createRandom(t)

	testCases := []struct {
		tname         string
		body          *domain.CreateRequestParams
		stubs         func(store *mock.MockRepository)
		checkResponse func(t *testing.T, res int64, err error)
	}{
		{
			tname: "OK",
			body:  mockArgs.createParams,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().Create(gomock.Any(), EqCreateParamsMatcher(mockArgs.createRepoParams)).Times(1).Return(newMockSqlResult(&mockArgs.repoResponse.ID), nil)
				store.EXPECT().FindByID(gomock.Any(), EqFindByIDMatcher(mockArgs.repoResponse.ID)).Times(1).Return(mockArgs.repoResponse, nil)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.repoResponse.ID, 10)), gomock.Eq(mockArgs.repoResponse)).Times(1)
			},
			checkResponse: func(t *testing.T, res int64, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, res)
			},
		},
		{
			tname: "ERR_CREATE_INTERNAL_SERVER_ERROR",
			body:  mockArgs.createParams,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().Create(gomock.Any(), EqCreateParamsMatcher(mockArgs.createRepoParams)).Times(1).Return(nil, sql.ErrConnDone)
				store.EXPECT().FindByID(gomock.Any(), EqFindByIDMatcher(mockArgs.repoResponse.ID)).Times(0)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.repoResponse.ID, 10)), gomock.Eq(mockArgs.repoResponse)).Times(0)
			},
			checkResponse: func(t *testing.T, res int64, err error) {
				require.Error(t, err)
				require.Empty(t, res)
			},
		},
		{
			tname: "ERR_FIND_BY_ID_NOT_FOUND",
			body:  mockArgs.createParams,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().Create(gomock.Any(), EqCreateParamsMatcher(mockArgs.createRepoParams)).Times(1).Return(newMockSqlResult(&mockArgs.repoResponse.ID), nil)
				store.EXPECT().FindByID(gomock.Any(), EqFindByIDMatcher(mockArgs.repoResponse.ID)).Times(1).Return(nil, sql.ErrNoRows)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.repoResponse.ID, 10)), gomock.Eq(mockArgs.repoResponse)).Times(0)
			},
			checkResponse: func(t *testing.T, res int64, err error) {
				require.Error(t, err)
				require.Empty(t, res)
			},
		},
		{
			tname: "ERR_FIND_BY_ID_INTERNAL_SERVER_ERROR",
			body:  mockArgs.createParams,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().Create(gomock.Any(), EqCreateParamsMatcher(mockArgs.createRepoParams)).Times(1).Return(newMockSqlResult(&mockArgs.repoResponse.ID), nil)
				store.EXPECT().FindByID(gomock.Any(), EqFindByIDMatcher(mockArgs.repoResponse.ID)).Times(1).Return(nil, sql.ErrConnDone)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.repoResponse.ID, 10)), gomock.Eq(mockArgs.repoResponse)).Times(0)
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

			actualBody, actualError := u.Create(context.TODO(), tc.body)
			tc.checkResponse(t, actualBody, actualError)
		})
	}
}

type eqCreateParamsMatcher struct {
	arg *repository.CreateParams
}

func EqCreateParamsMatcher(arg *repository.CreateParams) gomock.Matcher {
	return &eqCreateParamsMatcher{arg: arg}
}

func (ex *eqCreateParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(*repository.CreateParams)
	if !ok {
		return false
	}

	ex.arg.AffiliatedDealerName = arg.AffiliatedDealerName
	ex.arg.CreatedAt = arg.CreatedAt

	if ex.arg.AffiliatedDealerName == "" {
		return false
	}

	_, err := time.Parse(time.RFC3339Nano, ex.arg.CreatedAt)
	if err != nil {
		return false
	}

	return reflect.DeepEqual(ex.arg, arg)
}

func (ex *eqCreateParamsMatcher) String() string {
	var errMsg string

	if ex.arg.AffiliatedDealerName == "" {
		errMsg += "affiliated dealer name should not be empty\n"
	}

	_, err := time.Parse(time.RFC3339Nano, ex.arg.CreatedAt)
	if err != nil {
		errMsg += "created at doesnt reflect the time.RFC3339Nano layout\n"
	}

	return errMsg + fmt.Sprintf("matches arg: %v\n", ex.arg)
}
