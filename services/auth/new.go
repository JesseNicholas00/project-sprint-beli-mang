package auth

import "github.com/KerakTelor86/GoBoiler/repos/auth"

type authServiceImpl struct {
	repo       auth.AuthRepository
	jwtSecret  []byte
	bcryptCost int
}

func NewAuthService(
	repo auth.AuthRepository,
	jwtSecret string,
	bcryptCost int,
) AuthService {
	return &authServiceImpl{
		repo:       repo,
		jwtSecret:  []byte(jwtSecret),
		bcryptCost: bcryptCost,
	}
}
