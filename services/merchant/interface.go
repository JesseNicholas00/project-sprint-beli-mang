package merchant

import (
	"context"
)

type MerchantService interface {
	CreateMerchant(
		ctx context.Context,
		req CreateMerchantReq,
		res *CreateMerchantRes,
	) error
	AdminListMerchant(
		ctx context.Context,
		req AdminListMerchantReq,
		res *AdminListMerchantRes,
	) error
	CreateMerchantItem(
		ctx context.Context,
		req CreateMerchantItemReq,
		res *CreateMerchantItemRes,
	) error
}
