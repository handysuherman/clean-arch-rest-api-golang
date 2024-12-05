package domain

import "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"

type CreateRequestParams struct {
	ConsumerID         int64   `json:"consumer_id"`
	AdminFeeAmount     *string `json:"admin_fee_amount"`
	InstallmentAmount  *string `json:"installment_amount"`
	OtrAmount          *string `json:"otr_amount"`
	InterestRate       *string `json:"interest_rate"`
	AffiliatedDealerID int64   `json:"affiliated_dealer_id"`
}

type UpdateRequestParams struct {
	AdminFeeAmount    *string `json:"admin_fee_amount"`
	InstallmentAmount *string `json:"installment_amount"`
	OtrAmount         *string `json:"otr_amount"`
	InterestRate      *string `json:"interest_rate"`
}

type FetchParams struct {
	SearchText string             `json:"search_text"`
	Pagination *helper.Pagination `json:"pagination"`
}
