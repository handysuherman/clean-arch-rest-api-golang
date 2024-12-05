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
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/repository"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/logger"
	"github.com/shopspring/decimal"
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
	idempotencyKey   string

	repoResponse *repository.ConsumerTransaction
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

	return &mockArgs{
		createParams:     createParams,
		idempotencyKey:   helper.RandomString(32),
		createRepoParams: createRepoParams,
		repoResponse:     repoResponse,
	}
}

func createParams(t *testing.T, repoResponse *repository.ConsumerTransaction) *domain.CreateRequestParams {
	adminFeeAmountStr := repoResponse.AdminFeeAmount.Decimal.String()
	installmentAmountStr := repoResponse.InstallmentAmount.Decimal.String()
	otrAmountStr := repoResponse.OtrAmount.Decimal.String()
	interestRateStr := repoResponse.InterestRate.Decimal.String()

	return &domain.CreateRequestParams{
		ConsumerID:         repoResponse.ConsumerID,
		AffiliatedDealerID: repoResponse.AffiliatedDealerID,
		AdminFeeAmount:     &adminFeeAmountStr,
		InstallmentAmount:  &installmentAmountStr,
		OtrAmount:          &otrAmountStr,
		InterestRate:       &interestRateStr,
	}
}

func createRepositoryParams(t *testing.T, repoResponse *repository.ConsumerTransaction) *repository.CreateParams {
	return &repository.CreateParams{
		ConsumerID:         repoResponse.ConsumerID,
		ContractNumber:     repoResponse.ContractNumber,
		AdminFeeAmount:     repoResponse.AdminFeeAmount,
		InstallmentAmount:  repoResponse.InstallmentAmount,
		OtrAmount:          repoResponse.OtrAmount,
		InterestRate:       repoResponse.InterestRate,
		TransactionDate:    repoResponse.TransactionDate,
		CreatedAt:          repoResponse.CreatedAt,
		AffiliatedDealerID: repoResponse.AffiliatedDealerID,
	}
}

func createRepositoryResponse(t *testing.T) *repository.ConsumerTransaction {
	currentTime := time.Now().Format(time.RFC3339Nano)

	return &repository.ConsumerTransaction{
		ID:                 helper.RandomInt(1, 12),
		ConsumerID:         helper.RandomInt(1, 100),
		AffiliatedDealerID: helper.RandomInt(100, 200),
		ContractNumber:     helper.RandomString(16),
		AdminFeeAmount:     decimal.NewNullDecimal(decimal.NewFromInt(helper.RandomInt(1000000, 10000000))),
		InstallmentAmount:  decimal.NewNullDecimal(decimal.NewFromInt(helper.RandomInt(1000000, 10000000))),
		OtrAmount:          decimal.NewNullDecimal(decimal.NewFromInt(helper.RandomInt(1000000, 10000000))),
		InterestRate:       decimal.NewNullDecimal(decimal.NewFromInt(helper.RandomInt(1, 30))),
		TransactionDate:    currentTime,
		CreatedAt:          currentTime,
		UpdatedAt:          currentTime,
		UpdatedBy: sql.NullString{
			String: helper.RandomString(12),
			Valid:  true,
		},
	}
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
