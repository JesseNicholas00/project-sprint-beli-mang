package merchantitem

import (
	merchant "github.com/JesseNicholas00/BeliMang/repos/merchant"
	merchantitem "github.com/JesseNicholas00/BeliMang/repos/merchantitem"
)

type merchantItemServiceImpl struct {
	mRepo  merchant.MerchantRepository
	miRepo merchantitem.MerchantItemRepository
}

func NewMerchantItemService(
	mRepo merchant.MerchantRepository,
	miRepo merchantitem.MerchantItemRepository,
) MerchantItemService {
	return &merchantItemServiceImpl{miRepo: miRepo, mRepo: mRepo}
}
