package merchant

import (
	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	"github.com/JesseNicholas00/BeliMang/utils/ctxrizz"
)

type merchantServiceImpl struct {
	repo       merchant.MerchantRepository
	dbRizzer   ctxrizz.DbContextRizzer
	categories map[string]string
}

func NewMerchantService(
	repo merchant.MerchantRepository,
	dbRizzer ctxrizz.DbContextRizzer,
) MerchantService {
	// gyatt categories enum
	categories := make(map[string]string)
	categories["SmallRestaurant"] = "SmallRestaurant"
	categories["MediumRestaurant"] = "MediumRestaurant"
	categories["LargeRestaurant"] = "LargeRestaurant"
	categories["MerchandiseRestaurant"] = "MerchandiseRestaurant"
	categories["BoothKiosk"] = "BoothKiosk"
	categories["ConvenienceStore"] = "ConvenienceStore"

	return &merchantServiceImpl{repo: repo, dbRizzer: dbRizzer, categories: categories}
}
