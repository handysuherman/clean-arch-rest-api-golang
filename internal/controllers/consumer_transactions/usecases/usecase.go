package usecase

import "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/domain"

type usecaseImpl struct{}

func New() domain.Usecase {
	return usecaseImpl{}
}
