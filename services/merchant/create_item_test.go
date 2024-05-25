package merchant

import (
	"context"
	"testing"

	merchantRepo "github.com/JesseNicholas00/BeliMang/repos/merchant"
	"github.com/JesseNicholas00/BeliMang/types/location"
	"github.com/JesseNicholas00/BeliMang/utils/helper"
	"github.com/JesseNicholas00/BeliMang/utils/unittesting"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateMerchantItem(t *testing.T) {
	Convey("When creating a merchant item", t, func() {
		mockCtrl, service, mockedRepo := NewWithMockedRepo(t)
		defer mockCtrl.Finish()

		name := "bread"
		category := "SmallRestaurant"
		imageUrl := "https://bread.com/bread.png"
		price := 1000
		coords := location.Location{
			Latitude:  helper.ToPointer(6.9),
			Longitude: helper.ToPointer(42.0),
		}
		merchantId := uuid.NewString()

		Convey("When called with a valid request", func() {
			req := CreateMerchantItemReq{
				MerchantId: merchantId,
				Name:       name,
				Category:   category,
				Price:      price,
				ImageUrl:   imageUrl,
			}

			var res CreateMerchantItemRes

			Convey("Should pass the request to the repo layer", func() {
				unittesting.FixNextUuid()
				id := uuid.NewString()

				merchant := merchantRepo.Merchant{
					Id:        merchantId,
					Name:      req.Name,
					Category:  req.Category,
					ImageUrl:  req.ImageUrl,
					Latitude:  *coords.Latitude,
					Longitude: *coords.Longitude,
				}

				merchantItem := merchantRepo.MerchantItem{
					Id:         id,
					MerchantId: merchantId,
					Name:       req.Name,
					Category:   req.Category,
					Price:      req.Price,
					ImageUrl:   req.ImageUrl,
				}

				mockedRepo.
					EXPECT().
					FindMerchantById(gomock.Any(), req.MerchantId).
					Return(merchant, nil).
					Times(1)

				mockedRepo.
					EXPECT().
					CreateMerchantItem(gomock.Any(), merchantItem).
					Return(nil).
					Times(1)

				Convey("And return the expected id", func() {
					unittesting.FixNextUuid()
					err := service.CreateMerchantItem(context.TODO(), req, &res)
					So(err, ShouldBeNil)
					So(res.ItemId, ShouldEqual, id)
				})
			})
		})
	})
}
