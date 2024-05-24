package role

type Role string

const (
	Admin = "admin"
	User  = "user"
)

func GetRole(isAdmin bool) Role {
	if isAdmin {
		return Admin
	}
	return User
}
