package order

import "github.com/JesseNicholas00/BeliMang/types/location"

type EstimateOrderReq struct {
	UserId       string
	UserLocation location.Location `json:"userLocation" validate:"required"`
	// validate len <= 16 to prevent our TSP solver from blowing up
	Orders []MerchantOrder `json:"orders"       validate:"required,max=16,dive"`
}

type MerchantOrder struct {
	MerchantId      string              `json:"merchantId"      validate:"required"`
	IsStartingPoint *bool               `json:"isStartingPoint" validate:"required"`
	Items           []MerchantOrderItem `json:"items"           validate:"required,dive"`
}

type MerchantOrderItem struct {
	ItemId   string `json:"itemId"   validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
}

type EstimateOrderRes struct {
	TotalPrice               int64   `json:"totalPrice"`
	EstimatedDeliveryTimeMin float64 `json:"estimatedDeliveryTimeInMinutes"`
	Id                       string  `json:"calculatedEstimateId"`
}

type CreateOrderReq struct {
	UserId          string
	EstimateOrderId string `json:"calculatedEstimateId" validate:"required"`
}

type CreateOrderRes struct {
	OrderId string `json:"orderId"`
}

type OrderHistoryReq struct {
	MerchantId       *string `query:"merchantId"`
	Limit            *int    `query:"limit"`
	Offset           *int    `query:"offset"`
	Name             *string `query:"name"`
	MerchantCategory *string `query:"merchantCategory"`
	UserId           string
}

type OrderHistoryRes struct {
	Entries []OrderHistoryEntry
}

type OrderHistoryEntry struct {
	OrderId string  `json:"orderId"`
	Orders  []Order `json:"orders"`
}

type Order struct {
	Merchant OrderMerchant `json:"merchant"`
	Items    []OrderItem   `json:"items"`
}

type OrderMerchant struct {
	MerchantId       string            `json:"merchantId"`
	Name             string            `json:"name"`
	MerchantCategory string            `json:"merchantCategory"`
	ImageUrl         string            `json:"imageUrl"`
	Location         location.Location `json:"location"`
	CreatedAt        string            `json:"createdAt"`
}

type OrderItem struct {
	ItemId          string `json:"itemId"`
	Name            string `json:"name"`
	ProductCategory string `json:"productCategory"`
	Price           int64  `json:"price"`
	Quantity        int    `json:"quantity"`
	ImageUrl        string `json:"imageUrl"`
	CreatedAt       string `json:"createdAt"`
}
