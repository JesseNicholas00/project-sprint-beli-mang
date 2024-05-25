package order

import "time"

type Estimate struct {
	Id        string    `db:"estimate_id"`
	UserId    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	Items     []EstimateItem
}

type EstimateItem struct {
	MerchantId string `db:"merchant_id"`
	ItemId     string `db:"merchant_item_id"`
	Quantity   int    `db:"quantity"`
}
