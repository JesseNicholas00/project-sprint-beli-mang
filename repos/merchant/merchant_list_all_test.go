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

func TestListAllMerchant(t *testing.T) {
	Convey("With dummy data", t, func() {
		repo := NewWithTestDatabase(t)

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
			err := repo.CreateMerchant(context.TODO(), merchant)
			So(err, ShouldBeNil)
		}

		Convey("When querying with MerchantIdFilter filter", func() {
			Convey("Should return the matching merchants only", func() {
				Convey("When MID exists", func() {
					expectedMid := "220132512344"
					req := merchant.MerchantListAllFilter{
						MerchantId: &expectedMid,
					}

					res, err := repo.ListAllMerchant(context.TODO(), req)
					So(err, ShouldBeNil)
					So(res, ShouldHaveLength, 1)

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
					req := merchant.MerchantListAllFilter{
						MerchantId: &expectedMid,
					}

					res, err := repo.ListAllMerchant(context.TODO(), req)
					So(err, ShouldBeNil)
					So(res, ShouldHaveLength, 0)
				})
			})
		})

		Convey("When querying with name filter", func() {
			Convey("Should return the matching merchants only", func() {
				nameFilter := "pic"
				req := merchant.MerchantListAllFilter{
					Name: &nameFilter,
				}

				res, err := repo.ListAllMerchant(context.TODO(), req)
				So(err, ShouldBeNil)
				So(res, ShouldHaveLength, 2)

				var returnedNames []string
				for _, merchant := range res {
					returnedNames = append(
						returnedNames,
						merchant.Name,
					)
				}

				So(returnedNames, ShouldContain, names[0])
				So(returnedNames, ShouldContain, names[1])
				So(returnedNames, ShouldNotContain, names[2])
			})
		})

		Convey("When querying with merchantCategory filter", func() {
			Convey("Should return the matching merchants only", func() {
				catFilter := "lolMewGyattRizz"
				req := merchant.MerchantListAllFilter{
					MerchantCategory: &catFilter,
				}

				res, err := repo.ListAllMerchant(context.TODO(), req)
				So(err, ShouldBeNil)
				So(res, ShouldHaveLength, 1)

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
