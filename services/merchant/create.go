package merchant

import (
	"context"
	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/google/uuid"
)

func (svc *merchantServiceImpl) CreateMerchant(
	ctx context.Context,
	req CreateMerchantReq,
	res *CreateMerchantRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	id := uuid.NewString()
	m := merchant.Merchant{
		Id:        id,
		Name:      req.Name,
		Category:  req.Category,
		ImageUrl:  req.ImageUrl,
		Latitude:  *req.Location.Latitude,
		Longitude: *req.Location.Longitude,
	}

	err := svc.repo.CreateMerchant(ctx, m)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	res.Id = id
	return nil
}
