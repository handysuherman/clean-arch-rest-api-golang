package domain

import "context"

type Usecase interface {
	Create(ctx context.Context, arg *CreateRequestParams) (int64, error)
	Update(ctx context.Context, id int64, arg *UpdateRequestParams) (int64, error)
	FindByID(ctx context.Context, id int64) (*ConsumerLoanLimit, error)
	List(ctx context.Context, arg *FetchParams) (*ConsumerLoanLimitList, error)
}
