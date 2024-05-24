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
