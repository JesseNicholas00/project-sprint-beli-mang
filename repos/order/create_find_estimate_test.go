package order_test

import (
	"context"
	"errors"
	"testing"

	"github.com/JesseNicholas00/BeliMang/repos/order"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateAndFindEstimate(t *testing.T) {
	repo, _ := NewWithTestDatabase(t)
	Convey("When attempting to retrieve a non-existent estimate", t, func() {
		Convey("Should return err", func() {
			_, err := repo.FindEstimateById(context.TODO(), "none")
			So(errors.Is(err, order.ErrEstimateNotFound), ShouldBeTrue)
		})
	})

	Convey("When creating an estimate", t, func() {
		est := order.Estimate{
			Id:     "gamer-moment",
			UserId: "epic-gamer",
			Items: []order.EstimateItem{
				{
					MerchantId: "gamer-1",
					ItemId:     "juan",
					Quantity:   1,
				},
				{
					MerchantId: "gamer-2",
					ItemId:     "henry",
					Quantity:   2,
				},
				{
					MerchantId: "gamer-3",
					ItemId:     "gyatt",
					Quantity:   3,
				},
			},
		}

		Convey("Should not error out", func() {
			err := repo.CreateEstimate(context.TODO(), est)
			So(err, ShouldBeNil)

			Convey("And the item should be retrievable", func() {
				res, err := repo.FindEstimateById(context.TODO(), est.Id)
				So(err, ShouldBeNil)
				// copy over createdAt
				est.CreatedAt = res.CreatedAt
				So(res, ShouldEqual, est)
			})
		})
	})
}
