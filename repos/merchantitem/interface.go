package merchantitem

import "context"

type MerchantItemRepository interface {
	CreateMerchantItem(ctx context.Context, mi MerchantItem) error
}
