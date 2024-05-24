package merchant

import (
	"github.com/JesseNicholas00/BeliMang/repos/merchant"
)

type merchantServiceImpl struct {
	repo merchant.MerchantRepository
}

func NewMerchantService(repo merchant.MerchantRepository) MerchantService {
	return &merchantServiceImpl{repo: repo}
}
