package domain

type CreateDTORequestParams struct {
	Nik         string  `json:"nik" validate:"required,gte=0,lte=16"`
	FullName    string  `json:"full_name" validate:"required,gte=5,lte=255"`
	LegalName   string  `json:"legal_name"`
	BirthPlace  string  `json:"birth_place"`
	BirthDate   string  `json:"birth_date"`
	Salary      float64 `json:"salary"`
	KTPPhoto    string  `json:"ktp_photo"`
	SelfiePhoto string  `json:"selfie_photo"`
}

type UpdateDTORequestParams struct {
	Nik         *string  `json:"nik,omitempty"`
	FullName    *string  `json:"full_name,omitempty"`
	LegalName   *string  `json:"legal_name,omitempty"`
	BirthPlace  *string  `json:"birth_place,omitempty"`
	BirthDate   *string  `json:"birth_date,omitempty"`
	Salary      *float64 `json:"salary,omitempty"`
	KTPPhoto    *string  `json:"ktp_photo,omitempty"`
	SelfiePhoto *string  `json:"selfie_photo,omitempty"`
}
