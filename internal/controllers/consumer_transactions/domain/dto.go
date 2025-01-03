package domain

import "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"

type ConsumerTransaction struct {
	ID                 int64   `json:"id"`
	ConsumerID         int64   `json:"consumer_id"`
	ContractNumber     string  `json:"contract_number"`
	AdminFeeAmount     *string `json:"admin_fee_amount,omitempty"`
	InstallmentAmount  *string `json:"installment_amount,omitempty"`
	OtrAmount          *string `json:"otr_amount,omitempty"`
	InterestRate       *string `json:"interest_rate,omitempty"`
	TransactionDate    string  `json:"transaction_date"`
	AffiliatedDealerID int64   `json:"affiliated_dealer_id"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          string  `json:"updated_at"`
	UpdatedBy          *string `json:"updated_by,omitempty"`
}

type ConsumerTransactionList struct {
	TotalCount  int                    `json:"total_count"`
	TotalPages  int                    `json:"total_pages"`
	Page        int                    `json:"page"`
	Size        int                    `json:"size"`
	HasNextPage bool                   `json:"has_next_page"`
	List        []*ConsumerTransaction `json:"list"`
}

type CreateConsumerTransactionDTORequestParams struct {
	ConsumerID         int64   `json:"consumer_id" validate:"required"`
	AdminFeeAmount     *string `json:"admin_fee_amount"`
	InstallmentAmount  *string `json:"installment_amount"`
	OtrAmount          *string `json:"otr_amount"`
	InterestRate       *string `json:"interest_rate"`
	AffiliatedDealerID int64   `json:"affiliated_dealer_id" validate:"required"`
}

type UpdateConsumerTransactionDTORequestParams struct {
	AdminFeeAmount    *string `json:"admin_fee_amount"`
	InstallmentAmount *string `json:"installment_amount"`
	OtrAmount         *string `json:"otr_amount"`
	InterestRate      *string `json:"interest_rate"`
}

type FetchDTORequestParams struct {
	ConsumerID int64              `json:"consumer_id"`
	Pagination *helper.Pagination `json:"pagination"`
}
