package merchant_test

import (
	"context"
	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestListByIds(t *testing.T) {
	Convey("When listing by IDs", t, func() {
		repo := NewWithTestDatabase(t)
		merchants := []merchant.Merchant{
			{
				Id:        "bruh",
				Name:      "bruh1",
				Category:  "cat1",
				ImageUrl:  "http://image.com/image.png",
				Latitude:  4.20,
				Longitude: 6.9,
			},
			{
				Id:        "bruh2",
				Name:      "bruh3",
				Category:  "cat1",
				ImageUrl:  "http://image.com/image.png",
				Latitude:  4.20,
				Longitude: 6.9,
			},
			{
				Id:        "bruh4",
				Name:      "bruh5",
				Category:  "cat1",
				ImageUrl:  "http://image.com/image.png",
				Latitude:  4.20,
				Longitude: 6.9,
			},
		}

		for _, m := range merchants {
			err := repo.CreateMerchant(context.TODO(), m)
			So(err, ShouldBeNil)
		}

		Convey("Should return only the matching merchants", func() {
			res, err := repo.ListMerchantsByIds(
				context.TODO(),
				[]string{"bruh2", "bruh4"},
			)
			So(err, ShouldBeNil)

			So(res, ShouldHaveLength, 2)
			var resIds []string
			for _, m := range res {
				resIds = append(resIds, m.Id)
			}
			So(resIds, ShouldContain, "bruh2")
			So(resIds, ShouldContain, "bruh4")
		})
	})
}
