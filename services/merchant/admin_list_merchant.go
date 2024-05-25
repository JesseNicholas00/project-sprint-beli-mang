package merchant

import (
	"context"
	"time"

	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	"github.com/JesseNicholas00/BeliMang/types/location"
	"github.com/JesseNicholas00/BeliMang/types/pagination"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
)

// AdminListMerchant implements MerchantService.
func (svc *merchantServiceImpl) AdminListMerchant(
	ctx context.Context,
	req AdminListMerchantReq, res *AdminListMerchantRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	repoRes, totalRes, err := svc.repo.AdminListMerchant(ctx, merchant.AdminMerchantListFilter{
		MerchantId:       req.MerchantId,
		Limit:            *req.Limit,
		Offset:           *req.Offset,
		Name:             req.Name,
		MerchantCategory: req.MerchantCategory,
		CreatedAtSort:    req.CreatedAtSort,
	})
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	for _, merchant := range repoRes {
		res.Data = append(res.Data, ListMerchantResData{
			MerchantId:       merchant.Id,
			Name:             merchant.Name,
			MerchantCategory: merchant.Category,
			ImageUrl:         merchant.ImageUrl,
			Location: location.Location{
				Latitude:  &merchant.Latitude,
				Longitude: &merchant.Longitude,
			},
			CreatedAt: merchant.CreatedAt.Format(time.RFC3339Nano),
		})
	}
	res.Meta = pagination.Page{
		Limit:  req.Limit,
		Offset: req.Offset,
		Total:  &totalRes,
	}

	return nil
}
