package usecases

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/repository"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/repository/mock"
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
				store.EXPECT().Create(gomock.Any(), EqCreateParamsMatcher(mockArgs.createRepoParams)).Times(1).Return(newMockSqlResult(), nil)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.repoResponse.ID, 10)), gomock.Eq(mockArgs.repoResponse)).Times(1)
			},
			checkResponse: func(t *testing.T, res int64, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, res)
			},
		},
		{
			tname: "ERR_INTERNAL_SERVER_ERROR",
			body:  mockArgs.createParams,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().Create(gomock.Any(), EqCreateParamsMatcher(mockArgs.createRepoParams)).Times(1).Return(nil, sql.ErrConnDone)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.repoResponse.ID, 10)), gomock.Eq(mockArgs.repoResponse)).Times(0)
			},
			checkResponse: func(t *testing.T, res int64, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, res)
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

	ex.arg.Nik = arg.Nik
	ex.arg.FullName = arg.FullName
	ex.arg.LegalName = arg.LegalName
	ex.arg.BirthPlace = arg.BirthPlace
	ex.arg.BirthDate = arg.BirthDate
	ex.arg.Salary = arg.Salary
	ex.arg.KtpPhoto = arg.KtpPhoto
	ex.arg.CreatedAt = arg.CreatedAt
	ex.arg.SelfiePhoto = arg.SelfiePhoto

	if ex.arg.Nik == "" {
		return false
	}

	if ex.arg.FullName == "" {
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

	if ex.arg.Nik == "" {
		errMsg += "nik should not be empty\n"
	}

	if ex.arg.FullName == "" {
		errMsg += "full name should not be empty\n"
	}

	_, err := time.Parse(time.RFC3339Nano, ex.arg.CreatedAt)
	if err != nil {
		errMsg += "created at doesnt reflect the time.RFC3339Nano layout\n"
	}

	return errMsg + fmt.Sprintf("matches arg: %v\n", ex.arg)
}