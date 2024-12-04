package usecase

import "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_loan_limits/domain"

type usecaseImpl struct{}

func New() domain.Usecase {
	return usecaseImpl{}
}
