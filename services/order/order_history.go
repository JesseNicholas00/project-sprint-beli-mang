package order

import (
	"context"
	"time"

	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	"github.com/JesseNicholas00/BeliMang/repos/order"
	"github.com/JesseNicholas00/BeliMang/types/location"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
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

	err = transaction.RunWithAutoCommit(&sess, func() error {
		var merchants []merchant.Merchant
		if req.MerchantCategory != nil || req.MerchantId != nil ||
			req.Name != nil {
			merchants, err := getMerchantsForFilters(
				svc,
				ctx,
				req.MerchantCategory,
				req.MerchantId,
				req.Name,
				[]string{},
			)
			if err != nil {
				return err
			}

			if len(merchants) == 0 {
				return nil
			}
		}

		var merchantIds []string
		nameFilter := req.Name
		if len(merchants) != 0 {
			merchantIds = getMerchantIds(merchants)
			if (req.MerchantCategory != nil || req.MerchantId != nil) &&
				req.Name != nil {
				nameFilter = nil
			}
		}

		summaryRes, err := svc.repo.ListOrderSummary(
			ctx,
			order.OrderSummaryListFilter{
				MerchantIds: merchantIds,
				Limit:       *req.Limit,
				Offset:      *req.Offset,
				ItemName:    nameFilter,
				UserId:      req.UserId,
			},
		)
		if err != nil {
			return err
		}

		if len(summaryRes) > 0 {
			if nameFilter != nil {
				merchantIds = getMerchantIdsFromSummary(summaryRes)
				merchants, err = getMerchantsForFilters(
					svc,
					ctx,
					nil,
					nil,
					nil,
					merchantIds,
				)
				if err != nil {
					return err
				}
			}
			merchantIdMap := make(map[string]OrderHistoryMerchantDetail)
			for _, merchant := range merchants {
				merchantIdMap[merchant.Id] = OrderHistoryMerchantDetail{
					MerchantId:       merchant.Id,
					Name:             merchant.Name,
					MerchantCategory: merchant.Category,
					ImageUrl:         merchant.ImageUrl,
					Location: location.Location{
						Latitude:  &merchant.Latitude,
						Longitude: &merchant.Longitude,
					},
					CreatedAt: merchant.CreatedAt.Format(time.RFC3339Nano),
				}
			}

			mappedOrderIdMap := make(map[string]*int)
			mappedOrderIdMerchantIdMap := make(map[string]map[string]*int)
			for _, summary := range summaryRes {
				if mappedOrderIdMap[summary.OrderID] == nil {
					currentSize := len(res.Result)
					mappedOrderIdMap[summary.OrderID] = &currentSize
					res.Result = append(res.Result, OrderHistoryDTO{
						OrderId: summary.OrderID,
					})
				}
				currOrderLoc := mappedOrderIdMap[summary.OrderID]

				if mappedOrderIdMerchantIdMap[summary.OrderID] == nil {
					mappedOrderIdMerchantIdMap[summary.OrderID] = make(
						map[string]*int,
					)
				}
				if mappedOrderIdMerchantIdMap[summary.OrderID][summary.MerchantID] == nil {
					currentSize := len(res.Result[*currOrderLoc].Orders)
					mappedOrderIdMerchantIdMap[summary.OrderID][summary.MerchantID] = &currentSize
				} else {
					res.Result[*currOrderLoc].Orders[*mappedOrderIdMerchantIdMap[summary.OrderID][summary.MerchantID]].Items = append(
						res.Result[*currOrderLoc].Orders[*mappedOrderIdMerchantIdMap[summary.OrderID][summary.MerchantID]].Items,
						OrderHistoryItemDetail{
							ItemId:          summary.MerchantItemID,
							Name:            summary.ItemName,
							ProductCategory: summary.ItemCategory,
							Price:           summary.ItemPrice,
							Quantity:        summary.Quantity,
							ImageUrl:        summary.ItemImageURL,
							CreatedAt:       summary.ItemCreatedAt.Format(time.RFC3339Nano),
						},
					)
				}
			}
		}

		return nil
	})
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return nil
}

func getMerchantsForFilters(svc *orderServiceImpl,
	ctx context.Context,
	merchantCategory, merchantId, merchantName *string,
	merchantIds []string,
) (res []merchant.Merchant, err error) {

	filter := merchant.MerchantListAllFilter{}
	if merchantCategory != nil {
		filter.MerchantCategory = merchantCategory
	}
	if merchantId != nil {
		filter.MerchantId = merchantId
	}
	if merchantName != nil {
		filter.Name = merchantName
	}
	filter.MerchantIds = merchantIds

	res, err = svc.merchantRepo.ListAllMerchant(ctx, filter)

	return
}

func getMerchantIds(merchants []merchant.Merchant) []string {
	merchantIds := []string{}

	for _, merchant := range merchants {
		merchantIds = append(merchantIds, merchant.Id)
	}

	return merchantIds
}

func getMerchantIdsFromSummary(
	orderSummaries []order.OrderSummaryView,
) []string {
	merchantIds := []string{}

	for _, orderSummary := range orderSummaries {
		merchantIds = append(merchantIds, orderSummary.MerchantID)
	}

	return merchantIds
}
