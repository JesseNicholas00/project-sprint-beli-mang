package merchant

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JesseNicholas00/BeliMang/services/merchant"
	"github.com/JesseNicholas00/BeliMang/types/location"
	"github.com/JesseNicholas00/BeliMang/types/pagination"
	"github.com/JesseNicholas00/BeliMang/utils/helper"
	"github.com/JesseNicholas00/BeliMang/utils/unittesting"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAdminList(t *testing.T) {
	Convey("When given a valid request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		merchantId := "220132512342"
		limit := 5
		offset := 0
		name := "epic"
		merchantCategory := "goodCategory"
		createdAtSort := "wrong"
		lat := float64(1.1)
		lon := float64(2.2)

		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodGet,
			"/admin/merchants",
			rec,
			unittesting.WithQueryParams(map[string]string{
				"merchantId":       fmt.Sprint(merchantId),
				"name":             name,
				"merchantCategory": merchantCategory,
				"createdAt":        createdAtSort,
			}),
		)

		Convey("Should forward the request to the service layer", func() {
			expectedReq := merchant.AdminListMerchantReq{
				MerchantId:       &merchantId,
				Name:             &name,
				MerchantCategory: &merchantCategory,
				CreatedAtSort:    nil,
				Limit:            &limit,
				Offset:           &offset,
			}

			expectedRes := []merchant.ListMerchantResData{
				{
					MerchantId:       merchantId,
					Name:             name,
					MerchantCategory: merchantCategory,
					ImageUrl:         "https://verygoodImage.com",
					Location:         location.Location{Latitude: &lat, Longitude: &lon},
					CreatedAt:        "now",
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
					AdminListMerchant(gomock.Any(), expectedReq, gomock.Any()).
					Do(
						func(
							_ context.Context,
							req merchant.AdminListMerchantReq,
							res *merchant.AdminListMerchantRes,
						) {
							res.Data = expectedRes
							res.Meta = expectedMeta
						},
					).
					Return(nil).
					Times(1)

				unittesting.CallController(ctx, controller.adminList)

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
					AdminListMerchant(gomock.Any(), expectedReq, gomock.Any()).
					Do(
						func(
							_ context.Context,
							req merchant.AdminListMerchantReq,
							res *merchant.AdminListMerchantRes,
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

				unittesting.CallController(ctx, controller.adminList)

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
