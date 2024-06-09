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

type Order struct {
	OrderId    string `db:"order_id"`
	EstimateId string `db:"estimate_id"`
}

type OrderSummaryListFilter struct {
	MerchantIds []string
	Limit       int
	Offset      int
	ItemName    *string
	UserId      string
}

type OrderSummaryView struct {
	OrderID        string    `db:"order_id"`
	EstimateID     string    `db:"estimate_id"`
	UserID         string    `db:"user_id"`
	EstimateItemID int       `db:"estimate_item_id"`
	Quantity       int       `db:"quantity"`
	MerchantID     string    `db:"merchant_id"`
	MerchantItemID string    `db:"merchant_item_id"`
	ItemName       string    `db:"item_name"`
	ItemCategory   string    `db:"item_category"`
	ItemPrice      int64     `db:"item_price"`
	ItemImageURL   string    `db:"item_image_url"`
	ItemCreatedAt  time.Time `db:"item_created_at"`
}
