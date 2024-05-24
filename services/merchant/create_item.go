package merchant

import (
	"context"
	"errors"

	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/transaction"
	"github.com/google/uuid"
)

func (svc *merchantServiceImpl) CreateMerchantItem(
	ctx context.Context,
	req CreateMerchantItemReq,
	res *CreateMerchantItemRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	ctx, sess, err := svc.dbRizzer.GetOrAppendTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return transaction.RunWithAutoCommit(&sess, func() error {
		_, err := svc.repo.FindMerchantById(ctx, req.MerchantId)

		if err != nil {
			if errors.Is(err, merchant.ErrMerchantNotFound) {
				return ErrMerchantNotFound
			}
			return errorutil.AddCurrentContext(err)
		}

		id := uuid.NewString()
		mi := merchant.MerchantItem{
			Id:         id,
			MerchantId: req.MerchantId,
			Name:       req.Name,
			Category:   req.Category,
			Price:      req.Price,
			ImageUrl:   req.ImageUrl,
		}

		err = svc.repo.CreateMerchantItem(ctx, mi)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		res.ItemId = id
		return nil
	})

}
