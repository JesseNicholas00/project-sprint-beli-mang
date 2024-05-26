package order

import (
	"context"
	"errors"
	"fmt"
	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	"github.com/JesseNicholas00/BeliMang/repos/order"
	"github.com/JesseNicholas00/BeliMang/types/location"
	"github.com/JesseNicholas00/BeliMang/utils/helper"
	"github.com/JesseNicholas00/BeliMang/utils/unittesting"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestCreateEstimate(t *testing.T) {
	Convey("When creating an order estimate", t, func() {
		mockCtrl, service, mockedRepo, mockedMerchantRepo := NewWithMockedRepo(
			t,
		)
		defer mockCtrl.Finish()

		reqMerchantIds := []string{"gamer", "tsukuba", "bruh"}
		reqStartingPoints := []bool{false, false, true}
		reqItems := [][]MerchantOrderItem{
			{
				{
					ItemId:   "hakos",
					Quantity: 1,
				},
				{
					ItemId:   "baelz",
					Quantity: 2,
				},
			},
			{
				{
					ItemId:   "banban",
					Quantity: 127,
				},
				{
					ItemId:   "bigecho",
					Quantity: 427,
				},
				{
					ItemId:   "manekineko",
					Quantity: 67,
				},
			},
			{
				{
					ItemId:   "station",
					Quantity: 3,
				},
			},
		}

		var reqItemIds []string
		for _, items := range reqItems {
			for _, item := range items {
				reqItemIds = append(reqItemIds, item.ItemId)
			}
		}

		req := EstimateOrderReq{
			UserId: "five-fox-fubuki",
			UserLocation: location.Location{
				Latitude:  helper.ToPointer(36.1059758),
				Longitude: helper.ToPointer(140.1102648),
			},
			Orders: make([]MerchantOrder, 0),
		}
		for i := range reqMerchantIds {
			req.Orders = append(req.Orders, MerchantOrder{
				MerchantId:      reqMerchantIds[i],
				IsStartingPoint: &reqStartingPoints[i],
				Items:           reqItems[i],
			})
		}
		repoLats := []float64{36.106305, 36.1075609, 36.1057501}
		repoLongs := []float64{140.1066288, 140.1080447, 140.1090157}

		var repoMerchants []merchant.Merchant
		for i := range repoLats {
			repoMerchants = append(repoMerchants, merchant.Merchant{
				Id:        reqMerchantIds[i],
				Name:      "name of " + reqMerchantIds[i],
				Category:  "SmallRestaurant",
				ImageUrl:  "http://bread.com/bread.png",
				Latitude:  repoLats[i],
				Longitude: repoLongs[i],
			})
		}

		callWithMerchantRes := func(
			merchantRes []merchant.Merchant,
			merchantErr error,
		) (res EstimateOrderRes, err error) {
			mockedMerchantRepo.
				EXPECT().
				ListMerchantsByIds(gomock.Any(), reqMerchantIds).
				Return(merchantRes, merchantErr).
				Times(1)
			err = service.EstimateOrder(context.TODO(), req, &res)
			return
		}

		Convey("When merchant is not found", func() {
			Convey("Should return ErrMerchantNotFound", func() {
				_, err := callWithMerchantRes(nil, nil)
				So(errors.Is(err, ErrMerchantNotFound), ShouldBeTrue)
			})
		})

		var repoItems []merchant.MerchantItem
		for i := range reqMerchantIds {
			for j := range reqItems[i] {
				repoItems = append(repoItems, merchant.MerchantItem{
					Id:         reqItems[i][j].ItemId,
					MerchantId: reqMerchantIds[i],
					Name:       fmt.Sprintf("item-%d-%d", i, j),
					Category:   "category",
					Price:      69420,
					ImageUrl:   "https://bread.com/bread.png",
					CreatedAt:  time.Now().Add(-3 * time.Hour),
				})
			}
		}

		callWithItemRes := func(
			itemRes []merchant.MerchantItem,
			itemErr error,
		) (res EstimateOrderRes, err error) {
			mockedMerchantRepo.
				EXPECT().
				ListItemsByIds(gomock.Any(), reqItemIds).
				Return(itemRes, itemErr).
				Times(1)
			res, err = callWithMerchantRes(repoMerchants, nil)
			return
		}

		Convey("When item is not found", func() {
			Convey("When item doesn't exist", func() {
				Convey("Should return ErrItemNotFound", func() {
					_, err := callWithItemRes(nil, nil)
					So(errors.Is(err, ErrItemNotFound), ShouldBeTrue)
				})
			})
			Convey("When item belongs to the wrong merchant", func() {
				Convey("Should return ErrItemNotFound", func() {
					repoItems[len(repoItems)-1].MerchantId = reqMerchantIds[0]
					_, err := callWithItemRes(repoItems, nil)
					So(errors.Is(err, ErrItemNotFound), ShouldBeTrue)
				})
			})
		})

		callWithEverything := func() (res EstimateOrderRes, err error) {
			res, err = callWithItemRes(repoItems, nil)
			return
		}

		Convey("When distance is too far", func() {
			repoMerchants[0].Latitude = 35.7092181
			repoMerchants[0].Longitude = 139.7925196
			Convey("Should return ErrTooFar", func() {
				_, err := callWithEverything()
				So(errors.Is(err, ErrTooFar), ShouldBeTrue)
			})
		})

		Convey("When everything is good", func() {
			unittesting.FixNextUuid()
			repoEst := order.Estimate{
				Id:     uuid.NewString(),
				UserId: req.UserId,
				Items:  nil,
			}
			unittesting.FixNextUuid()

			for _, ord := range req.Orders {
				for _, item := range ord.Items {
					repoEst.Items = append(repoEst.Items, order.EstimateItem{
						MerchantId: ord.MerchantId,
						ItemId:     item.ItemId,
						Quantity:   item.Quantity,
					})
				}
			}

			mockedRepo.
				EXPECT().
				CreateEstimate(gomock.Any(), repoEst).
				Return(nil).
				Times(1)

			Convey("Should return the expected result", func() {
				res, err := callWithEverything()
				So(err, ShouldBeNil)
				So(res.Id, ShouldEqual, repoEst.Id)
				So(
					helper.IsEqualFloat(
						res.EstimatedDeliveryTimeMin,
						1.0172947900053564,
					),
					ShouldBeTrue,
				)
				So(res.TotalPrice, ShouldEqual, 43526340)
			})
		})
	})
}
