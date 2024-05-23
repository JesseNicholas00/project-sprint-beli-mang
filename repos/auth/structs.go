package auth

type User struct {
	Id       string `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
	IsAdmin  bool   `db:"is_admin"`
}
