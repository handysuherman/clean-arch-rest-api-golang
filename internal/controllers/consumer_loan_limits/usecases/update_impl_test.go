package usecases

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_loan_limits/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_loan_limits/repository"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_loan_limits/repository/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_MOCK_UPDATE(t *testing.T) {
	mockArgs := createRandom(t)

	testCases := []struct {
		tname         string
		id            int64
		body          *domain.UpdateRequestParams
		stubs         func(store *mock.MockRepository)
		checkResponse func(t *testing.T, res int64, err error)
	}{
		{
			tname: "OK",
			id:    mockArgs.updateRepoParams.ID,
			body:  mockArgs.updateParams,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().Update(gomock.Any(), EqUpdateParamsMatcher(mockArgs.updateRepoParams)).Times(1).Return(nil)
				store.EXPECT().FindByID(gomock.Any(), EqFindByIDMatcher(mockArgs.updateRepoParams.ID)).Times(1).Return(mockArgs.repoResponse, nil)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.updateRepoParams.ID, 10)), gomock.Eq(mockArgs.repoResponse)).Times(1)
			},
			checkResponse: func(t *testing.T, res int64, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, res)
			},
		},
		{
			tname: "ERR_UPDATE_INTERNAL_SERVER_ERROR",
			id:    mockArgs.updateRepoParams.ID,
			body:  mockArgs.updateParams,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().Update(gomock.Any(), EqUpdateParamsMatcher(mockArgs.updateRepoParams)).Times(1).Return(sql.ErrConnDone)
				store.EXPECT().FindByID(gomock.Any(), EqFindByIDMatcher(mockArgs.updateRepoParams.ID)).Times(0)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.updateRepoParams.ID, 10)), gomock.Eq(mockArgs.repoResponse)).Times(0)
			},
			checkResponse: func(t *testing.T, res int64, err error) {
				require.Error(t, err)
				require.Empty(t, res)
			},
		},
		{
			tname: "ERR_FIND_BY_ID_NOT_FOUND",
			id:    mockArgs.updateRepoParams.ID,
			body:  mockArgs.updateParams,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().Update(gomock.Any(), EqUpdateParamsMatcher(mockArgs.updateRepoParams)).Times(1).Return(nil)
				store.EXPECT().FindByID(gomock.Any(), EqFindByIDMatcher(mockArgs.updateRepoParams.ID)).Times(1).Return(nil, sql.ErrNoRows)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.updateRepoParams.ID, 10)), gomock.Eq(mockArgs.repoResponse)).Times(0)
			},
			checkResponse: func(t *testing.T, res int64, err error) {
				require.Error(t, err)
				require.Empty(t, res)
			},
		},
		{
			tname: "ERR_FIND_BY_ID_INTERNAL_SERVER_ERROR",
			id:    mockArgs.updateRepoParams.ID,
			body:  mockArgs.updateParams,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().Update(gomock.Any(), EqUpdateParamsMatcher(mockArgs.updateRepoParams)).Times(1).Return(nil)
				store.EXPECT().FindByID(gomock.Any(), EqFindByIDMatcher(mockArgs.updateRepoParams.ID)).Times(1).Return(nil, sql.ErrConnDone)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.updateRepoParams.ID, 10)), gomock.Eq(mockArgs.repoResponse)).Times(0)
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

			actualBody, actualError := u.Update(context.TODO(), tc.id, tc.body)
			tc.checkResponse(t, actualBody, actualError)
		})
	}
}

type eqUpdateParamsMatcher struct {
	arg *repository.UpdateParams
}

func EqUpdateParamsMatcher(arg *repository.UpdateParams) gomock.Matcher {
	return &eqUpdateParamsMatcher{arg: arg}
}

func (ex *eqUpdateParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(*repository.UpdateParams)
	if !ok {
		return false
	}

	ex.arg.ID = arg.ID
	ex.arg.Tenor = arg.Tenor
	ex.arg.Amount = arg.Amount

	ex.arg.UpdatedAt = arg.UpdatedAt
	ex.arg.UpdatedBy = arg.UpdatedBy

	if ex.arg.ID == int64(0) {
		return false
	}

	if ex.arg.Tenor.Valid || ex.arg.Amount.Valid {
		if ex.arg.UpdatedAt.String == "" {
			return false
		}

		_, err := time.Parse(time.RFC3339Nano, ex.arg.UpdatedAt.String)
		if err != nil {
			return false
		}

		if ex.arg.UpdatedBy.String == "" {
			return false
		}
	}

	return reflect.DeepEqual(ex.arg, arg)
}

func (ex *eqUpdateParamsMatcher) String() string {
	var errMsg string

	if ex.arg.ID == int64(0) {
		errMsg += "id should not be empty or zero\n"
	}

	if ex.arg.ID == int64(0) {
		errMsg += "id should not be empty\n"
	}

	if ex.arg.Tenor.Valid || ex.arg.Amount.Valid {
		if ex.arg.UpdatedAt.String == "" {
			errMsg += "updated_at should not be empty if any of updated field related is set\n"
		}

		_, err := time.Parse(time.RFC3339Nano, ex.arg.UpdatedAt.String)
		if err != nil {
			errMsg += "updated at doesnt reflect the time.RFC3339Nano layout \n"
		}

		if ex.arg.UpdatedBy.String == "" {
			errMsg += "updated by should not be empty \n"
		}
	}

	return errMsg + fmt.Sprintf("matches arg: %v\n", ex.arg)
}
