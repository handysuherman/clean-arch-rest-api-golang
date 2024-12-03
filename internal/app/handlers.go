package app

import (
	dropboxHandler "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/delivery/http/v1"
	// dropboxRepo "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/repository"
	dropboxUsecase "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/usecases"
)

func (a *app) dropboxHandlers() *dropboxHandler.Handler {
	// repo := dropboxRepo.New(a.log, a.cfg)
	usecase := dropboxUsecase.New()
	handler := dropboxHandler.New(a.log, a.cfg, usecase, a.server, a.metrics)

	return handler
}
