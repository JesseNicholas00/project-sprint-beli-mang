package auth

type Staff struct {
	Id        string `db:"staff_id"`
	Name      string `db:"staff_name"`
	Phone     string `db:"staff_phone_number"`
	Password  string `db:"staff_password"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
