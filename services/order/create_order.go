package order

import (
	"context"
	"errors"

	"github.com/JesseNicholas00/BeliMang/repos/order"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/transaction"
	"github.com/google/uuid"
)

func (svc *orderServiceImpl) CreateOrder(
	ctx context.Context,
	req CreateOrderReq,
	res *CreateOrderRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := svc.dbRizzer.GetOrAppendTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return transaction.RunWithAutoCommit(&sess, func() error {
		estimate, err := svc.repo.FindEstimateById(ctx, req.EstimateOrderId)
		if err != nil {
			if errors.Is(err, order.ErrEstimateNotFound) {
				return ErrEstimateNotFound
			}
			return errorutil.AddCurrentContext(err)
		}

		// cannot use someone else's estimate
		if estimate.UserId != req.UserId {
			return ErrEstimateNotFound
		}

		orderId := uuid.NewString()

		if err := svc.repo.CreateOrder(ctx, order.Order{
			OrderId:    orderId,
			EstimateId: estimate.Id,
		}); err != nil {
			return errorutil.AddCurrentContext(err)
		}

		res.OrderId = orderId

		return nil
	})
}
