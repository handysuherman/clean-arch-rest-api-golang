package domain

import "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"

type AffiliatedDealer struct {
	ID                   int64   `json:"id"`
	AffiliatedDealerName string  `json:"affiliated_dealer_name"`
	CreatedAt            string  `json:"created_at"`
	UpdatedAt            string  `json:"updated_at"`
	UpdatedBy            *string `json:"updated_by"`
	IsActivated          bool    `json:"is_activated"`
	IsActivatedAt        string  `json:"is_activated_at"`
	IsActivatedUpdatedAt string  `json:"is_activated_updated_at"`
}

type AffiliatedDealerList struct {
	TotalCount  int                 `json:"total_count"`
	TotalPages  int                 `json:"total_pages"`
	Page        int                 `json:"page"`
	Size        int                 `json:"size"`
	HasNextPage bool                `json:"has_next_page"`
	List        []*AffiliatedDealer `json:"list"`
}

type CreateAffiliatedDealerDTORequestParams struct {
	AffiliatedDealerName string `json:"affiliated_dealer_name" validate:"required,gte=0,lte=255"`
}

type UpdateAffiliatedDealerDTORequestParams struct {
	AffiliatedDealerName *string `json:"affiliated_dealer_name,omitempty"`
	IsActivated          *bool   `json:"is_activated"`
}

type FetchDTORequestParams struct {
	Query      *string            `json:"query"`
	Pagination *helper.Pagination `json:"pagination"`
}
