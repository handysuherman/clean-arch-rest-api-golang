package app

import (
	affiliatedDealersHandler "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/delivery/http/v1"
	// affiliatedDealersRepo "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/repository"
	affiliatedDealersUsecase "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/usecases"

	consumersHandler "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/delivery/http/v1"
	// consumersRepo "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/repository"
	consumersUsecase "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/usecases"

	consumerLoanLimitsHandler "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_loan_limits/delivery/http/v1"
	// consumer_loan_limitsRepo "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_loan_limits/repository"
	consumerLoanLimitsUsecase "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_loan_limits/usecases"

	consumerTransactionsHandler "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/delivery/http/v1"
	consumerTransactionsRepo "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/repository"
	consumerTransactionsUsecase "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/usecases"
)

func (a *app) affiliatedDealersHandlers() *affiliatedDealersHandler.Handler {
	// repo := affiliatedDealersRepo.New(a.log, a.cfg)
	usecase := affiliatedDealersUsecase.New()
	handler := affiliatedDealersHandler.New(a.log, a.cfg, usecase, a.server, a.metrics)

	return handler
}

func (a *app) consumersHandlers() *consumersHandler.Handler {
	// repo := consumersRepo.New(a.log, a.cfg)
	usecase := consumersUsecase.New()
	handler := consumersHandler.New(a.log, a.cfg, usecase, a.server, a.metrics)

	return handler
}

func (a *app) consumerLoanLimitsHandlers() *consumerLoanLimitsHandler.Handler {
	// repo := consumerLoanLimitsRepo.New(a.log, a.cfg)
	usecase := consumerLoanLimitsUsecase.New()
	handler := consumerLoanLimitsHandler.New(a.log, a.cfg, usecase, a.server, a.metrics)

	return handler
}

func (a *app) consumerTransactionsHandlers() *consumerTransactionsHandler.Handler {
	repo := consumerTransactionsRepo.NewStore(a.log, a.cfg, a.mysqlConnection, a.redisConnection)
	usecase := consumerTransactionsUsecase.New(a.log, a.cfg, repo)
	handler := consumerTransactionsHandler.New(a.log, a.cfg, usecase, a.server, a.metrics, a.validator)

	return handler
}
