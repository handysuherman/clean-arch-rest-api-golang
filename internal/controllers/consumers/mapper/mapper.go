package mapper

import (
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/domain"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/controllers/consumers/repository"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"
)

func NewCreateRequestParams(arg *domain.CreateDTORequestParams) *domain.CreateRequestParams {
	return &domain.CreateRequestParams{
		Nik:         arg.Nik,
		FullName:    arg.FullName,
		LegalName:   arg.LegalName,
		BirthPlace:  arg.BirthPlace,
		BirthDate:   arg.BirthDate,
		Salary:      arg.Salary,
		KTPPhoto:    arg.KTPPhoto,
		SelfiePhoto: arg.SelfiePhoto,
	}
}

func NewUpdateRequestParams(arg *domain.UpdateDTORequestParams) *domain.UpdateRequestParams {
	return &domain.UpdateRequestParams{
		Nik:         arg.Nik,
		FullName:    arg.FullName,
		LegalName:   arg.LegalName,
		BirthPlace:  arg.BirthPlace,
		BirthDate:   arg.BirthDate,
		Salary:      arg.Salary,
		KTPPhoto:    arg.KTPPhoto,
		SelfiePhoto: arg.SelfiePhoto,
	}
}

func ToDTO(arg *repository.Consumer) *domain.Consumer {
	res := &domain.Consumer{
		ID:                   arg.ID,
		Nik:                  arg.Nik,
		FullName:             arg.FullName,
		CreatedAt:            arg.CreatedAt,
		UpdatedAt:            arg.UpdatedAt,
		IsActivated:          arg.IsActivated,
		IsActivatedAt:        arg.IsActivatedAt,
		IsActivatedUpdatedAt: arg.IsActivatedUpdatedAt,
	}

	if arg.LegalName.Valid {
		res.LegalName = &arg.LegalName.String
	}

	if arg.BirthPlace.Valid {
		res.BirthPlace = &arg.BirthPlace.String
	}

	if arg.BirthDate.Valid {
		dateStr := arg.BirthDate.Time.String()
		res.BirthDate = &dateStr
	}

	if arg.Salary.Valid {
		salaryStr := arg.Salary.Decimal.String()
		res.Salary = &salaryStr
	}

	if arg.KtpPhoto.Valid {
		res.KtpPhoto = &arg.KtpPhoto.String
	}

	if arg.SelfiePhoto.Valid {
		res.SelfiePhoto = &arg.SelfiePhoto.String
	}

	if arg.UpdatedBy.Valid {
		res.UpdatedBy = &arg.UpdatedBy.String
	}

	return res
}

func ListToDTO(args []*repository.Consumer) []*domain.Consumer {
	list := make([]*domain.Consumer, 0, len(args))

	for _, item := range args {
		list = append(list, ToDTO(item))
	}

	return list
}

func NewFetchProductsParams(arg *domain.FetchDTORequestParams) *domain.FetchParams {
	var (
		defaultPage       = 1
		defaultSize       = 10
		defaultSearchText = ""
	)

	if arg.Page == nil {
		arg.Page = &defaultPage
	}

	if arg.Size == nil {
		arg.Size = &defaultSize
	}

	if arg.Query == nil {
		arg.Query = &defaultSearchText
	}

	pq := helper.NewPaginationQuery(*arg.Size, *arg.Page)
	return &domain.FetchParams{
		SearchText: *arg.Query,
		Pagination: pq,
	}
}

func NewConsumerList(
	list []*repository.Consumer,
	count int64,
	pagination *helper.Pagination,
) *domain.ConsumerList {
	return &domain.ConsumerList{
		TotalCount:  int(count),
		TotalPages:  pagination.GetTotalPages(int(count)),
		Page:        pagination.GetPage(),
		Size:        pagination.GetSize(),
		HasNextPage: pagination.GetHasMore(int(count)),
		List:        ListToDTO(list),
	}
}
