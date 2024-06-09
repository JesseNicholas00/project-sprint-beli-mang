package merchant

import (
	"context"
	"time"

	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	"github.com/JesseNicholas00/BeliMang/types/pagination"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
)

func (svc *merchantServiceImpl) FindMerchantItemList(ctx context.Context,
	req MerchantItemListReq,
	res *MerchantItemListRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	repoRes, totalRes, err := svc.repo.FindMerchantItemsByFilter(
		ctx,
		merchant.MerchantItemListFilter{
			MerchantId:     req.MerchantId,
			MerchantItemId: req.MerchantItemId,
			Limit:          *req.Limit,
			Offset:         *req.Offset,
			Name:           req.Name,
			Category:       req.Category,
			CreatedAtSort:  req.CreatedAtSort,
		},
	)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	for _, merchantItem := range repoRes {
		res.Data = append(res.Data, ListMerchantItemResData{
			MerchantItemId: merchantItem.Id,
			Name:           merchantItem.Name,
			Price:          merchantItem.Price,
			Category:       merchantItem.Category,
			ImageUrl:       merchantItem.ImageUrl,
			CreatedAt:      merchantItem.CreatedAt.Format(time.RFC3339Nano),
		})
	}
	res.Meta = pagination.Page{
		Limit:  req.Limit,
		Offset: req.Offset,
		Total:  &totalRes,
	}

	return nil
}
