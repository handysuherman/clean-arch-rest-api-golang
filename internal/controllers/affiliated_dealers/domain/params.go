package domain

type CreateRequestParams struct {
	Nik         string  `json:"nik" validate:"required,gte=0,lte=16"`
	FullName    string  `json:"full_name" validate:"required,gte=5,lte=255"`
	LegalName   string  `json:"legal_name"`
	BirthPlace  string  `json:"birth_place"`
	BirthDate   string  `json:"birth_date"`
	Salary      float64 `json:"salary"`
	KTPPhoto    string  `json:"ktp_photo"`
	SelfiePhoto string  `json:"selfie_photo"`
}

func NewCreateRequestParams(arg *CreateDTORequestParams) *CreateRequestParams {
	return &CreateRequestParams{
		Nik:         arg.Nik,
		FullName:    arg.FullName,
		LegalName:   arg.LegalName,
		BirthPlace:  arg.BirthPlace,
		BirthDate:   arg.BirthDate,
		Salary:      arg.Salary,
		KTPPhoto:    arg.KTPPhoto,
		SelfiePhoto: arg.SelfiePhoto,
	}
}

type UpdateRequestParams struct {
	Nik         *string  `json:"nik"`
	FullName    *string  `json:"full_name"`
	LegalName   *string  `json:"legal_name"`
	BirthPlace  *string  `json:"birth_place"`
	BirthDate   *string  `json:"birth_date"`
	Salary      *float64 `json:"salary"`
	KTPPhoto    *string  `json:"ktp_photo"`
	SelfiePhoto *string  `json:"selfie_photo"`
}

func NewUpdateRequestParams(arg *UpdateDTORequestParams) *UpdateRequestParams {
	return &UpdateRequestParams{
		Nik:         arg.Nik,
		FullName:    arg.FullName,
		LegalName:   arg.LegalName,
		BirthPlace:  arg.BirthPlace,
		BirthDate:   arg.BirthDate,
		Salary:      arg.Salary,
		KTPPhoto:    arg.KTPPhoto,
		SelfiePhoto: arg.SelfiePhoto,
	}
}
