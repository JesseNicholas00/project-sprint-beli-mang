package merchant

import "time"

type Merchant struct {
	Id        string    `db:"merchant_id"`
	Name      string    `db:"name"`
	Category  string    `db:"category"`
	ImageUrl  string    `db:"image_url"`
	Latitude  float64   `db:"latitude"`
	Longitude float64   `db:"longitude"`
	CreatedAt time.Time `db:"created_at"`
}

type Total struct {
	Total int64 `db:"total"`
}

type AdminMerchantListFilter struct {
	MerchantId       *string
	Limit            int
	Offset           int
	Name             *string
	MerchantCategory *string
	CreatedAtSort    *string
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

type MerchantItemListFilter struct {
	MerchantItemId *string
	Limit          int
	Offset         int
	Name           *string
	Category       *string
	CreatedAtSort  *string
}
