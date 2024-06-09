package order

import (
	"context"
	"time"

	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	"github.com/JesseNicholas00/BeliMang/repos/order"
	"github.com/JesseNicholas00/BeliMang/types/location"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/helper"
	"github.com/JesseNicholas00/BeliMang/utils/transaction"
)

// OrderHistory implements OrderService.
func (svc *orderServiceImpl) OrderHistory(ctx context.Context,
	req OrderHistoryReq,
	res *OrderHistoryRes,
) error {

	ctx, sess, err := svc.dbRizzer.GetOrAppendTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return transaction.RunWithAutoCommit(&sess, func() error {
		merchants, err := svc.merchantRepo.ListAllMerchant(
			ctx,
			merchant.MerchantListAllFilter{
				MerchantId:       req.MerchantId,
				Name:             req.Name,
				MerchantCategory: req.MerchantCategory,
			},
		)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		if len(merchants) == 0 {
			return nil
		}
		merchantDetailsById := make(map[string]OrderMerchant)
		for _, reqMerchant := range merchants {
			merchantDetailsById[reqMerchant.Id] = OrderMerchant{
				MerchantId:       reqMerchant.Id,
				Name:             reqMerchant.Name,
				MerchantCategory: reqMerchant.Category,
				ImageUrl:         reqMerchant.ImageUrl,
				Location: location.Location{
					Latitude:  helper.ToPointer(reqMerchant.Latitude),
					Longitude: helper.ToPointer(reqMerchant.Longitude),
				},
				CreatedAt: reqMerchant.CreatedAt.Format(time.RFC3339Nano),
			}
		}

		summaryReq := order.OrderSummaryListFilter{
			MerchantIds: nil,
			Limit:       *req.Limit,
			Offset:      *req.Offset,
			ItemName:    req.Name,
			UserId:      req.UserId,
		}
		for _, reqMerchant := range merchants {
			summaryReq.MerchantIds = append(
				summaryReq.MerchantIds,
				reqMerchant.Id,
			)
		}

		summaries, err := svc.repo.ListOrderSummary(ctx, summaryReq)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		type mapKey struct {
			orderId    string
			merchantId string
		}
		itemListByKey := make(map[mapKey][]OrderItem)

		for _, summary := range summaries {
			key := mapKey{
				summary.OrderID,
				summary.MerchantID,
			}

			curItem := OrderItem{
				ItemId:          summary.MerchantItemID,
				Name:            summary.ItemName,
				ProductCategory: summary.ItemCategory,
				Price:           summary.ItemPrice,
				Quantity:        summary.Quantity,
				ImageUrl:        summary.ItemImageURL,
				CreatedAt:       summary.ItemCreatedAt.Format(time.RFC3339Nano),
			}

			if itemList, ok := itemListByKey[key]; ok {
				itemList = append(itemList, curItem)
			} else {
				itemListByKey[key] = []OrderItem{curItem}
			}
		}

		ordersById := make(map[string][]Order)

		for key, items := range itemListByKey {
			curOrder := Order{
				Merchant: merchantDetailsById[key.merchantId],
				Items:    nil,
			}

			for _, item := range items {
				curOrder.Items = append(curOrder.Items, item)
			}

			if orders, ok := ordersById[key.orderId]; ok {
				orders = append(orders, curOrder)
			} else {
				ordersById[key.orderId] = []Order{curOrder}
			}
		}

		for id, orders := range ordersById {
			res.Entries = append(res.Entries, OrderHistoryEntry{
				OrderId: id,
				Orders:  orders,
			})
		}

		return nil
	})
}
