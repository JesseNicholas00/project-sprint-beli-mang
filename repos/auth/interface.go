package auth

import "context"

type AuthRepository interface {
	CreateStaff(ctx context.Context, staff Staff) (Staff, error)
	FindStaffByPhone(ctx context.Context, phone string) (Staff, error)
}
