package merchant

import (
	"context"
	"github.com/JesseNicholas00/BeliMang/services/merchant"
	"github.com/JesseNicholas00/BeliMang/types/location"
	"github.com/JesseNicholas00/BeliMang/utils/helper"
	"github.com/JesseNicholas00/BeliMang/utils/unittesting"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateMerchantValid(t *testing.T) {
	Convey("Given a valid request", t, func() {
		mockCtrl, controller, mockedService := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		id := "bread-press"
		name := "bread"
		category := "SmallRestaurant"
		imageUrl := "https://bread.com/bread.png"
		coords := location.Location{
			Latitude:  helper.ToPointer(float32(6.9)),
			Longitude: helper.ToPointer(float32(42.0)),
		}

		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodPost,
			"/admin/merchants",
			rec,
			unittesting.WithJsonPayload(map[string]interface{}{
				"name":             name,
				"merchantCategory": category,
				"imageUrl":         imageUrl,
				"location": map[string]interface{}{
					"lat":  coords.Latitude,
					"long": coords.Longitude,
				},
			}),
		)

		Convey("Should forward the request to the service layer", func() {
			mockedService.
				EXPECT().
				CreateMerchant(gomock.Any(), gomock.Any(), gomock.Any()).
				Do(
					func(
						_ context.Context,
						req merchant.CreateMerchantReq,
						res *merchant.CreateMerchantRes,
					) {
						So(req.Name, ShouldEqual, name)
						So(req.Category, ShouldEqual, category)
						So(req.ImageUrl, ShouldEqual, imageUrl)
						So(
							helper.IsEqualFloat(
								*req.Location.Longitude,
								*coords.Longitude,
							),
							ShouldBeTrue,
						)
						So(
							helper.IsEqualFloat(
								*req.Location.Longitude,
								*coords.Longitude,
							),
							ShouldBeTrue,
						)
						res.Id = id
					},
				).
				Return(nil).
				Times(1)

			unittesting.CallController(ctx, controller.createMerchant)

			Convey("And return the results with status 201", func() {
				expectedRes := merchant.CreateMerchantRes{Id: id}
				So(rec.Code, ShouldEqual, http.StatusCreated)
				So(
					rec.Body.String(),
					ShouldEqualJSON,
					string(helper.MustMarshalJson(expectedRes)),
				)
			})
		})
	})
}

func TestCreateMerchantInvalid(t *testing.T) {
	Convey("Given an invalid request", t, func() {
		mockCtrl, controller, _ := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		name := "bread"
		// bad category name
		category := "wrong category lol"
		imageUrl := "https://bread.com/bread.png"
		coords := location.Location{
			Latitude:  helper.ToPointer(float32(6.9)),
			Longitude: helper.ToPointer(float32(42.0)),
		}

		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodPost,
			"/admin/merchants",
			rec,
			unittesting.WithJsonPayload(map[string]interface{}{
				"name":             name,
				"merchantCategory": category,
				"imageUrl":         imageUrl,
				"location": map[string]interface{}{
					"lat":  coords.Latitude,
					"long": coords.Longitude,
				},
			}),
		)

		Convey("Should return error 400", func() {
			unittesting.CallController(ctx, controller.createMerchant)
			So(rec.Code, ShouldEqual, http.StatusBadRequest)
		})
	})
}
