package app

import (
	affiliatedDealersHandler "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/delivery/http/v1"
	// affiliatedDealersRepo "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/repository"
	affiliatedDealersUsecase "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/usecases"

	consumersHandler "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/delivery/http/v1"
	// consumersRepo "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/repository"
	consumersUsecase "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/usecases"
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
