package merchant

import (
	"time"

	"github.com/JesseNicholas00/BeliMang/types/location"
)

type Merchant struct {
	Id        string    `db:"merchant_id"`
	Name      string    `db:"name"`
	Category  string    `db:"category"`
	ImageUrl  string    `db:"image_url"`
	Latitude  float64   `db:"latitude"`
	Longitude float64   `db:"longitude"`
	CreatedAt time.Time `db:"created_at"`
}

type AdminMerchantListFilter struct {
	MerchantId       *string
	Limit            int
	Offset           int
	Name             *string
	MerchantCategory *string
	CreatedAtSort    *string
}

type MerchantFilter struct {
	MerchantId       *string
	Name             *string
	MerchantCategory *string
	Location         location.GyattLocation
	Limit            int
	Offset           int
}

type MerchantListAllFilter struct {
	MerchantId       *string
	MerchantIds      []string
	Name             *string
	MerchantCategory *string
}

type MerchantItem struct {
	Id         string    `db:"merchant_item_id"`
	MerchantId string    `db:"merchant_id"`
	Name       string    `db:"name"`
	Category   string    `db:"category"`
	Price      int64     `db:"price"`
	ImageUrl   string    `db:"image_url"`
	CreatedAt  time.Time `db:"created_at"`
}

type MerchantItemDetail struct {
	Id         string    `db:"item_id"`
	MerchantId string    `db:"item_merchant_id"`
	Name       string    `db:"item_name"`
	Category   string    `db:"item_category"`
	Price      int64     `db:"item_price"`
	ImageUrl   string    `db:"item_image_url"`
	CreatedAt  time.Time `db:"item_created_at"`
}

type MerchantWithItems struct {
	Id        string    `db:"merchant_id"`
	Name      string    `db:"merchant_name"`
	Category  string    `db:"merchant_category"`
	ImageUrl  string    `db:"merchant_image_url"`
	Latitude  float64   `db:"latitude"`
	Longitude float64   `db:"longitude"`
	CreatedAt time.Time `db:"merchant_created_at"`
	Items     []MerchantItemDetail
}

type MerchantItemListFilter struct {
	MerchantItemId *string
	Limit          int
	Offset         int
	Name           *string
	Category       *string
	CreatedAtSort  *string
}

type Total struct {
	Total int64 `db:"total"`
}
