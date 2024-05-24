package merchantitem

import (
	"context"
)

type MerchantItemService interface {
	CreateMerchantItem(
		ctx context.Context,
		req CreateMerchantItemReq,
		res *CreateMerchantItemRes,
	) error
}
