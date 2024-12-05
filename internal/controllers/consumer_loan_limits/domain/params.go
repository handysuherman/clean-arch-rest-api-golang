package domain

import "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"

type CreateRequestParams struct {
	ConsumerID int64  `json:"consumer_id" validate:"required"`
	Tenor      int16  `json:"tenor" validate:"required"`
	Amount     string `json:"amount" validate:"required"`
}

type UpdateRequestParams struct {
	Tenor  *int16  `json:"tenor"`
	Amount *string `json:"amount"`
}

type FetchParams struct {
	ConsumerID int64              `json:"consumer_id"`
	Pagination *helper.Pagination `json:"pagination"`
}
