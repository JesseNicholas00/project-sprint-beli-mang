package merchantitem

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
