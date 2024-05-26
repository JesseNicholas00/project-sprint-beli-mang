package merchant

import (
	"time"

	"github.com/JesseNicholas00/BeliMang/types/location"
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
	MerchantId string `validate:"required,min=1"           param:"merchantId"`
	Name       string `validate:"required,min=2,max=30"                       json:"name"`
	Category   string `validate:"required,productCategory"                    json:"productCategory"`
	Price      int64  `validate:"required,min=1"                              json:"price"`
	ImageUrl   string `validate:"required,url,imageExt"                       json:"imageUrl"`
}

type CreateMerchantItemRes struct {
	ItemId string `json:"itemId"`
}

type Merchant struct {
	MerchantId       string            `json:"merchantId"`
	Name             string            `json:"name"`
	MerchantCategory string            `json:"merchantCategory"`
	ImageUrl         string            `json:"imageUrl"`
	Location         location.Location `json:"location"`
	CreatedAt        time.Time         `json:"createdAt"`
}

type Item struct {
	ItemId          string    `json:"itemId"`
	Name            string    `json:"name"`
	ProductCategory string    `json:"productCategory"`
	Price           int       `json:"price"`
	ImageUrl        string    `json:"imageUrl"`
	CreatedAt       time.Time `json:"createdAt"`
}

type MerchantAndItems struct {
	Merchant Merchant `json:"merchant"`
	Items    []Item   `json:"items"`
}

type FindMerchantReq struct {
	LatLongSeparatedByCommaIdkWhyItsLikeThatButOk string  `param:"location"             validate:"required,min=3"`
	MerchantId                                    *string `query:"merchantId"`
	Name                                          *string `query:"name"`
	MerchantCategory                              *string `query:"category"`
	Limit                                         *int    `query:"limit"`
	Offset                                        int     `query:"offset"`
}

type FindMerchantRes struct {
	Data []MerchantAndItems `json:"data"`
}
