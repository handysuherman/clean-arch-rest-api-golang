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
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/repository"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/logger"
	"github.com/shopspring/decimal"
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
	createParams        *domain.CreateRequestParams
	createRepoParams    *repository.CreateParams
	updateParams        *domain.UpdateRequestParams
	updateRepoParams    *repository.UpdateParams
	listParams          *domain.FetchParams
	listRepoParams      *repository.ListParams
	countListRepoParams *repository.CountListParams

	idempotencyKey string

	repoResponse *repository.Consumer
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
	listRepoParams, countListRepoParams := listRepoParams(t, listParams)

	return &mockArgs{
		createParams:        createParams,
		idempotencyKey:      helper.RandomString(32),
		createRepoParams:    createRepoParams,
		repoResponse:        repoResponse,
		updateParams:        updateParams,
		updateRepoParams:    updateRepoParams,
		listParams:          listParams,
		listRepoParams:      listRepoParams,
		countListRepoParams: countListRepoParams,
	}
}

func listParams(t *testing.T, repoResponse *repository.Consumer) *domain.FetchParams {
	return &domain.FetchParams{
		SearchText: repoResponse.FullName,
		Pagination: helper.NewPaginationFromQueryParams("1", "1"),
	}
}

func listRepoParams(t *testing.T, params *domain.FetchParams) (*repository.ListParams, *repository.CountListParams) {
	searchText := "%" + params.SearchText + "%s"

	countListArg := &repository.CountListParams{
		FullName: searchText,
		LegalName: sql.NullString{
			String: searchText,
			Valid:  true,
		},
	}

	return &repository.ListParams{
		FullName:  countListArg.FullName,
		LegalName: countListArg.LegalName,
		Limit:     int32(params.Pagination.GetLimit()),
		Offset:    int32(params.Pagination.GetOffset()),
	}, countListArg
}

func createParams(t *testing.T, repoResponse *repository.Consumer) *domain.CreateRequestParams {
	salaryFloat, _ := repoResponse.Salary.Decimal.Float64()
	birthDateStr := repoResponse.BirthDate.Time.String()

	return &domain.CreateRequestParams{
		Nik:         repoResponse.Nik,
		FullName:    repoResponse.FullName,
		LegalName:   &repoResponse.LegalName.String,
		BirthPlace:  &repoResponse.BirthPlace.String,
		BirthDate:   &birthDateStr,
		Salary:      &salaryFloat,
		SelfiePhoto: &repoResponse.SelfiePhoto.String,
		KTPPhoto:    &repoResponse.KtpPhoto.String,
	}
}

func updateParams(t *testing.T, repoResponse *repository.Consumer) *domain.UpdateRequestParams {
	salaryFloat, _ := repoResponse.Salary.Decimal.Float64()
	birthDateStr := repoResponse.BirthDate.Time.String()

	return &domain.UpdateRequestParams{
		FullName:    &repoResponse.FullName,
		Salary:      &salaryFloat,
		BirthPlace:  &repoResponse.BirthPlace.String,
		BirthDate:   &birthDateStr,
		KTPPhoto:    &repoResponse.KtpPhoto.String,
		SelfiePhoto: &repoResponse.SelfiePhoto.String,
	}
}

func updateRepositoryParams(t *testing.T, repoResponse *repository.Consumer) *repository.UpdateParams {
	birthDateStr := fmt.Sprintf("%d-%02d-%02d", helper.RandomInt(2000, 2099), helper.RandomInt(1, 12), helper.RandomInt(1, 28))
	birthDateLayoutStr := "2006-01-02"

	parsedBirthDate, err := time.Parse(birthDateLayoutStr, birthDateStr)
	require.NoError(t, err)
	require.NotEmpty(t, parsedBirthDate)

	return &repository.UpdateParams{
		FullName: sql.NullString{
			String: repoResponse.FullName,
		},
		Salary: decimal.NewNullDecimal(decimal.NewFromInt(helper.RandomInt(1000000, 10000000))),
		BirthPlace: sql.NullString{
			String: helper.RandomString(21),
			Valid:  true,
		},
		BirthDate: sql.NullTime{
			Time:  parsedBirthDate,
			Valid: true,
		},
		KtpPhoto: sql.NullString{
			String: fmt.Sprintf("%s/%s.jpeg", helper.RandomUrl(), helper.RandomString(12)),
			Valid:  true,
		},
		SelfiePhoto: sql.NullString{
			String: fmt.Sprintf("%s/%s.jpeg", helper.RandomUrl(), helper.RandomString(12)),
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

func createRepositoryParams(t *testing.T, repoResponse *repository.Consumer) *repository.CreateParams {
	return &repository.CreateParams{
		Nik:         repoResponse.Nik,
		FullName:    repoResponse.FullName,
		LegalName:   repoResponse.LegalName,
		BirthPlace:  repoResponse.BirthPlace,
		BirthDate:   repoResponse.BirthDate,
		Salary:      repoResponse.Salary,
		CreatedAt:   repoResponse.CreatedAt,
		SelfiePhoto: repoResponse.SelfiePhoto,
		KtpPhoto:    repoResponse.KtpPhoto,
	}
}

func createRepositoryResponse(t *testing.T) *repository.Consumer {
	birthDateStr := fmt.Sprintf("%d-%02d-%02d", helper.RandomInt(2000, 2099), helper.RandomInt(1, 12), helper.RandomInt(1, 28))
	birthDateLayoutStr := "2006-01-02"

	parsedBirthDate, err := time.Parse(birthDateLayoutStr, birthDateStr)
	require.NoError(t, err)
	require.NotEmpty(t, parsedBirthDate)

	arg := repository.Consumer{
		ID:       helper.RandomInt(100, 100000),
		Nik:      helper.RandomString(12),
		FullName: helper.RandomString(32),
		LegalName: sql.NullString{
			String: helper.RandomString(32),
			Valid:  true,
		},
		BirthPlace: sql.NullString{
			String: helper.RandomString(21),
			Valid:  true,
		},
		BirthDate: sql.NullTime{
			Time:  parsedBirthDate,
			Valid: true,
		},
		Salary: decimal.NewNullDecimal(decimal.NewFromInt(helper.RandomInt(1000000, 10000000))),
		KtpPhoto: sql.NullString{
			String: fmt.Sprintf("%s/%s.jpeg", helper.RandomUrl(), helper.RandomString(12)),
			Valid:  true,
		},
		SelfiePhoto: sql.NullString{
			String: fmt.Sprintf("%s/%s.jpeg", helper.RandomUrl(), helper.RandomString(12)),
			Valid:  true,
		},
		CreatedAt: time.Now().Format(time.RFC3339Nano),
		UpdatedAt: time.Now().Format(time.RFC3339Nano),
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

type mockSqlRes struct{}

func newMockSqlResult() sql.Result {
	return &mockSqlRes{}
}

func (m *mockSqlRes) LastInsertId() (int64, error) {
	return helper.RandomInt(10, 1000000), nil
}

func (m *mockSqlRes) RowsAffected() (int64, error) {
	return helper.RandomInt(10, 1000000), nil
}
