package merchant

import "github.com/JesseNicholas00/BeliMang/utils/ctxrizz"

type merchantRepoImpl struct {
	dbRizz     ctxrizz.DbContextRizzer
	statements statements
}

func NewMerchantRepository(dbRizz ctxrizz.DbContextRizzer) MerchantRepository {
	return &merchantRepoImpl{dbRizz: dbRizz, statements: prepareStatements()}
}
