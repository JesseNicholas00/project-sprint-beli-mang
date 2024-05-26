package merchant

import (
	"context"
	"strconv"
	"strings"

	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	"github.com/JesseNicholas00/BeliMang/types/location"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/helper"
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
		_, exists := svc.categories[*req.MerchantCategory]

		if !exists {
			res.Data = []MerchantAndItems{}
			return err
		}
	}

	fanum_tax := strings.Split(req.LatLongSeparatedByCommaIdkWhyItsLikeThatButOk, ",")
	var fanum_tax_paid location.Location

	if len(fanum_tax) == 2 {
		fanum, err_fanum := strconv.ParseFloat(fanum_tax[0], 64)
		if err_fanum != nil {
			err = errorutil.AddCurrentContext(err_fanum)
		}
		tax, err_tax := strconv.ParseFloat(fanum_tax[0], 64)
		if err_tax != nil {
			err = errorutil.AddCurrentContext(err_tax)
		}
		fanum_tax_paid.Latitude = &fanum
		fanum_tax_paid.Longitude = &tax
	} else {
		err = ErrLatLangNotValid
		return err
	}

	if req.Limit == nil {
		req.Limit = helper.ToPointer(5)
	}
	filter := merchant.MerchantFilter{
		Name:             req.Name,
		Limit:            *req.Limit,
		Offset:           req.Offset,
		MerchantCategory: req.MerchantCategory,
		MerchantId:       req.MerchantId,
		Location:         fanum_tax_paid,
	}

	merchants, err := svc.repo.FindMerchantByFilter(ctx, filter)

	var kaiCenat []MerchantAndItems

	for _, merchanta := range merchants {
		merch := Merchant{
			MerchantId:       merchanta.Id,
			Name:             merchanta.Name,
			MerchantCategory: merchanta.Category,
			ImageUrl:         merchanta.ImageUrl,
			Location: location.Location{
				Latitude:  &merchanta.Latitude,
				Longitude: &merchanta.Longitude,
			},
			CreatedAt: merchanta.CreatedAt,
		}
		var items []Item

		for _, item := range merchanta.Items {
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
