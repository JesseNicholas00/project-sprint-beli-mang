package auth

import (
	"github.com/JesseNicholas00/BeliMang/utils/ctxrizz"
)

type authRepositoryImpl struct {
	dbRizzer ctxrizz.DbContextRizzer
}

func NewAuthRepository(dbRizzer ctxrizz.DbContextRizzer) AuthRepository {
	return &authRepositoryImpl{
		dbRizzer: dbRizzer,
	}
}
