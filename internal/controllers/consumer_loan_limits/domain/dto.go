package domain

import "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"

type ConsumerLoanLimit struct {
	ID         int64  `json:"id"`
	ConsumerID int64  `json:"consumer_id"`
	Tenor      int16  `json:"tenor"`
	Amount     string `json:"amount"`
	// format should be like 0001-01-01 00:00:00Z
	CreatedAt string `json:"created_at"`
	// format should be like 0001-01-01 00:00:00Z
	UpdatedAt string  `json:"updated_at"`
	UpdatedBy *string `json:"updated_by,omitempty"`
}

type ConsumerLoanLimitList struct {
	TotalCount  int                  `json:"total_count"`
	TotalPages  int                  `json:"total_pages"`
	Page        int                  `json:"page"`
	Size        int                  `json:"size"`
	HasNextPage bool                 `json:"has_next_page"`
	List        []*ConsumerLoanLimit `json:"list"`
}

type CreateConsumerLoanLimitDTORequestParams struct {
	ConsumerID int64  `json:"consumer_id" validate:"required"`
	Tenor      int16  `json:"tenor" validate:"required"`
	Amount     string `json:"amount" validate:"required"`
}

type UpdateConsumerLoanLimitDTORequestParams struct {
	Tenor  *int16  `json:"tenor"`
	Amount *string `json:"amount"`
}

type FetchDTORequestParams struct {
	ConsumerID int64              `json:"consumer_id"`
	Pagination *helper.Pagination `json:"pagination"`
}
