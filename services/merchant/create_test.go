package merchant

import (
	"context"
	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	"github.com/JesseNicholas00/BeliMang/types/location"
	"github.com/JesseNicholas00/BeliMang/utils/helper"
	"github.com/JesseNicholas00/BeliMang/utils/unittesting"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCreateMerchant(t *testing.T) {
	Convey("When creating a merchant", t, func() {
		mockCtrl, service, mockedRepo := NewWithMockedRepo(t)
		defer mockCtrl.Finish()

		name := "bread"
		category := "SmallRestaurant"
		imageUrl := "https://bread.com/bread.png"
		coords := location.Location{
			Latitude:  helper.ToPointer(6.9),
			Longitude: helper.ToPointer(42.0),
		}

		Convey("When called with a valid request", func() {
			req := CreateMerchantReq{
				Name:     name,
				Category: category,
				ImageUrl: imageUrl,
				Location: coords,
			}
			var res CreateMerchantRes

			Convey("Should pass the request to the repo layer", func() {
				unittesting.FixNextUuid()
				id := uuid.NewString()

				m := merchant.Merchant{
					Id:        id,
					Name:      req.Name,
					Category:  req.Category,
					ImageUrl:  req.ImageUrl,
					Latitude:  *req.Location.Latitude,
					Longitude: *req.Location.Longitude,
				}

				mockedRepo.
					EXPECT().
					CreateMerchant(gomock.Any(), m).
					Return(nil).
					Times(1)

				Convey("And return the expected id", func() {
					unittesting.FixNextUuid()
					err := service.CreateMerchant(context.TODO(), req, &res)
					So(err, ShouldBeNil)
					So(res.Id, ShouldEqual, id)
				})
			})
		})
	})
}
