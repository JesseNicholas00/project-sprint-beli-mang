package merchant

import "context"

type MerchantRepository interface {
	CreateMerchant(ctx context.Context, m Merchant) error
	AdminListMerchant(
		ctx context.Context,
		filters AdminMerchantListFilter,
	) (res []Merchant, total int64, err error)
	ListAllMerchant(
		ctx context.Context,
		filters MerchantListAllFilter,
	) (res []Merchant, err error)
	CreateMerchantItem(ctx context.Context, mi MerchantItem) error
	FindMerchantById(ctx context.Context, merchantId string) (Merchant, error)
	FindMerchantByFilter(
		ctx context.Context,
		filter MerchantFilter,
	) (res []MerchantWithItems, err error)
	ListMerchantsByIds(
		ctx context.Context,
		merchantIds []string,
	) ([]Merchant, error)
	ListItemsByIds(ctx context.Context, ids []string) ([]MerchantItem, error)
}
