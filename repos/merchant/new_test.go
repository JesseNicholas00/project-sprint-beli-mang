package merchant_test

import (
	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	"github.com/JesseNicholas00/BeliMang/utils/ctxrizz"
	"github.com/JesseNicholas00/BeliMang/utils/unittesting"
	"testing"
)

func NewWithTestDatabase(t *testing.T) merchant.MerchantRepository {
	db := unittesting.SetupTestDatabase("../../migrations", t)
	return merchant.NewMerchantRepository(ctxrizz.NewDbContextRizzer(db))
}
