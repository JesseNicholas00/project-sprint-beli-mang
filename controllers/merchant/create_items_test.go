package merchant

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JesseNicholas00/BeliMang/services/merchant"
	"github.com/JesseNicholas00/BeliMang/utils/helper"
	"github.com/JesseNicholas00/BeliMang/utils/unittesting"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateMerchantItemValid(t *testing.T) {
	Convey("Given a valid request", t, func() {
		mockCtrl, controller, mockedService := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		id := "bread-press"
		merchantId := "merchantId"
		name := "bread"
		category := "Beverage"
		imageUrl := "https://bread.com/bread.jpg"
		price := 10000
		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodPost,
			"/admin/merchants/merchantId/items",
			rec,
			unittesting.WithJsonPayload(map[string]interface{}{
				"name":            name,
				"productCategory": category,
				"imageUrl":        imageUrl,
				"price":           price,
			}),
			unittesting.WithPathParams(map[string]string{
				"merchantId": merchantId,
			}),
		)

		Convey("Should forward the request to the service layer", func() {
			mockedService.
				EXPECT().
				CreateMerchantItem(gomock.Any(), gomock.Any(), gomock.Any()).
				Do(
					func(
						_ context.Context,
						req merchant.CreateMerchantItemReq,
						res *merchant.CreateMerchantItemRes,
					) {
						So(req.Name, ShouldEqual, name)
						So(req.Category, ShouldEqual, category)
						So(req.ImageUrl, ShouldEqual, imageUrl)
						So(req.Price, ShouldEqual, price)
						res.ItemId = id
					},
				).
				Return(nil).
				Times(1)

			unittesting.CallController(ctx, controller.createMerchantItems)

			Convey("And return the results with status 201", func() {
				expectedRes := merchant.CreateMerchantItemRes{ItemId: id}
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

func TestCreateMerchantItemInvalid(t *testing.T) {
	Convey("Given an invalid request", t, func() {
		mockCtrl, controller, _ := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		merchantId := "merchantId"
		name := "bread"
		// bad category name
		category := "wrong category lol"
		imageUrl := "https://bread.com/bread.png"
		price := 10000

		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodPost,
			"/admin/merchants/merchantId/items",
			rec,
			unittesting.WithJsonPayload(map[string]interface{}{
				"name":            name,
				"productCategory": category,
				"imageUrl":        imageUrl,
				"price":           price,
			}),
			unittesting.WithPathParams(map[string]string{
				"merchantId": merchantId,
			}),
		)

		Convey("Should return error 400", func() {
			unittesting.CallController(ctx, controller.createMerchantItems)
			So(rec.Code, ShouldEqual, http.StatusBadRequest)
		})
	})
}
