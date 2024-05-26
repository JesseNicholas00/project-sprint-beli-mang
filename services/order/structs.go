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
