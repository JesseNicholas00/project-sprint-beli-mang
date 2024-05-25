package order

type Estimate struct {
	Id     string `db:"estimate_id"`
	UserId string `db:"user_id"`
	Items  []EstimateItem
}

type EstimateItem struct {
	MerchantId string `db:"merchant_id"`
	ItemId     string `db:"merchant_item_id"`
	Quantity   int    `db:"quantity"`
}

type dbEstimateItem struct {
	EstimateItem
	Id      int64  `db:"order_item_id"`
	OrderId string `db:"order_id"`
}
