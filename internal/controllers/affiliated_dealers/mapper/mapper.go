package mapper

import (
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/affiliated_dealers/repository"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
)

func NewCreateRequestParams(arg *domain.CreateAffiliatedDealerDTORequestParams) *domain.CreateRequestParams {
	return &domain.CreateRequestParams{
		AffiliatedDealerName: arg.AffiliatedDealerName,
	}
}

func NewUpdateRequestParams(arg *domain.UpdateAffiliatedDealerDTORequestParams) *domain.UpdateRequestParams {
	return &domain.UpdateRequestParams{
		AffiliatedDealerName: arg.AffiliatedDealerName,
	}
}

func ToDTO(arg *repository.AffiliatedDealer) *domain.AffiliatedDealer {
	res := &domain.AffiliatedDealer{
		ID:                   arg.ID,
		AffiliatedDealerName: arg.AffiliatedDealerName,
		CreatedAt:            arg.CreatedAt,
		UpdatedAt:            arg.UpdatedAt,
		IsActivated:          arg.IsActivated,
		IsActivatedAt:        arg.IsActivatedAt,
		IsActivatedUpdatedAt: arg.IsActivatedUpdatedAt,
	}

	if arg.UpdatedBy.Valid {
		res.UpdatedBy = &arg.UpdatedBy.String
	}

	return res
}

func ListToDTO(args []*repository.AffiliatedDealer) []*domain.AffiliatedDealer {
	list := make([]*domain.AffiliatedDealer, 0, len(args))

	for _, item := range args {
		list = append(list, ToDTO(item))
	}

	return list
}

func NewFetchParams(arg *domain.FetchDTORequestParams) *domain.FetchParams {
	var (
		searchText = ""
	)

	if arg.Query != nil {
		searchText = *arg.Query
	}

	return &domain.FetchParams{
		SearchText: searchText,
		Pagination: arg.Pagination,
	}
}

func NewAffiliatedDealerList(
	list []*repository.AffiliatedDealer,
	count int64,
	pagination *helper.Pagination,
) *domain.AffiliatedDealerList {
	return &domain.AffiliatedDealerList{
		TotalCount:  int(count),
		TotalPages:  pagination.GetTotalPages(int(count)),
		Page:        pagination.GetPage(),
		Size:        pagination.GetSize(),
		HasNextPage: pagination.GetHasMore(int(count)),
		List:        ListToDTO(list),
	}
}
