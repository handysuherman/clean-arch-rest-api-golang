package mapper

import (
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumer_transactions/repository"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
)

func NewCreateRequestParams(arg *domain.CreateDTORequestParams) *domain.CreateRequestParams {
	return &domain.CreateRequestParams{
		ConsumerID:         arg.ConsumerID,
		AffiliatedDealerID: arg.AffiliatedDealerID,
		AdminFeeAmount:     arg.AdminFeeAmount,
		InstallmentAmount:  arg.InstallmentAmount,
		OtrAmount:          arg.OtrAmount,
		InterestRate:       arg.InterestRate,
	}
}

func NewUpdateRequestParams(arg *domain.UpdateDTORequestParams) *domain.UpdateRequestParams {
	return &domain.UpdateRequestParams{
		AdminFeeAmount:    arg.AdminFeeAmount,
		InstallmentAmount: arg.InstallmentAmount,
		OtrAmount:         arg.OtrAmount,
		InterestRate:      arg.InterestRate,
	}
}

func ToDTO(arg *repository.ConsumerTransaction) *domain.ConsumerTransaction {
	res := &domain.ConsumerTransaction{
		ID:                 arg.ID,
		ConsumerID:         arg.ConsumerID,
		AffiliatedDealerID: arg.AffiliatedDealerID,
		ContractNumber:     arg.ContractNumber,
		TransactionDate:    arg.TransactionDate,
		CreatedAt:          arg.CreatedAt,
		UpdatedAt:          arg.UpdatedAt,
	}

	if arg.UpdatedBy.Valid {
		res.UpdatedBy = &arg.UpdatedBy.String
	}

	if arg.AdminFeeAmount.Valid {
		amount := arg.AdminFeeAmount.Decimal.String()
		res.AdminFeeAmount = &amount
	}

	if arg.InstallmentAmount.Valid {
		amount := arg.InstallmentAmount.Decimal.String()
		res.InstallmentAmount = &amount
	}

	if arg.OtrAmount.Valid {
		amount := arg.OtrAmount.Decimal.String()
		res.OtrAmount = &amount
	}

	if arg.InterestRate.Valid {
		amount := arg.InterestRate.Decimal.String()
		res.InterestRate = &amount
	}

	return res
}

func ListToDTO(args []*repository.ConsumerTransaction) []*domain.ConsumerTransaction {
	list := make([]*domain.ConsumerTransaction, 0, len(args))

	for _, item := range args {
		list = append(list, ToDTO(item))
	}

	return list
}

func NewFetchParams(arg *domain.FetchDTORequestParams) *domain.FetchParams {
	var (
		defaultPage = 1
		defaultSize = 10
		// defaultSearchText = ""
	)

	if arg.Page == nil {
		arg.Page = &defaultPage
	}

	if arg.Size == nil {
		arg.Size = &defaultSize
	}

	pq := helper.NewPaginationQuery(*arg.Size, *arg.Page)
	return &domain.FetchParams{
		ConsumerID: arg.ConsumerID,
		Pagination: pq,
	}
}

func NewConsumerTransactionList(
	list []*repository.ConsumerTransaction,
	count int64,
	pagination *helper.Pagination,
) *domain.ConsumerTransactionList {
	return &domain.ConsumerTransactionList{
		TotalCount:  int(count),
		TotalPages:  pagination.GetTotalPages(int(count)),
		Page:        pagination.GetPage(),
		Size:        pagination.GetSize(),
		HasNextPage: pagination.GetHasMore(int(count)),
		List:        ListToDTO(list),
	}
}
