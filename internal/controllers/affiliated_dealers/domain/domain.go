package domain

import "context"

type Usecase interface {
	Create(ctx context.Context, arg *CreateRequestParams) (int, error)
	Update(ctx context.Context, id int, arg *UpdateRequestParams) error
	FindByID(ctx context.Context, id int) (*AffiliatedDealer, error)
	List(ctx context.Context, arg *FetchParams) (*AffiliatedDealerList, error)
}
