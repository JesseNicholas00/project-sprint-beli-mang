package merchant

import "context"

type MerchantRepository interface {
	CreateMerchant(ctx context.Context, m Merchant) error
	AdminListMerchant(ctx context.Context, filters AdminMerchantListFilter) (res []Merchant, total int64, err error)
}
