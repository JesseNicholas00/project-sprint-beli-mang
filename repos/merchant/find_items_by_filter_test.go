//go:build integration
// +build integration

package merchant_test

import (
	"context"
	"testing"
	"time"

	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMerchantItemList(t *testing.T) {
	Convey("With dummy data", t, func() {
		repo := NewWithTestDatabase(t)

		merchantIds := []string{"220132512342", "220132512344", "220132512343"}
		names := []string{"epic", "epic_gamer_2", "gamer3"}
		prices := []int64{100, 200, 300}
		merchantCategory := []string{"goodCategory", "badCategory", "lolMewGyattRizz"}
		createtAts := []time.Time{
			time.Now(),
			time.Now().AddDate(0, 0, 2),
			time.Now().AddDate(0, 0, 3),
		}

		var merchantItems []merchant.MerchantItem
		for i := 0; i < 3; i++ {
			merchantItems = append(
				merchantItems,
				merchant.MerchantItem{
					Name:      names[i],
					Id:        merchantIds[i],
					Price:     prices[i],
					Category:  merchantCategory[i],
					ImageUrl:  "https://bread.com/bread.png",
					CreatedAt: createtAts[i],
				},
			)
		}

		for _, merchant := range merchantItems {
			err := repo.CreateMerchantItem(context.TODO(), merchant)
			So(err, ShouldBeNil)
		}

		Convey("When querying with MerchantIdFilter filter", func() {
			Convey("Should return the matching merchantItems only", func() {
				Convey("When MID exists", func() {
					expectedMid := "220132512344"
					req := merchant.MerchantItemListFilter{
						MerchantItemId: &expectedMid,
						Limit:          5,
						Offset:         0,
					}

					res, total, err := repo.FindMerchantItemsByFilter(context.TODO(), req)
					So(err, ShouldBeNil)
					So(res, ShouldHaveLength, 1)
					So(total, ShouldEqual, 1)

					var returnedNames []string
					for _, merchant := range res {
						returnedNames = append(
							returnedNames,
							merchant.Name,
						)
					}

					So(returnedNames, ShouldContain, names[1])
					So(returnedNames, ShouldNotContain, names[0])
					So(returnedNames, ShouldNotContain, names[2])
				})
				Convey("When MID not exists", func() {
					expectedMid := "lmaoGyattLoL"
					req := merchant.MerchantItemListFilter{
						MerchantItemId: &expectedMid,
						Limit:          5,
						Offset:         0,
					}

					res, total, err := repo.FindMerchantItemsByFilter(context.TODO(), req)
					So(err, ShouldBeNil)
					So(res, ShouldHaveLength, 0)
					So(total, ShouldEqual, 0)
				})
			})
		})

		Convey("When querying with name filter", func() {
			Convey("Should return the matching merchantItems only", func() {
				nameFilter, sort := "pic", "asc"
				req := merchant.MerchantItemListFilter{
					Name:          &nameFilter,
					Limit:         1,
					Offset:        0,
					CreatedAtSort: &sort,
				}

				res, total, err := repo.FindMerchantItemsByFilter(context.TODO(), req)
				So(err, ShouldBeNil)
				So(res, ShouldHaveLength, 1)
				So(total, ShouldEqual, 2)

				var returnedNames []string
				for _, merchant := range res {
					returnedNames = append(
						returnedNames,
						merchant.Name,
					)
				}

				So(returnedNames, ShouldContain, names[0])
				So(returnedNames, ShouldNotContain, names[1])
				So(returnedNames, ShouldNotContain, names[2])
			})
		})

		Convey("When querying with merchantCategory filter", func() {
			Convey("Should return the matching merchantItems only", func() {
				thirdFilter := "lolMewGyattRizz"
				req := merchant.MerchantItemListFilter{
					Category: &thirdFilter,
					Limit:    10,
					Offset:   0,
				}

				res, total, err := repo.FindMerchantItemsByFilter(context.TODO(), req)
				So(err, ShouldBeNil)
				So(res, ShouldHaveLength, 1)
				So(total, ShouldEqual, 1)

				var returnedNames []string
				for _, merchant := range res {
					returnedNames = append(
						returnedNames,
						merchant.Name,
					)
				}

				So(returnedNames, ShouldContain, names[2])
				So(returnedNames, ShouldNotContain, names[1])
				So(returnedNames, ShouldNotContain, names[0])
			})
		})
	})
}
