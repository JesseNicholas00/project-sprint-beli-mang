package merchant

import (
	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	"github.com/JesseNicholas00/BeliMang/utils/ctxrizz"
)

type merchantServiceImpl struct {
	repo     merchant.MerchantRepository
	dbRizzer ctxrizz.DbContextRizzer
}

func NewMerchantService(
	repo merchant.MerchantRepository,
	dbRizzer ctxrizz.DbContextRizzer,
) MerchantService {
	return &merchantServiceImpl{repo: repo, dbRizzer: dbRizzer}
}
