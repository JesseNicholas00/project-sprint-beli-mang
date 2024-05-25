package merchant

import "time"

type Merchant struct {
	Id        string  `db:"merchant_id"`
	Name      string  `db:"name"`
	Category  string  `db:"category"`
	ImageUrl  string  `db:"image_url"`
	Latitude  float64 `db:"latitude"`
	Longitude float64 `db:"longitude"`
}

type MerchantItem struct {
	Id         string    `db:"merchant_item_id"`
	MerchantId string    `db:"merchant_id"`
	Name       string    `db:"name"`
	Category   string    `db:"category"`
	Price      int       `db:"price"`
	ImageUrl   string    `db:"image_url"`
	CreatedAt  time.Time `db:"created_at"`
}
