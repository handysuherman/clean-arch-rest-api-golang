package usecases

import (
	"context"
	"database/sql"
	"strconv"
	"testing"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/repository/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_MOCK_FIND_BY_ID(t *testing.T) {
	mockArgs := createRandom(t)

	testCases := []struct {
		tname         string
		id            int64
		stubs         func(store *mock.MockRepository)
		checkResponse func(t *testing.T, res *domain.Consumer, err error)
	}{
		{
			tname: "OK_CACHE_HIT",
			id:    mockArgs.repoResponse.ID,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().Get(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.repoResponse.ID, 10))).Times(1).Return(mockArgs.repoResponse, nil)
				store.EXPECT().FindByID(gomock.Any(), EqFindByIDMatcher(mockArgs.repoResponse.ID)).Times(0)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.repoResponse.ID, 10)), gomock.Eq(mockArgs.repoResponse)).Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.Consumer, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, res)

				require.Equal(t, mockArgs.repoResponse.ID, res.ID)
				require.Equal(t, mockArgs.repoResponse.Nik, res.Nik)
				require.Equal(t, mockArgs.repoResponse.FullName, res.FullName)
				require.Equal(t, mockArgs.repoResponse.CreatedAt, res.CreatedAt)
				require.Equal(t, mockArgs.repoResponse.UpdatedAt, res.UpdatedAt)
				require.Equal(t, mockArgs.repoResponse.IsActivated, res.IsActivated)
				require.Equal(t, mockArgs.repoResponse.IsActivatedAt, res.IsActivatedAt)
				require.Equal(t, mockArgs.repoResponse.IsActivatedUpdatedAt, res.IsActivatedUpdatedAt)

				if mockArgs.repoResponse.LegalName.Valid {
					require.Equal(t, mockArgs.repoResponse.LegalName.String, *res.LegalName)
				} else {
					require.Nil(t, res.LegalName)
				}

				if mockArgs.repoResponse.BirthPlace.Valid {
					require.Equal(t, mockArgs.repoResponse.BirthPlace.String, *res.BirthPlace)
				} else {
					require.Nil(t, res.BirthPlace)
				}

				if mockArgs.repoResponse.BirthDate.Valid {
					require.Equal(t, mockArgs.repoResponse.BirthDate.Time.Format("2006-01-02"), *res.BirthDate)
				} else {
					require.Nil(t, res.BirthDate)
				}

				if mockArgs.repoResponse.Salary.Valid {
					require.Equal(t, mockArgs.repoResponse.Salary.Decimal.String(), *res.Salary)
				} else {
					require.Nil(t, res.Salary)
				}

				if mockArgs.repoResponse.KtpPhoto.Valid {
					require.Equal(t, mockArgs.repoResponse.KtpPhoto.String, *res.KtpPhoto)
				} else {
					require.Nil(t, res.KtpPhoto)
				}

				if mockArgs.repoResponse.SelfiePhoto.Valid {
					require.Equal(t, mockArgs.repoResponse.SelfiePhoto.String, *res.SelfiePhoto)
				} else {
					require.Nil(t, res.SelfiePhoto)
				}

				if mockArgs.repoResponse.UpdatedBy.Valid {
					require.Equal(t, mockArgs.repoResponse.UpdatedBy.String, *res.UpdatedBy)
				} else {
					require.Nil(t, res.UpdatedBy)
				}
			},
		},
		{
			tname: "OK_CACHE_MISS",
			id:    mockArgs.repoResponse.ID,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().Get(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.repoResponse.ID, 10))).Times(1).Return(nil, nil)
				store.EXPECT().FindByID(gomock.Any(), EqFindByIDMatcher(mockArgs.repoResponse.ID)).Times(1).Return(mockArgs.repoResponse, nil)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.repoResponse.ID, 10)), gomock.Eq(mockArgs.repoResponse)).Times(1)
			},
			checkResponse: func(t *testing.T, res *domain.Consumer, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, res)

				require.Equal(t, mockArgs.repoResponse.ID, res.ID)
				require.Equal(t, mockArgs.repoResponse.Nik, res.Nik)
				require.Equal(t, mockArgs.repoResponse.FullName, res.FullName)
				require.Equal(t, mockArgs.repoResponse.CreatedAt, res.CreatedAt)
				require.Equal(t, mockArgs.repoResponse.UpdatedAt, res.UpdatedAt)

				require.Equal(t, mockArgs.repoResponse.IsActivated, res.IsActivated)
				require.Equal(t, mockArgs.repoResponse.IsActivatedAt, res.IsActivatedAt)
				require.Equal(t, mockArgs.repoResponse.IsActivatedUpdatedAt, res.IsActivatedUpdatedAt)

				if mockArgs.repoResponse.LegalName.Valid {
					require.Equal(t, mockArgs.repoResponse.LegalName.String, *res.LegalName)
				} else {
					require.Nil(t, res.LegalName)
				}

				if mockArgs.repoResponse.BirthPlace.Valid {
					require.Equal(t, mockArgs.repoResponse.BirthPlace.String, *res.BirthPlace)
				} else {
					require.Nil(t, res.BirthPlace)
				}

				if mockArgs.repoResponse.BirthDate.Valid {
					require.Equal(t, mockArgs.repoResponse.BirthDate.Time.Format("2006-01-02"), *res.BirthDate)
				} else {
					require.Nil(t, res.BirthDate)
				}

				if mockArgs.repoResponse.Salary.Valid {
					require.Equal(t, mockArgs.repoResponse.Salary.Decimal.String(), *res.Salary)
				} else {
					require.Nil(t, res.Salary)
				}

				if mockArgs.repoResponse.KtpPhoto.Valid {
					require.Equal(t, mockArgs.repoResponse.KtpPhoto.String, *res.KtpPhoto)
				} else {
					require.Nil(t, res.KtpPhoto)
				}

				if mockArgs.repoResponse.SelfiePhoto.Valid {
					require.Equal(t, mockArgs.repoResponse.SelfiePhoto.String, *res.SelfiePhoto)
				} else {
					require.Nil(t, res.SelfiePhoto)
				}

				if mockArgs.repoResponse.UpdatedBy.Valid {
					require.Equal(t, mockArgs.repoResponse.UpdatedBy.String, *res.UpdatedBy)
				} else {
					require.Nil(t, res.UpdatedBy)
				}
			},
		},
		{
			tname: "ERR_FIND_BY_ID_NOT_FOUND",
			id:    mockArgs.repoResponse.ID,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().Get(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.repoResponse.ID, 10))).Times(1).Return(nil, nil)
				store.EXPECT().FindByID(gomock.Any(), EqFindByIDMatcher(mockArgs.repoResponse.ID)).Times(1).Return(nil, sql.ErrNoRows)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.repoResponse.ID, 10)), gomock.Eq(mockArgs.repoResponse)).Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.Consumer, err error) {
				require.Error(t, err)
				require.Empty(t, res)
			},
		},
		{
			tname: "ERR_FIND_BY_ID_INTERNAL_SERVER_ERROR",
			id:    mockArgs.repoResponse.ID,
			stubs: func(store *mock.MockRepository) {
				store.EXPECT().Get(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.repoResponse.ID, 10))).Times(1).Return(nil, nil)
				store.EXPECT().FindByID(gomock.Any(), EqFindByIDMatcher(mockArgs.repoResponse.ID)).Times(1).Return(nil, sql.ErrConnDone)
				store.EXPECT().Put(gomock.Any(), gomock.Eq(strconv.FormatInt(mockArgs.repoResponse.ID, 10)), gomock.Eq(mockArgs.repoResponse)).Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.Consumer, err error) {
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

			actualBody, actualError := u.FindByID(context.TODO(), tc.id)
			tc.checkResponse(t, actualBody, actualError)
		})
	}
}
