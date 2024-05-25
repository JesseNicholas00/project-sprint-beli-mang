package merchant

import "github.com/JesseNicholas00/BeliMang/types/location"

type CreateMerchantReq struct {
	Name     string            `json:"name"             validate:"required,min=2,max=30"`
	Category string            `json:"merchantCategory" validate:"required,merchantCategory"`
	ImageUrl string            `json:"imageUrl"         validate:"required,url,imageExt"`
	Location location.Location `json:"location"         validate:"required"`
}

type CreateMerchantRes struct {
	Id string `json:"merchantId"`
}

type CreateMerchantItemReq struct {
	MerchantId string `validate:"required,min=1"           param:"merchantId"`
	Name       string `validate:"required,min=2,max=30"                       json:"name"`
	Category   string `validate:"required,productCategory"                    json:"productCategory"`
	Price      int    `validate:"required,min=1"                              json:"price"`
	ImageUrl   string `validate:"required,url,imageExt"                       json:"imageUrl"`
}

type CreateMerchantItemRes struct {
	ItemId string `json:"itemId"`
}
