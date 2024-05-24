package merchantitem

import "github.com/JesseNicholas00/BeliMang/utils/ctxrizz"

type merchantItemRepoImpl struct {
	dbRizz     ctxrizz.DbContextRizzer
	statements statements
}

func NewMerchantItemRepository(dbRizz ctxrizz.DbContextRizzer) MerchantItemRepository {
	return &merchantItemRepoImpl{dbRizz: dbRizz, statements: prepareStatements()}
}
