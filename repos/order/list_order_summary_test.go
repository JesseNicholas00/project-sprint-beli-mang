//go:build integration
// +build integration

package order_test

import (
	"context"
	"testing"
	"time"

	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	"github.com/JesseNicholas00/BeliMang/repos/order"
	. "github.com/smartystreets/goconvey/convey"
)

func TestListOrderSummary(t *testing.T) {
	Convey("With dummy data", t, func() {
		repo, merchantRepo := NewWithTestDatabase(t)

		merchantIds := []string{"220132512342", "220132512344", "220132512343"}
		names := []string{"epic", "epic_gamer_2", "gamer3"}
		merchantCategory := []string{"goodCategory", "badCategory", "lolMewGyattRizz"}
		lat := []float64{1.1, 3.3, 5.5}
		lon := []float64{2.2, 4.4, 6.6}
		createtAts := []time.Time{time.Now().AddDate(0, 0, -2), time.Now().AddDate(0, 0, -1), time.Now()}

		var merchants []merchant.Merchant
		for i := 0; i < 3; i++ {
			merchants = append(
				merchants,
				merchant.Merchant{
					Name:      names[i],
					Id:        merchantIds[i],
					Category:  merchantCategory[i],
					Latitude:  lat[i],
					Longitude: lon[i],
					ImageUrl:  "https://bread.com/bread.png",
					CreatedAt: createtAts[i],
				},
			)
		}

		for _, merchant := range merchants {
			err := merchantRepo.CreateMerchant(context.TODO(), merchant)
			So(err, ShouldBeNil)
		}

		items := []merchant.MerchantItem{
			{
				Id:         "bread1",
				MerchantId: merchantIds[0],
				Name:       "goodName1",
				Category:   "goodCat1",
				Price:      1000,
				ImageUrl:   "https://lmao.image/1",
				CreatedAt:  time.Now().Add(time.Duration(-3)),
			},
			{
				Id:         "bread2",
				MerchantId: merchantIds[0],
				Name:       "goodName2",
				Category:   "goodCat2",
				Price:      2000,
				ImageUrl:   "https://lmao.image/2",
				CreatedAt:  time.Now().Add(time.Duration(-2)),
			},
			{
				Id:         "bread3",
				MerchantId: merchantIds[1],
				Name:       "goodName3",
				Category:   "goodCat3",
				Price:      3000,
				ImageUrl:   "https://lmao.image/3",
				CreatedAt:  time.Now().Add(time.Duration(-1)),
			},
		}

		for _, item := range items {
			err := merchantRepo.CreateMerchantItem(context.TODO(), item)
			So(err, ShouldBeNil)
		}

		estimates := []order.Estimate{
			{
				Id:     "gamer-moment1",
				UserId: "epic-gamer",
				Items: []order.EstimateItem{
					{
						MerchantId: items[1].MerchantId,
						ItemId:     items[1].Id,
						Quantity:   1,
					},
					{
						MerchantId: items[2].MerchantId,
						ItemId:     items[2].Id,
						Quantity:   2,
					},
				},
			},
			{
				Id:     "gamer-moment2",
				UserId: "epic-gamer",
				Items: []order.EstimateItem{
					{
						MerchantId: items[0].MerchantId,
						ItemId:     items[0].Id,
						Quantity:   1,
					},
				},
			},
			{
				Id:     "gamer-moment3",
				UserId: "epic-gamer",
				Items: []order.EstimateItem{
					{
						MerchantId: items[0].MerchantId,
						ItemId:     items[0].Id,
						Quantity:   1,
					},
					{
						MerchantId: items[2].MerchantId,
						ItemId:     items[2].Id,
						Quantity:   2,
					},
				},
			},
		}

		for _, estimate := range estimates {
			err := repo.CreateEstimate(context.TODO(), estimate)
			So(err, ShouldBeNil)
		}

		orders := []order.Order{
			{
				OrderId:    "gamer-order1",
				EstimateId: "gamer-moment1",
			},
			{
				OrderId:    "gamer-order2",
				EstimateId: "gamer-moment2",
			},
			{
				OrderId:    "gamer-order3",
				EstimateId: "gamer-moment3",
			},
		}

		for _, order := range orders {
			err := repo.CreateOrder(context.TODO(), order)
			So(err, ShouldBeNil)
		}

		Convey("When querying with MerchantIds and Name Filter", func() {
			Convey("Should return the matching merchants and itemName only", func() {
				Convey("When mid and itemName exists", func() {
					expectedNameFilter := "me3"
					req := order.OrderSummaryListFilter{
						MerchantIds: []string{merchantIds[0]},
						Limit:       5,
						Offset:      0,
						ItemName:    &expectedNameFilter,
						UserId:      "epic-gamer",
					}

					res, err := repo.ListOrderSummary(context.TODO(), req)
					So(err, ShouldBeNil)
					So(res, ShouldHaveLength, 5)

					var returnedItemNames []string
					for _, ordreSummary := range res {
						returnedItemNames = append(
							returnedItemNames,
							ordreSummary.ItemName,
						)
					}

					So(returnedItemNames, ShouldContain, items[2].Name)
					So(returnedItemNames, ShouldContain, items[1].Name)
					So(returnedItemNames, ShouldContain, items[0].Name)
				})
				Convey("When mid not extist but itemName exists", func() {
					expectedMid := "gyattlmaolol"
					expectedNameFilter := "me3"
					req := order.OrderSummaryListFilter{
						MerchantIds: []string{expectedMid},
						Limit:       5,
						Offset:      0,
						ItemName:    &expectedNameFilter,
						UserId:      "epic-gamer",
					}

					res, err := repo.ListOrderSummary(context.TODO(), req)
					So(err, ShouldBeNil)
					So(res, ShouldHaveLength, 2)

					var returnedItemNames []string
					for _, ordreSummary := range res {
						returnedItemNames = append(
							returnedItemNames,
							ordreSummary.ItemName,
						)
					}

					So(returnedItemNames, ShouldContain, items[2].Name)
					So(returnedItemNames, ShouldNotContain, items[1].Name)
					So(returnedItemNames, ShouldNotContain, items[0].Name)
				})
				Convey("When mid and itemName not extist", func() {
					expectedMid := "gyattlmaolol"
					expectedNameFilter := "me4"
					req := order.OrderSummaryListFilter{
						MerchantIds: []string{expectedMid},
						Limit:       5,
						Offset:      0,
						ItemName:    &expectedNameFilter,
						UserId:      "epic-gamer",
					}

					res, err := repo.ListOrderSummary(context.TODO(), req)
					So(err, ShouldBeNil)
					So(res, ShouldHaveLength, 0)
				})
			})
		})

		Convey("When querying with MerchantIds Filter", func() {
			Convey("Should return the matching merchants only", func() {
				Convey("When mid exists", func() {
					req := order.OrderSummaryListFilter{
						MerchantIds: []string{merchantIds[1]},
						Limit:       5,
						Offset:      0,
						UserId:      "epic-gamer",
					}

					res, err := repo.ListOrderSummary(context.TODO(), req)
					So(err, ShouldBeNil)
					So(res, ShouldHaveLength, 2)

					var returnedItemNames []string
					for _, ordreSummary := range res {
						returnedItemNames = append(
							returnedItemNames,
							ordreSummary.ItemName,
						)
					}

					So(returnedItemNames, ShouldContain, items[2].Name)
					So(returnedItemNames, ShouldNotContain, items[1].Name)
					So(returnedItemNames, ShouldNotContain, items[0].Name)
				})
				Convey("When mid not exists", func() {
					req := order.OrderSummaryListFilter{
						MerchantIds: []string{merchantIds[2]},
						Limit:       5,
						Offset:      0,
						UserId:      "epic-gamer",
					}

					res, err := repo.ListOrderSummary(context.TODO(), req)
					So(err, ShouldBeNil)
					So(res, ShouldHaveLength, 0)
				})
			})
		})

		Convey("When querying with ItemName Filter", func() {
			Convey("Should return the matching itemName only", func() {
				Convey("When itemName exists", func() {
					expectedNameFilter := "me3"
					req := order.OrderSummaryListFilter{
						Limit:    5,
						Offset:   0,
						ItemName: &expectedNameFilter,
						UserId:   "epic-gamer",
					}

					res, err := repo.ListOrderSummary(context.TODO(), req)
					So(err, ShouldBeNil)
					So(res, ShouldHaveLength, 2)

					var returnedItemNames []string
					for _, ordreSummary := range res {
						returnedItemNames = append(
							returnedItemNames,
							ordreSummary.ItemName,
						)
					}

					So(returnedItemNames, ShouldContain, items[2].Name)
					So(returnedItemNames, ShouldNotContain, items[1].Name)
					So(returnedItemNames, ShouldNotContain, items[0].Name)
				})
				Convey("When itemName not exists", func() {
					expectedNameFilter := "meLmaoJuan"
					req := order.OrderSummaryListFilter{
						Limit:    5,
						Offset:   0,
						ItemName: &expectedNameFilter,
						UserId:   "epic-gamer",
					}

					res, err := repo.ListOrderSummary(context.TODO(), req)
					So(err, ShouldBeNil)
					So(res, ShouldHaveLength, 0)
				})
			})
		})

		Convey("When querying with no filter", func() {
			Convey("Should return the matching UserId only", func() {
				req := order.OrderSummaryListFilter{
					Limit:  5,
					Offset: 0,
					UserId: "epic-gamer",
				}

				res, err := repo.ListOrderSummary(context.TODO(), req)
				So(err, ShouldBeNil)
				So(res, ShouldHaveLength, 5)

				var returnedItemNames []string
				for _, ordreSummary := range res {
					returnedItemNames = append(
						returnedItemNames,
						ordreSummary.ItemName,
					)
				}

				So(returnedItemNames, ShouldContain, items[2].Name)
				So(returnedItemNames, ShouldContain, items[1].Name)
				So(returnedItemNames, ShouldContain, items[0].Name)
			})
		})
	})
}
