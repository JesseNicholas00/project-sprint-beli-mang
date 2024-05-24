package merchantitem

import (
	merchant "github.com/JesseNicholas00/BeliMang/repos/merchant"
	merchantitem "github.com/JesseNicholas00/BeliMang/repos/merchantitem"
	"github.com/JesseNicholas00/BeliMang/utils/ctxrizz"
)

type merchantItemServiceImpl struct {
	mRepo    merchant.MerchantRepository
	miRepo   merchantitem.MerchantItemRepository
	dbRizzer ctxrizz.DbContextRizzer
}

func NewMerchantItemService(
	mRepo merchant.MerchantRepository,
	miRepo merchantitem.MerchantItemRepository,
	dbRizzer ctxrizz.DbContextRizzer,
) MerchantItemService {
	return &merchantItemServiceImpl{
		miRepo:   miRepo,
		mRepo:    mRepo,
		dbRizzer: dbRizzer,
	}
}
