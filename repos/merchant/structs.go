package merchant

type Merchant struct {
	Id        string  `db:"merchant_id"`
	Name      string  `db:"name"`
	Category  string  `db:"category"`
	ImageUrl  string  `db:"image_url"`
	Latitude  float64 `db:"latitude"`
	Longitude float64 `db:"longitude"`
}
