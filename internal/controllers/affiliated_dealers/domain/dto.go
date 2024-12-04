package domain

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

type CreateDTORequestParams struct {
	AffiliatedDealerName string `json:"affiliated_dealer_name" validate:"required,gte=0,lte=16"`
}

type UpdateDTORequestParams struct {
	AffiliatedDealerName *string `json:"affiliated_dealer_name,omitempty"`
}

type FetchDTORequestParams struct {
	Query *string `json:"query"`
	Page  *int    `json:"page"`
	Size  *int    `json:"size"`
}
