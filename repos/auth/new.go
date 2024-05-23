package auth

import (
	"github.com/JesseNicholas00/BeliMang/utils/ctxrizz"
)

type authRepostioryImpl struct {
	dbRizzer ctxrizz.DbContextRizzer
}

func NewAuthRepository(dbRizzer ctxrizz.DbContextRizzer) AuthRepository {
	return &authRepostioryImpl{
		dbRizzer: dbRizzer,
	}
}
