//go:build integration
// +build integration

package merchant_test

import (
	"context"
	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCreateMerchant(t *testing.T) {
	Convey("When creating a new merchant", t, func() {
		repo := NewWithTestDatabase(t)

		reqMerchant := merchant.Merchant{
			Id:        "bread",
			Name:      "chocolate croissant",
			Category:  "SmallRestaurant",
			ImageUrl:  "https://bread.com/bread.png",
			Latitude:  6.9,
			Longitude: 42.0,
		}

		Convey("Should return nil", func() {
			err := repo.CreateMerchant(context.TODO(), reqMerchant)
			So(err, ShouldBeNil)
		})
	})
}
