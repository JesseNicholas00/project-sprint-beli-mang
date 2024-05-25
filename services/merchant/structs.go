package merchant

import (
	"github.com/JesseNicholas00/BeliMang/types/location"
	"github.com/JesseNicholas00/BeliMang/types/pagination"
)

type CreateMerchantReq struct {
	Name     string            `json:"name"             validate:"required,min=2,max=30"`
	Category string            `json:"merchantCategory" validate:"required,merchantCategory"`
	ImageUrl string            `json:"imageUrl"         validate:"required,url,imageExt"`
	Location location.Location `json:"location"         validate:"required"`
}

type CreateMerchantRes struct {
	Id string `json:"merchantId"`
}

type AdminListMerchantReq struct {
	MerchantId       *string `query:"merchantId"`
	Limit            *int    `query:"limit"`
	Offset           *int    `query:"offset"`
	Name             *string `query:"name"`
	MerchantCategory *string `query:"merchantCategory"`
	CreatedAtSort    *string `query:"createdAt"`
}

type AdminListMerchantRes struct {
	Data []ListMerchantResData `json:"data"`
	Meta pagination.Page       `json:"meta"`
}

type ListMerchantResData struct {
	MerchantId       string            `json:"merchantId"`
	Name             string            `json:"name"`
	MerchantCategory string            `json:"merchantCategory"`
	ImageUrl         string            `json:"imageUrl"`
	Location         location.Location `json:"location"`
	CreatedAt        string            `json:"createdAt"`
}

type CreateMerchantItemReq struct {
	MerchantId string `validate:"required,min=1"            param:"merchantId"`
	Name       string `validate:"required,min=2,max=30"                        json:"name"`
	Category   string `validate:"required,merchantCategory"                    json:"productCategory"`
	Price      int    `validate:"required,min=1"                               json:"price"`
	ImageUrl   string `validate:"required,url,imageExt"                        json:"imageUrl"`
}

type CreateMerchantItemRes struct {
	ItemId string `json:"itemId"`
}

type MerchantItemListReq struct {
	MerchantItemId *string `query:"itemId"`
	Limit          *int    `query:"limit"`
	Offset         *int    `query:"offset"`
	Name           *string `query:"name"`
	Category       *string `query:"productCategory"`
	CreatedAtSort  *string `query:"createdAt"`
}

type ListMerchantItemResData struct {
	MerchantItemId string `json:"itemId"`
	Name           string `json:"name"`
	Category       string `json:"productCategory"`
	Price          int    `json:"price"`
	ImageUrl       string `json:"imageUrl"`
	CreatedAt      string `json:"createdAt"`
}

type MerchantItemListRes struct {
	Data []ListMerchantItemResData `json:"data"`
	Meta pagination.Page           `json:"meta"`
}
