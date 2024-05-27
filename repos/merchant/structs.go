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
