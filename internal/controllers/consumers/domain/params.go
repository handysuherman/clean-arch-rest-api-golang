package domain

import "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"

type CreateRequestParams struct {
	Nik         string   `json:"nik" validate:"required,gte=0,lte=16"`
	FullName    string   `json:"full_name" validate:"required,gte=5,lte=255"`
	LegalName   *string  `json:"legal_name"`
	BirthPlace  *string  `json:"birth_place"`
	BirthDate   *string  `json:"birth_date"`
	Salary      *float64 `json:"salary"`
	KTPPhoto    *string  `json:"ktp_photo"`
	SelfiePhoto *string  `json:"selfie_photo"`
}

type UpdateRequestParams struct {
	FullName    *string  `json:"full_name"`
	BirthPlace  *string  `json:"birth_place"`
	BirthDate   *string  `json:"birth_date"`
	Salary      *float64 `json:"salary"`
	KTPPhoto    *string  `json:"ktp_photo"`
	SelfiePhoto *string  `json:"selfie_photo"`
}

type FetchParams struct {
	SearchText string             `json:"search_text"`
	Pagination *helper.Pagination `json:"pagination"`
}
