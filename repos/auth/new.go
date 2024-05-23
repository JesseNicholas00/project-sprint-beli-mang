package auth

import (
	"github.com/KerakTelor86/GoBoiler/utils/ctxrizz"
)

type authRepostioryImpl struct {
	dbRizzer ctxrizz.DbContextRizzer
}

func NewAuthRepository(dbRizzer ctxrizz.DbContextRizzer) AuthRepository {
	return &authRepostioryImpl{
		dbRizzer: dbRizzer,
	}
}
