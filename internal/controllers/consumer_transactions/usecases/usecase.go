package usecases

import (
	"fmt"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/config"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/repository"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/constants"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/logger"
)

type usecaseImpl struct {
	log        logger.Logger
	cfg        *config.Config
	repository repository.Repository
}

func New(
	log logger.Logger,
	cfg *config.Config,
	repository repository.Repository,
) domain.Usecase {
	return usecaseImpl{
		log:        log.WithPrefix(fmt.Sprintf("%s-%s", "consumer-transactions", constants.Usecase)),
		cfg:        cfg,
		repository: repository,
	}
}
