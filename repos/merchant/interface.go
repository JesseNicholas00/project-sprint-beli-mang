package merchant

import "context"

type MerchantRepository interface {
	CreateMerchant(ctx context.Context, m Merchant) error
	FindMerchantById(ctx context.Context, merchant_id string) (Merchant, error)
}
