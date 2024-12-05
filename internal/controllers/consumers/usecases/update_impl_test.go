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
				store.EXPECT().FindByID(gomock.Any(), EqFindByIDMatcher(mockArgs.updateRepoParams.ID)).Times(1).Return(sql.ErrNoRows)
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
				store.EXPECT().FindByID(gomock.Any(), EqFindByIDMatcher(mockArgs.updateRepoParams.ID)).Times(1).Return(sql.ErrConnDone)
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
	ex.arg.FullName = arg.FullName
	ex.arg.BirthPlace = arg.BirthPlace
	ex.arg.BirthDate = arg.BirthDate
	ex.arg.Salary = arg.Salary
	ex.arg.KtpPhoto = arg.KtpPhoto
	ex.arg.SelfiePhoto = arg.SelfiePhoto

	ex.arg.IsActivated = arg.IsActivated
	ex.arg.IsActivatedAt = arg.IsActivatedAt
	ex.arg.IsActivatedUpdatedAt = arg.IsActivatedUpdatedAt

	ex.arg.UpdatedAt = arg.UpdatedAt
	ex.arg.UpdatedBy = arg.UpdatedBy

	if ex.arg.ID == int64(0) {
		return false
	}

	if ex.arg.FullName.Valid || ex.arg.BirthPlace.Valid || ex.arg.BirthDate.Valid || ex.arg.Salary.Valid || ex.arg.KtpPhoto.Valid || ex.arg.SelfiePhoto.Valid {
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

	if ex.arg.IsActivated.Valid {
		if ex.arg.IsActivatedUpdatedAt.String == "" {
			return false
		}

		_, err := time.Parse(time.RFC3339Nano, ex.arg.IsActivatedUpdatedAt.String)
		if err != nil {
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

	if ex.arg.FullName.Valid || ex.arg.BirthPlace.Valid || ex.arg.BirthDate.Valid || ex.arg.Salary.Valid || ex.arg.KtpPhoto.Valid || ex.arg.SelfiePhoto.Valid {
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

	if ex.arg.IsActivated.Valid {
		if ex.arg.IsActivatedUpdatedAt.String == "" {
			errMsg += "IsActivatedUpdatedAt should not be empty if IsActivated updated \n"
		}

		_, err := time.Parse(time.RFC3339Nano, ex.arg.IsActivatedUpdatedAt.String)
		if err != nil {
			errMsg += "IsActivatedUpdatedAt doesnt reflect the time.RFC3339Nano layout \n"
		}
	}

	return errMsg + fmt.Sprintf("matches arg: %v\n", ex.arg)
}
