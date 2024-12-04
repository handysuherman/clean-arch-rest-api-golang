package metrics

import (
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/config"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/constants"
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	SuccessHTTPRequest prometheus.Counter
	ErrorHTTPRequest   prometheus.Counter

	CreateConsumerHTTPRequest   prometheus.Counter
	UpdateConsumerHTTPRequest   prometheus.Counter
	FindConsumerHTTPRequest     prometheus.Counter
	FindByIDConsumerHTTPRequest prometheus.Counter

	CreateConsumerLoanLimitHTTPRequest   prometheus.Counter
	UpdateConsumerLoanLimitHTTPRequest   prometheus.Counter
	FindConsumerLoanLimitHTTPRequest     prometheus.Counter
	FindByIDConsumerLoanLimitHTTPRequest prometheus.Counter

	CreateConsumerTransactionHTTPRequest   prometheus.Counter
	UpdateConsumerTransactionHTTPRequest   prometheus.Counter
	FindConsumerTransactionHTTPRequest     prometheus.Counter
	FindByIDConsumerTransactionHTTPRequest prometheus.Counter

	CreateAffiliatedDealerHTTPRequest   prometheus.Counter
	UpdateAffiliatedDealerHTTPRequest   prometheus.Counter
	FindAffiliatedDealerHTTPRequest     prometheus.Counter
	FindByIDAffiliatedDealerHTTPRequest prometheus.Counter
}

func New(cfg *config.Internal) *Metrics {
	return &Metrics{
		SuccessHTTPRequest: NewCounter(cfg, "success", constants.HTTP),
		ErrorHTTPRequest:   NewCounter(cfg, "error", constants.HTTP),

		CreateConsumerHTTPRequest:   NewCounter(cfg, "create_consumer", constants.HTTP),
		UpdateConsumerHTTPRequest:   NewCounter(cfg, "update_consumer", constants.HTTP),
		FindConsumerHTTPRequest:     NewCounter(cfg, "find_consumer", constants.HTTP),
		FindByIDConsumerHTTPRequest: NewCounter(cfg, "find_by_id_consumer", constants.HTTP),

		CreateConsumerLoanLimitHTTPRequest:   NewCounter(cfg, "create_consumer_loan_limit", constants.HTTP),
		UpdateConsumerLoanLimitHTTPRequest:   NewCounter(cfg, "update_consumer_loan_limit", constants.HTTP),
		FindConsumerLoanLimitHTTPRequest:     NewCounter(cfg, "find_consumer_loan_limit", constants.HTTP),
		FindByIDConsumerLoanLimitHTTPRequest: NewCounter(cfg, "find_by_id_consumer_loan_limit", constants.HTTP),

		CreateConsumerTransactionHTTPRequest:   NewCounter(cfg, "create_consumer_transaction", constants.HTTP),
		UpdateConsumerTransactionHTTPRequest:   NewCounter(cfg, "update_consumer_transaction", constants.HTTP),
		FindConsumerTransactionHTTPRequest:     NewCounter(cfg, "find_consumer_transaction", constants.HTTP),
		FindByIDConsumerTransactionHTTPRequest: NewCounter(cfg, "find_by_id_consumer_transaction", constants.HTTP),

		CreateAffiliatedDealerHTTPRequest:   NewCounter(cfg, "create_affiliated_dealer", constants.HTTP),
		UpdateAffiliatedDealerHTTPRequest:   NewCounter(cfg, "update_affiliated_dealer", constants.HTTP),
		FindAffiliatedDealerHTTPRequest:     NewCounter(cfg, "find_affiliated_dealer", constants.HTTP),
		FindByIDAffiliatedDealerHTTPRequest: NewCounter(cfg, "find_by_id_affiliated_dealer", constants.HTTP),
	}
}
