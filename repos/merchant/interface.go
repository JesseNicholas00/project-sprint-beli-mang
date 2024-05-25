package merchant

import "context"

type MerchantRepository interface {
	CreateMerchant(ctx context.Context, m Merchant) error
	CreateMerchantItem(ctx context.Context, mi MerchantItem) error
	FindMerchantById(ctx context.Context, merchantId string) (Merchant, error)
	ListMerchantsByIds(
		ctx context.Context,
		merchantIds []string,
	) ([]Merchant, error)
	ListItemsByIds(ctx context.Context, ids []string) ([]MerchantItem, error)
}
