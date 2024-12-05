package mapper

import (
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_loan_limits/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_loan_limits/repository"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
)

func NewCreateRequestParams(arg *domain.CreateDTORequestParams) *domain.CreateRequestParams {
	return &domain.CreateRequestParams{
		ConsumerID: arg.ConsumerID,
		Tenor:      arg.Tenor,
		Amount:     arg.Amount,
	}
}

func NewUpdateRequestParams(arg *domain.UpdateDTORequestParams) *domain.UpdateRequestParams {
	return &domain.UpdateRequestParams{
		Tenor:  arg.Tenor,
		Amount: arg.Amount,
	}
}

func ToDTO(arg *repository.ConsumerLoanLimit) *domain.ConsumerLoanLimit {
	res := &domain.ConsumerLoanLimit{
		ID:         arg.ID,
		ConsumerID: arg.ConsumerID,
		Tenor:      arg.Tenor,
		Amount:     arg.Amount.String(),
		CreatedAt:  arg.CreatedAt,
		UpdatedAt:  arg.UpdatedAt,
	}

	if arg.UpdatedBy.Valid {
		res.UpdatedBy = &arg.UpdatedBy.String
	}

	return res
}

func ListToDTO(args []*repository.ConsumerLoanLimit) []*domain.ConsumerLoanLimit {
	list := make([]*domain.ConsumerLoanLimit, 0, len(args))

	for _, item := range args {
		list = append(list, ToDTO(item))
	}

	return list
}

func NewFetchParams(arg *domain.FetchDTORequestParams) *domain.FetchParams {
	return &domain.FetchParams{
		ConsumerID: arg.ConsumerID,
		Pagination: arg.Pagination,
	}
}

func NewConsumerLoanLimitList(
	list []*repository.ConsumerLoanLimit,
	count int64,
	pagination *helper.Pagination,
) *domain.ConsumerLoanLimitList {
	return &domain.ConsumerLoanLimitList{
		TotalCount:  int(count),
		TotalPages:  pagination.GetTotalPages(int(count)),
		Page:        pagination.GetPage(),
		Size:        pagination.GetSize(),
		HasNextPage: pagination.GetHasMore(int(count)),
		List:        ListToDTO(list),
	}
}
