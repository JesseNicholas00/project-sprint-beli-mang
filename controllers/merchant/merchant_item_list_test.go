package merchant

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JesseNicholas00/BeliMang/services/merchant"
	"github.com/JesseNicholas00/BeliMang/types/pagination"
	"github.com/JesseNicholas00/BeliMang/utils/helper"
	"github.com/JesseNicholas00/BeliMang/utils/unittesting"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMerchantItemList(t *testing.T) {
	Convey("When given a valid request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		merchantId := "merchantId"
		merchantItemId := "merchantItemId"
		limit := 5
		offset := 0
		name := "epic"
		productCategory := "goodCategory"
		createdAtSort := "wrong"

		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodGet,
			"/admin/merchants/merchantId/items",
			rec,
			unittesting.WithQueryParams(map[string]string{
				"itemId":          merchantItemId,
				"name":            name,
				"productCategory": productCategory,
				"createdAt":       createdAtSort,
			}),
			unittesting.WithPathParams(map[string]string{
				"merchantId": merchantId,
			}),
		)

		Convey("Should forward the request to the service layer", func() {
			expectedReq := merchant.MerchantItemListReq{
				MerchantItemId: &merchantItemId,
				Name:           &name,
				Category:       &productCategory,
				CreatedAtSort:  nil,
				Limit:          &limit,
				Offset:         &offset,
			}

			expectedRes := []merchant.ListMerchantItemResData{
				{
					MerchantItemId: merchantItemId,
					Name:           name,
					Category:       productCategory,
					Price:          100,
					ImageUrl:       "https://verygoodImage.com/image.jpg",
					CreatedAt:      "now",
				},
			}

			Convey("When the result data is not empty", func() {
				expectedTotal := int64(1)
				expectedMeta := pagination.Page{
					Offset: &offset,
					Limit:  &limit,
					Total:  &expectedTotal,
				}
				service.
					EXPECT().
					FindMerchantItemList(gomock.Any(), expectedReq, gomock.Any()).
					Do(
						func(
							_ context.Context,
							req merchant.MerchantItemListReq,
							res *merchant.MerchantItemListRes,
						) {
							res.Data = expectedRes
							res.Meta = expectedMeta
						},
					).
					Return(nil).
					Times(1)

				unittesting.CallController(ctx, controller.findMerchantItemList)

				Convey(
					"Should return HTTP 200 and the resulting array",
					func() {
						So(rec.Code, ShouldEqual, http.StatusOK)

						expectedBody := helper.MustMarshalJson(
							map[string]interface{}{
								"meta": expectedMeta,
								"data": expectedRes,
							},
						)

						So(
							rec.Body.String(),
							ShouldEqualJSON,
							string(expectedBody),
						)
					},
				)
			})

			Convey("When the result data is empty", func() {
				expectedTotal := int64(0)
				expectedMeta := pagination.Page{
					Offset: &offset,
					Limit:  &limit,
					Total:  &expectedTotal,
				}
				service.
					EXPECT().
					FindMerchantItemList(gomock.Any(), expectedReq, gomock.Any()).
					Do(
						func(
							_ context.Context,
							req merchant.MerchantItemListReq,
							res *merchant.MerchantItemListRes,
						) {
							res.Data = nil
							res.Meta = pagination.Page{
								Offset: req.Offset,
								Limit:  req.Limit,
								Total:  &expectedTotal,
							}
						},
					).
					Return(nil).
					Times(1)

				unittesting.CallController(ctx, controller.findMerchantItemList)

				Convey(
					"Should return HTTP 200 and an empty array",
					func() {
						So(rec.Code, ShouldEqual, http.StatusOK)

						expectedBody := helper.MustMarshalJson(
							map[string]interface{}{
								"meta": expectedMeta,
								"data": make([]struct{}, 0),
							},
						)

						So(
							rec.Body.String(),
							ShouldEqualJSON,
							string(expectedBody),
						)
					},
				)
			})
		})
	})
}
