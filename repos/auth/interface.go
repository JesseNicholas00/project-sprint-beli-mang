package auth

import "context"

type AuthRepository interface {
	CreateUser(ctx context.Context, user User) (User, error)
	FindUserByUsername(ctx context.Context, username string) (User, error)
	FindUserByEmailAndIsAdmin(ctx context.Context, email string, isAdmin bool) (User, error)
}
