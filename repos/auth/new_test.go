package auth_test

import (
	"testing"

	"github.com/KerakTelor86/GoBoiler/repos/auth"
	"github.com/KerakTelor86/GoBoiler/utils/ctxrizz"
	"github.com/KerakTelor86/GoBoiler/utils/unittesting"
)

func NewWithTestDatabase(t *testing.T) auth.AuthRepository {
	db := unittesting.SetupTestDatabase("../../migrations", t)
	return auth.NewAuthRepository(ctxrizz.NewDbContextRizzer(db))
}
