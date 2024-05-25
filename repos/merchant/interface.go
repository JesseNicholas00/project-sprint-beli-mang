package merchant

import "context"

type MerchantRepository interface {
	CreateMerchant(ctx context.Context, m Merchant) error
	AdminListMerchant(
		ctx context.Context,
		filters AdminMerchantListFilter,
	) (res []Merchant, total int64, err error)
	CreateMerchantItem(ctx context.Context, mi MerchantItem) error
	FindMerchantById(ctx context.Context, merchant_id string) (Merchant, error)
	FindMerchantItemsByFilter(
		ctx context.Context,
		filters MerchantItemListFilter,
	) (res []MerchantItem, total int64, err error)
}
