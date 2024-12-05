package domain

import "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/helper"

type Consumer struct {
	ID                   int64   `json:"id"`
	Nik                  string  `json:"nik"`
	FullName             string  `json:"full_name"`
	LegalName            *string `json:"legal_name,omitempty"`
	BirthPlace           *string `json:"birth_place,omitempty"`
	BirthDate            *string `json:"birth_date,omitempty"`
	Salary               *string `json:"salary,omitempty"`
	KtpPhoto             *string `json:"ktp_photo,omitempty"`
	SelfiePhoto          *string `json:"selfie_photo,omitempty"`
	CreatedAt            string  `json:"created_at"`
	UpdatedAt            string  `json:"updated_at"`
	UpdatedBy            *string `json:"updated_by,omitempty"`
	IsActivated          bool    `json:"is_activated"`
	IsActivatedAt        string  `json:"is_activated_at"`
	IsActivatedUpdatedAt string  `json:"is_activated_updated_at"`
}

type ConsumerList struct {
	TotalCount  int         `json:"total_count"`
	TotalPages  int         `json:"total_pages"`
	Page        int         `json:"page"`
	Size        int         `json:"size"`
	HasNextPage bool        `json:"has_next_page"`
	List        []*Consumer `json:"list"`
}

type CreateConsumerDTORequestParams struct {
	Nik         string   `json:"nik" validate:"required,gte=0,lte=16"`
	FullName    string   `json:"full_name" validate:"required,gte=5,lte=255"`
	LegalName   *string  `json:"legal_name"`
	BirthPlace  *string  `json:"birth_place"`
	BirthDate   *string  `json:"birth_date"`
	Salary      *float64 `json:"salary"`
	KTPPhoto    *string  `json:"ktp_photo"`
	SelfiePhoto *string  `json:"selfie_photo"`
}

type UpdateConsumerDTORequestParams struct {
	FullName    *string  `json:"full_name,omitempty"`
	BirthPlace  *string  `json:"birth_place,omitempty"`
	BirthDate   *string  `json:"birth_date,omitempty"`
	Salary      *float64 `json:"salary,omitempty"`
	KTPPhoto    *string  `json:"ktp_photo,omitempty"`
	SelfiePhoto *string  `json:"selfie_photo,omitempty"`
	IsActivated *bool    `json:"is_activated,omitempty"`
}

type FetchDTORequestParams struct {
	Query      *string            `json:"query"`
	Pagination *helper.Pagination `json:"pagination"`
}
