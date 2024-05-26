package merchant

import "context"

type MerchantRepository interface {
	CreateMerchant(ctx context.Context, m Merchant) error
	CreateMerchantItem(ctx context.Context, mi MerchantItem) error
	FindMerchantById(ctx context.Context, merchant_id string) (Merchant, error)
	FindMerchantByFilter(ctx context.Context, filter MerchantFilter) (res []Merchant, err error)
}
