package merchant_test

import (
	"context"
	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestListItemsByIds(t *testing.T) {
	Convey("When listing items by IDs", t, func() {
		repo := NewWithTestDatabase(t)
		items := []merchant.MerchantItem{
			{
				Id:         "bread1",
				MerchantId: "asdf",
			},
			{
				Id:         "bread2",
				MerchantId: "asdf",
			},
			{
				Id:         "bread3",
				MerchantId: "hjkl",
			},
		}

		for _, it := range items {
			err := repo.CreateMerchantItem(context.TODO(), it)
			So(err, ShouldBeNil)
		}

		Convey("Should return only the matching items", func() {
			res, err := repo.ListItemsByIds(
				context.TODO(),
				[]string{"bread1", "bread3"},
			)
			So(err, ShouldBeNil)

			So(res, ShouldHaveLength, 2)
			var resIds []string
			for _, m := range res {
				resIds = append(resIds, m.Id)
			}
			So(resIds, ShouldContain, "bread1")
			So(resIds, ShouldContain, "bread3")
		})
	})
}
