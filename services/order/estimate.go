package order

import (
	"context"
	"math"

	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	"github.com/JesseNicholas00/BeliMang/repos/order"
	"github.com/JesseNicholas00/BeliMang/types/location"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/graphs"
	"github.com/JesseNicholas00/BeliMang/utils/helper"
	"github.com/JesseNicholas00/BeliMang/utils/transaction"
	"github.com/google/uuid"
)

const (
	speedKmh   = 40.0
	minPerHour = 60
	kmToMin    = 1 / speedKmh * minPerHour

	maxAcceptableOrderDistanceKm = 3.0
)

func (svc *orderServiceImpl) EstimateOrder(
	ctx context.Context,
	req EstimateOrderReq,
	res *EstimateOrderRes,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := svc.dbRizzer.GetOrAppendTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return transaction.RunWithAutoCommit(&sess, func() error {
		var merchantIds []string
		var itemIds []string

		for _, merchantOrder := range req.Orders {
			merchantIds = append(merchantIds, merchantOrder.MerchantId)
			for _, item := range merchantOrder.Items {
				itemIds = append(itemIds, item.ItemId)
			}
		}

		// check for merchant 404
		merchants, err := svc.merchantRepo.ListMerchantsByIds(
			ctx,
			merchantIds,
		)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}
		merchantById := make(map[string]merchant.Merchant)
		for _, m := range merchants {
			merchantById[m.Id] = m
		}
		for _, id := range merchantIds {
			if _, ok := merchantById[id]; !ok {
				return ErrMerchantNotFound
			}
		}

		// check for item 404
		items, err := svc.merchantRepo.ListItemsByIds(ctx, itemIds)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}
		itemByIds := make(map[string]merchant.MerchantItem)
		for _, it := range items {
			itemByIds[it.Id] = it
		}
		for _, merchantOrder := range req.Orders {
			for _, item := range merchantOrder.Items {
				// item not found at all (doesn't actually exist :skull:)
				repoItem, ok := itemByIds[item.ItemId]
				if !ok {
					return ErrItemNotFound
				}

				// item does not belong to expected merchant
				if repoItem.MerchantId != merchantOrder.MerchantId {
					return ErrItemNotFound
				}
			}
		}

		// check if distance is too big
		var coords []location.Location
		var startIdx int
		for _, merchantOrder := range req.Orders {
			if *merchantOrder.IsStartingPoint {
				startIdx = len(coords)
			}

			m := merchantById[merchantOrder.MerchantId]
			coords = append(coords, location.Location{
				Latitude:  helper.ToPointer(m.Latitude),
				Longitude: helper.ToPointer(m.Longitude),
			})
		}
		var endIdx = len(coords)
		coords = append(coords, req.UserLocation)

		var distMatrix = make([][]float64, len(coords))
		for i := range distMatrix {
			distMatrix[i] = make([]float64, len(coords))
			for j := range distMatrix[i] {
				distMatrix[i][j] = location.GetDistance(
					coords[i],
					coords[j],
				)
			}
		}

		radius := float64(0)
		for _, dist := range distMatrix[endIdx] {
			radius = max(radius, dist)
		}

		if radius*radius*math.Pi >= maxAcceptableOrderDistanceKm {
			return ErrTooFar
		}

		totalDistance, err := graphs.GetTspDistance(
			ctx,
			startIdx,
			endIdx,
			distMatrix,
		)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		totalPrice := int64(0)

		var estItems []order.EstimateItem
		for _, merchantOrder := range req.Orders {
			for _, item := range merchantOrder.Items {
				estItems = append(estItems, order.EstimateItem{
					MerchantId: merchantOrder.MerchantId,
					ItemId:     item.ItemId,
					Quantity:   item.Quantity,
				})
				totalPrice += int64(
					item.Quantity,
				) * itemByIds[item.ItemId].Price
			}
		}

		est := order.Estimate{
			Id:     uuid.NewString(),
			UserId: req.UserId,
			Items:  estItems,
		}
		if err := svc.repo.CreateEstimate(ctx, est); err != nil {
			return errorutil.AddCurrentContext(err)
		}

		*res = EstimateOrderRes{
			Id:                       est.Id,
			EstimatedDeliveryTimeMin: totalDistance * kmToMin,
			TotalPrice:               totalPrice,
		}
		return nil
	})
}
