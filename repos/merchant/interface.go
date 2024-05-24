package merchant

import "context"

type MerchantRepository interface {
	CreateMerchant(ctx context.Context, m Merchant) error
}
