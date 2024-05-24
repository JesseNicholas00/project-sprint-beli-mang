package merchantitem

import (
	"context"

	"github.com/JesseNicholas00/BeliMang/repos/merchantitem"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/transaction"
	"github.com/google/uuid"
)

func (svc *merchantItemServiceImpl) CreateMerchantItem(
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
		_, err := svc.mRepo.FindMerchantById(ctx, req.MerchantId)

		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		id := uuid.NewString()
		mi := merchantitem.MerchantItem{
			Id:         id,
			MerchantId: req.MerchantId,
			Name:       req.Name,
			Category:   req.Category,
			Price:      req.Price,
			ImageUrl:   req.ImageUrl,
		}

		err = svc.miRepo.CreateMerchantItem(ctx, mi)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		res.ItemId = id
		return nil
	})

}
