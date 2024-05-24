package merchant

type Merchant struct {
	Id        string  `db:"merchant_id"`
	Name      string  `db:"name"`
	Category  string  `db:"category"`
	ImageUrl  string  `db:"image_url"`
	Latitude  float32 `db:"latitude"`
	Longitude float32 `db:"longitude"`
}
