package domain

import "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"

type CreateRequestParams struct {
	AffiliatedDealerName string `json:"affiliated_dealer_name" validate:"required,gte=0,lte=16"`
}

type UpdateRequestParams struct {
	AffiliatedDealerName *string `json:"affiliated_dealer_name,omitempty"`
}

type FetchParams struct {
	SearchText string             `json:"search_text"`
	Pagination *helper.Pagination `json:"pagination"`
}
