package usecases

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/config"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/repository"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/logger"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

var (
	cfg  *config.Config
	tlog logger.Logger
)

func TestMain(m *testing.M) {
	_tlog := logger.NewLogger()

	tlog = _tlog

	wd, err := os.Getwd()
	if err != nil {
		tlog.Warnf("os.Getwd.err: %v", err)
		return
	}

	_cfg, err := config.New(fmt.Sprintf("%s/%s", findModuleRoot(wd), "config-dev.yaml"))
	if err != nil {
		tlog.Warnf("config.New.err: %v", err)
		return
	}

	cfg = _cfg

	os.Exit(m.Run())
}

type mockArgs struct {
	createParams     *domain.CreateRequestParams
	createRepoParams *repository.CreateParams
	updateParams     *domain.UpdateRequestParams
	updateRepoParams *repository.UpdateParams
	listParams       *domain.FetchParams
	listRepoParams   *repository.ListParams

	idempotencyKey string

	repoResponse *repository.AffiliatedDealer
}

type eqFindByIDMatcher struct {
	arg int64
}

func EqFindByIDMatcher(arg int64) gomock.Matcher {
	return &eqFindByIDMatcher{arg: arg}
}

func (ex *eqFindByIDMatcher) Matches(x interface{}) bool {
	arg, ok := x.(int64)
	if !ok {
		return false
	}

	ex.arg = arg

	if ex.arg == 0 {
		return false
	}

	if !reflect.DeepEqual(ex.arg, arg) {
		return false
	}

	return true
}

func (ex *eqFindByIDMatcher) String() string {
	if ex.arg == 0 {
		return "id should not be empty"
	}

	return fmt.Sprintf("matches arg: %v", ex.arg)
}

func createRandom(t *testing.T) *mockArgs {
	repoResponse := createRepositoryResponse(t)

	createParams := createParams(t, repoResponse)
	createRepoParams := createRepositoryParams(t, repoResponse)

	updateParams := updateParams(t, repoResponse)
	updateRepoParams := updateRepositoryParams(t, repoResponse)

	listParams := listParams(t, repoResponse)
	listRepoParams := listRepoParams(t, listParams)

	return &mockArgs{
		createParams:     createParams,
		idempotencyKey:   helper.RandomString(32),
		createRepoParams: createRepoParams,
		repoResponse:     repoResponse,
		updateParams:     updateParams,
		updateRepoParams: updateRepoParams,
		listParams:       listParams,
		listRepoParams:   listRepoParams,
	}
}

func listParams(t *testing.T, repoResponse *repository.AffiliatedDealer) *domain.FetchParams {
	return &domain.FetchParams{
		SearchText: repoResponse.AffiliatedDealerName,
		Pagination: helper.NewPaginationFromQueryParams("1", "1"),
	}
}

func listRepoParams(t *testing.T, params *domain.FetchParams) *repository.ListParams {
	searchText := "%" + params.SearchText + "%"

	return &repository.ListParams{
		AffiliatedDealerName: searchText,
		Limit:                int32(params.Pagination.GetLimit()),
		Offset:               int32(params.Pagination.GetOffset()),
	}
}

func createParams(t *testing.T, repoResponse *repository.AffiliatedDealer) *domain.CreateRequestParams {
	return &domain.CreateRequestParams{
		AffiliatedDealerName: repoResponse.AffiliatedDealerName,
	}
}

func updateParams(t *testing.T, repoResponse *repository.AffiliatedDealer) *domain.UpdateRequestParams {
	return &domain.UpdateRequestParams{
		AffiliatedDealerName: &repoResponse.AffiliatedDealerName,
	}
}

func updateRepositoryParams(t *testing.T, repoResponse *repository.AffiliatedDealer) *repository.UpdateParams {
	birthDateStr := fmt.Sprintf("%d-%02d-%02d", helper.RandomInt(2000, 2099), helper.RandomInt(1, 12), helper.RandomInt(1, 28))
	birthDateLayoutStr := "2006-01-02"

	parsedBirthDate, err := time.Parse(birthDateLayoutStr, birthDateStr)
	require.NoError(t, err)
	require.NotEmpty(t, parsedBirthDate)

	return &repository.UpdateParams{
		AffiliatedDealerName: sql.NullString{
			String: repoResponse.AffiliatedDealerName,
			Valid:  true,
		},
		IsActivated: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
		IsActivatedAt: sql.NullString{
			String: time.Now().Format(time.RFC3339Nano),
			Valid:  true,
		},
		IsActivatedUpdatedAt: sql.NullString{
			String: time.Now().Format(time.RFC3339Nano),
			Valid:  true,
		},
		UpdatedAt: sql.NullString{
			String: repoResponse.UpdatedAt,
			Valid:  true,
		},
		// for now, we need place holder, as soon authentication service involved this should be replaced by issuer id
		UpdatedBy: sql.NullString{
			String: "system",
			Valid:  true,
		},
		ID: repoResponse.ID,
	}
}

func createRepositoryParams(t *testing.T, repoResponse *repository.AffiliatedDealer) *repository.CreateParams {
	return &repository.CreateParams{
		AffiliatedDealerName: repoResponse.AffiliatedDealerName,
		CreatedAt:            repoResponse.CreatedAt,
	}
}

func createRepositoryResponse(t *testing.T) *repository.AffiliatedDealer {
	birthDateStr := fmt.Sprintf("%d-%02d-%02d", helper.RandomInt(2000, 2099), helper.RandomInt(1, 12), helper.RandomInt(1, 28))
	birthDateLayoutStr := "2006-01-02"

	parsedBirthDate, err := time.Parse(birthDateLayoutStr, birthDateStr)
	require.NoError(t, err)
	require.NotEmpty(t, parsedBirthDate)

	arg := repository.AffiliatedDealer{
		ID:                   helper.RandomInt(100, 100000),
		AffiliatedDealerName: helper.RandomString(12),
		CreatedAt:            time.Now().Format(time.RFC3339Nano),
		UpdatedAt:            time.Now().Format(time.RFC3339Nano),
		UpdatedBy: sql.NullString{
			String: "system",
			Valid:  true,
		},
		IsActivated:          true,
		IsActivatedAt:        time.Now().Format(time.RFC3339Nano),
		IsActivatedUpdatedAt: time.Now().Format(time.RFC3339Nano),
	}

	return &arg
}

func findModuleRoot(dir string) string {
	for {
		_, err := os.Stat(filepath.Join(dir, "go.mod"))
		if err == nil {
			return dir
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return ""
}

type mockSqlRes struct {
	res *int64
}

func newMockSqlResult(res *int64) sql.Result {
	return &mockSqlRes{
		res: res,
	}
}

func (m *mockSqlRes) LastInsertId() (int64, error) {
	if m.res != nil {
		return *m.res, nil
	}
	return helper.RandomInt(10, 1000000), nil
}

func (m *mockSqlRes) RowsAffected() (int64, error) {
	if m.res != nil {
		return *m.res, nil
	}
	return helper.RandomInt(10, 1000000), nil
}
