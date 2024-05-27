package merchant

import (
	"context"

	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	"github.com/JesseNicholas00/BeliMang/types/location"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	merchantCategory "github.com/JesseNicholas00/BeliMang/utils/validation/merchant"
)

func (svc *merchantServiceImpl) FindMerchantByFilter(
	ctx context.Context,
	req FindMerchantReq,
	res *FindMerchantRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	var err error
	if req.MerchantCategory != nil {
		validCategory := merchantCategory.IsValidMerchantCategory(*req.MerchantCategory)

		if !validCategory {
			res.Data = []MerchantAndItems{}
			return err
		}
	}

	filter := merchant.MerchantFilter{
		Name:             req.Name,
		Limit:            *req.Limit,
		Offset:           req.Offset,
		MerchantCategory: req.MerchantCategory,
		MerchantId:       req.MerchantId,
		Location:         req.Location,
	}

	merchants, err := svc.repo.FindMerchantByFilter(ctx, filter)

	var kaiCenat []MerchantAndItems

	for _, merchant := range merchants {
		merch := Merchant{
			MerchantId:       merchant.Id,
			Name:             merchant.Name,
			MerchantCategory: merchant.Category,
			ImageUrl:         merchant.ImageUrl,
			Location: location.Location{
				Latitude:  &merchant.Latitude,
				Longitude: &merchant.Longitude,
			},
			CreatedAt: merchant.CreatedAt,
		}
		var items []Item

		for _, item := range merchant.Items {
			merchItem := Item{
				ItemId:          item.Id,
				Name:            item.Name,
				ImageUrl:        item.ImageUrl,
				Price:           item.Price,
				ProductCategory: item.Category,
				CreatedAt:       item.CreatedAt,
			}
			items = append(items, merchItem)
		}

		kaiCenat = append(kaiCenat, MerchantAndItems{
			Merchant: merch,
			Items:    items,
		})
	}

	if err != nil {
		return errorutil.AddCurrentContext(err)
	}
	res.Data = kaiCenat

	return err
}
