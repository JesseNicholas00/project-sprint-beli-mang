package role

type Role int

const (
	Admin = iota
	User
)

func GetRole(isAdmin bool) Role {
	if isAdmin {
		return Admin
	}
	return User
}

func ToBoolean(role Role) bool {
	return role == Admin
}
