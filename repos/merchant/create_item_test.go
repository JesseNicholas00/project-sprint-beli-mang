//go:build integration
// +build integration

package merchant_test

import (
	"context"
	"testing"

	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateMerchantItem(t *testing.T) {
	Convey("When creating a new merchant item", t, func() {
		repo := NewWithTestDatabase(t)

		reqMerchantItem := merchant.MerchantItem{
			Id:       "bread",
			Name:     "chocolate croissant",
			Category: "SmallRestaurant",
			Price:    1001231,
			ImageUrl: "https://bread.com/bread.png",
		}

		Convey("Should return nil", func() {
			err := repo.CreateMerchantItem(context.TODO(), reqMerchantItem)
			So(err, ShouldBeNil)
		})
	})
}
